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
	msgTpl       = "Message `%q` was repeated *%d* time(s).\n"
	msgTplPeriod = "Message `%q` was repeated *%d* time(s) for last %.0f seconds.\n"
)

// StringAgg is a simple string aggregator which collects strings as is.
type StringAgg struct {
	sync.RWMutex

	//Buffer for messsages
	buffer map[string]int

	//Number of messages for Flush() call
	batchSize int

	//Transport used for sending messages.
	sender sender.Transport
}

// NewStringAgg initialises new StringAgg instance.
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

// Aggregate collects incoming messages.
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

// Flush flushes collected messages to sender.  It called if one of
// conditions happend: batchSize limit exceeded or period of time is up.
func (s *StringAgg) Flush(key string) {
	if err := s.sender.Send(fmt.Sprintf(msgTpl, key, s.buffer[key]), sender.DefaultMessage); err != nil {
		log.Println("agg: string: flush error:", err)
		return
	}

	// It don't need additional mutex lock here,
	// because it call under one from Aggregate() method.
	delete(s.buffer, key)
}

// FlushAll flushes all messages to sender and clear the buffer.
func (s *StringAgg) FlushAll(p time.Duration) {
	s.Lock()
	defer s.Unlock()
	b := bytes.NewBufferString("")
	for k, v := range s.buffer {
		b.WriteString(fmt.Sprintf(msgTplPeriod, k, v, p.Seconds()))
		delete(s.buffer, k)
	}
	if len(b.String()) > 0 {
		s.sender.Send(b.String(), sender.DefaultMessage)
	}
}

// Count returns number of collected messages.
func (s *StringAgg) Count(key string) int {
	s.RLock()
	defer s.RUnlock()
	return s.buffer[key]
}

// Len returns number of different messages
func (s *StringAgg) Len() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.buffer)
}
