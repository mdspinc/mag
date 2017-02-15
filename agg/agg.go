// Package agg provides typed aggregators for collecting messages.
package agg

import (
	"errors"
	"log"
	"time"
)

var (
	ErrInvalidAggregatorType = errors.New("agg: new: invalid aggregator type.")
)

type (
	Type byte

	Aggregator interface {
		//Collects incoming messages.
		Aggregate(interface{})

		// Flushes collected messages to sender.
		// Flush() called if one of conditions happend:
		// batchSize limit exceeded or period of time is up.
		Flush(string)

		// Flushes all messages to sender and clear the buffer.
		FlushAll(time.Duration)

		//Returns number of collected messages.
		Count(string) int
	}
)

const (
	AGGTYPE_STRING Type = iota
)

// Initializes new Aggregator instance.
func New(t Type, batchSize int, flushPeriod time.Duration) (Aggregator, error) {
	switch t {
	case AGGTYPE_STRING:
		sa := NewStringAgg(batchSize)
		runTicker(flushPeriod, sa)
		return sa, nil
	default:
		return nil, ErrInvalidAggregatorType
	}
}

// Runs Ticker for flushing messages by time.
func runTicker(fp time.Duration, a Aggregator) *time.Ticker {
	ticker := time.NewTicker(fp)
	go func() {
		for range ticker.C {
			log.Println("FLUSH BY TIME")
			a.FlushAll(fp)
		}
	}()
	return ticker
}
