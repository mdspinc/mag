package agg

import (
	"log"
	"sync"
)

type StringAgg struct {
	sync.RWMutex

	//Buffer for messsages
	buffer map[string]int

	//Number of messages for Flush() call
	batchSize int
}

func NewStringAgg(bs int) *StringAgg {
	return &StringAgg{
		buffer:    make(map[string]int),
		batchSize: bs,
	}
}

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

func (s *StringAgg) Flush(key string) {
	log.Println("flush:", key)

	delete(s.buffer, key)
}

func (s *StringAgg) Count(key string) int {
	s.RLock()
	defer s.RUnlock()
	return s.buffer[key]
}
