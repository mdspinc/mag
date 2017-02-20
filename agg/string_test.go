package agg

import (
	"testing"
	"time"
)

func TestAggregate(t *testing.T) {
	s, key, key2, _ := testInit(10)

	s.Aggregate(key)
	s.Aggregate(key)
	s.Aggregate(key2)

	if s.Count(key) != 2 {
		t.Error("invalid number of items")
	}

	if s.Count(key2) != 1 {
		t.Error("invalid number of items")
	}
}

func TestFlushes(t *testing.T) {
	s, key, key2, _ := testInit(10)

	s.Aggregate(key)
	s.Aggregate(key)
	s.Aggregate(key2)

	if s.Len() != 2 {
		t.Error("invalid buffer length")
	}

	s.Flush(key)

	if s.Len() != 1 {
		t.Error("invalid buffer length after Flush()")
	}

	s.FlushAll(time.Second)

	if s.Len() != 0 {
		t.Error("invalid buffer length after FlushAll()")
	}
}
