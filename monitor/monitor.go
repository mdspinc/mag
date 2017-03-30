// Package monitor track some kind of error in external services.
package monitor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/mdspinc/mag/sender"
)

// MonitorMessage is a notification message template.
const MonitorMessage = "Number of `FREQUENCY_KEY_PRESENT` errors for all campaigns is %d for last *%d* seconds."

type (
	// Monitor is a monitoring data and settings.
	Monitor struct {
		// Source of data
		BotsmetricsAPIAddress string
		BotsmetricsAPIToken   string
		interval              int
		ticker                *time.Ticker
		max                   int
		store                 []Store
		sender                sender.Transport
		fkpThreshold          int
	}

	// Data is a struct for JSON unmarshalng.
	data struct {
		Campaigns []struct {
			ID          string `json:"id"`
			Impressions int    `json:"imps"`
			Errors      struct {
				FrequencyKeyPresent int `json:"FREQUENCY_KEY_PRESENT"`
			} `json:"errors"`
		}
	}
)

// New initialized Monitor struct.
func New(address, token string, interval, maxItems, fkpThreshold int) *Monitor {
	ss, err := sender.NewSlackSender()
	if err != nil {
		log.Println("monitor: New: error:", err)
	}

	return &Monitor{
		BotsmetricsAPIAddress: address,
		BotsmetricsAPIToken:   token,
		interval:              interval,
		ticker:                time.NewTicker(time.Second * time.Duration(interval)),
		max:                   maxItems,
		sender:                ss,
		fkpThreshold:          fkpThreshold,
	}
}

// Start runs monitor process.
func (m *Monitor) Start() {
	if m.BotsmetricsAPIAddress == "" {
		log.Println("monitor: BotsmetricsApiAddress is empty, monitoring is disabled")
		return
	}

	go func() {
		for {
			<-m.ticker.C
			err := m.Fetch()
			if err != nil {
				log.Println(err)
			}

			m.Check()
		}
	}()

	log.Printf(
		"Monitor started with %d seconds interval and %d FKP threshold. \n",
		m.interval, m.fkpThreshold)
}

// Fetch requests data from API.
func (m *Monitor) Fetch() error {
	cc := &data{}
	r := NewStore()

	req, err := http.NewRequest("GET", m.BotsmetricsAPIAddress, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Bearer "+m.BotsmetricsAPIToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK ||
		resp.StatusCode != http.StatusTemporaryRedirect ||
		resp.StatusCode != http.StatusPermanentRedirect {
		return fmt.Errorf("monitor fetch: bad response status: %d", resp.StatusCode)
	}

	if err := json.Unmarshal(body, &cc); err != nil {
		return err
	}

	for _, c := range cc.Campaigns {
		r[c.ID] = newStoredCamp(c.Errors.FrequencyKeyPresent, c.Impressions)
	}

	m.store = append(m.store, r)

	if len(m.store) > m.max {
		m.store = append(m.store[:0], m.store[1:]...)
	}

	return nil
}

// Check checks number of errors of type FREQUENCY_KEY_PRESENT.
// If number of errors for all campaigns is less than FKP_THRESHOLD
// send notification to slack.
func (m *Monitor) Check() {
	if len(m.store) < 2 {
		return
	}

	last := len(m.store) - 1
	penult := len(m.store) - 2

	diff := m.store[last].Diff(m.store[penult])

	if diff < m.fkpThreshold {
		m.sender.Send(
			fmt.Sprintf(MonitorMessage, diff, m.interval), sender.MonitorMessage)
	}
}
