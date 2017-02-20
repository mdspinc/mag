package agg

import "testing"

func testInit(batchSize int) (sa *StringAgg, key, key2 string, bs int) {
	bs = batchSize
	sa = NewStringAgg(bs)
	key, key2 = "A", "B"
	return
}

func TestNewStringAgg(t *testing.T) {
	s, key, _, batchSize := testInit(10)

	if s.batchSize != batchSize {
		t.Error("invalid batch size")
	}

	if s.Count(key) != 0 {
		t.Error("ivalid items per key")
	}

	if s.sender == nil {
		t.Error("sender is not initialized")
	}

	if s.buffer == nil {
		t.Error("buffer is not initialzed")
	}
}
