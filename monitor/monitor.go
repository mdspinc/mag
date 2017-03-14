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

type (
	Monitor struct {
		// Source of data
		BotsmetricsApiAddress string
		interval              int
		ticker                *time.Ticker
		max                   int
		store                 []map[string]*storedCamp
		sender                sender.Transport
	}

	// Campaign info stored in modnitor.
	storedCamp struct {
		impressions         int
		frequencyKeyPresent int
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
func New(address string, interval int, maxItems int) *Monitor {
	ss, err := sender.NewSlackSender()
	if err != nil {
		log.Println("monitor: New: error:", err)
	}

	return &Monitor{
		BotsmetricsApiAddress: address,
		interval:              interval,
		ticker:                time.NewTicker(time.Second * time.Duration(interval)),
		max:                   maxItems,
		sender:                ss,
	}
}

// Start runs monitor process.
func (m *Monitor) Start() {
	if m.BotsmetricsApiAddress == "" {
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

	log.Printf("Monitor started with %d seconds interval.\n", m.interval)
}

// Fetch requests data from API.
func (m *Monitor) Fetch() error {
	cc := &data{}
	r := make(map[string]*storedCamp)

	resp, err := http.Get(m.BotsmetricsApiAddress)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &cc); err != nil {
		return err
	}

	for _, c := range cc.Campaigns {
		r[c.ID] = &storedCamp{
			frequencyKeyPresent: c.Errors.FrequencyKeyPresent,
			impressions:         c.Impressions,
		}
	}

	m.store = append(m.store, r)

	if len(m.store) > m.max {
		m.store = append(m.store[:0], m.store[1:]...)
	}

	return nil
}

// Check check values if impressions and error FREQUENCY_KEY_PRESENT. If
// FREQUENCY_KEY_PRESENT still the same while impressions happend  and frequency
// capping(didn't check yet) is set than notify someone in slack.
func (m *Monitor) Check() {
	if len(m.store) < 2 {
		return
	}

	last := len(m.store) - 1
	prelast := len(m.store) - 2

	for k, v := range m.store[last] {
		imp1 := v.impressions
		fkp1 := v.frequencyKeyPresent
		if fkp1 == 0 {
			continue
		}
		imp2 := m.store[prelast][k].impressions
		fkp2 := m.store[prelast][k].frequencyKeyPresent
		if fkp1 == fkp2 && (imp2+3) < imp1 && imp1 > 0 {
			//fmt.Printf("id: %s\t fkp2: %d,  imp2: %d \t fkp1: %d, imp1: %d\n",
			//k, fkp2, imp2, fkp1, imp1)
			m.sender.Send(fmt.Sprintf(
				"Campaign: *%s*, imps changed: %d -> %d for last *%d* seconds, but `FREQUENCY_KEY_PRESENT` is the same: %d.",
				k, imp2, imp1, m.interval, fkp1,
			), sender.MONITOR_MESSAGE)
		}
	}
}
