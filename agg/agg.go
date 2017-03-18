// Package agg provides typed aggregators for collecting messages.
package agg

import (
	"errors"
	"time"
)

var (
	// ErrInvalidAggregatorType returned when aggregation type is not string.
	ErrInvalidAggregatorType = errors.New("agg: new: invalid aggregator type")
)

type (
	// Type is a aggregation type.
	Type byte

	// Aggregator is an interface for aggregation.
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
	// AggtypeString represents aggresation for strings only.
	AggtypeString Type = iota
)

// New initializes new Aggregator instance.
func New(t Type, batchSize int, flushPeriod time.Duration) (Aggregator, error) {
	switch t {
	case AggtypeString:
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
			a.FlushAll(fp)
		}
	}()
	return ticker
}
