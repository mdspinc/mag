package agg

import (
	"bytes"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/mdspinc/mag/sender"
)

const (
	MsgTpl       = "Message `%q` was repeated *%d* time(s).\n"
	MsgTplPeriod = "Message `%q` was repeated *%d* time(s) for last %.0f seconds.\n"
)

// Simple string aggregator which collects strings as is.
type StringAgg struct {
	sync.RWMutex

	//Buffer for messsages
	buffer map[string]int

	//Number of messages for Flush() call
	batchSize int

	//Transport used for sending messages.
	sender sender.Transport
}

// Initialises new StringAgg instance.
func NewStringAgg(bs int) *StringAgg {
	ss, err := sender.NewSlackSender()
	if err != nil {
		log.Println("agg: string: new slack sedner:", err)
	}

	return &StringAgg{
		buffer:    make(map[string]int),
		batchSize: bs,
		sender:    ss,
	}
}

// See Aggregator interface desciption.
func (s *StringAgg) Aggregate(data interface{}) {
	s.Lock()
	defer s.Unlock()

	switch t := data.(type) {
	case string:
		s.buffer[t]++
		if s.buffer[t] >= s.batchSize {
			s.Flush(t)
		}
	}
}

// See Aggregator interface desciption.
func (s *StringAgg) Flush(key string) {
	if err := s.sender.Send(fmt.Sprintf(MsgTpl, key, s.buffer[key]), sender.DEFAULT_MESSAGE); err != nil {
		log.Println("agg: string: flush error:", err)
		return
	}

	// It don't need additional mutex lock here,
	// because it call under one from Aggregate() method.
	delete(s.buffer, key)
}

// See Aggregator interface desciption.
func (s *StringAgg) FlushAll(p time.Duration) {
	s.Lock()
	defer s.Unlock()
	b := bytes.NewBufferString("")
	for k, v := range s.buffer {
		b.WriteString(fmt.Sprintf(MsgTplPeriod, k, v, p.Seconds()))
		delete(s.buffer, k)
	}
	if len(b.String()) > 0 {
		s.sender.Send(b.String(), sender.DEFAULT_MESSAGE)
	}
}

// See Aggregator interface desciption.
func (s *StringAgg) Count(key string) int {
	s.RLock()
	defer s.RUnlock()
	return s.buffer[key]
}

// Returnd number of different messages
func (s *StringAgg) Len() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.buffer)
}
