package agg

type (
	Type byte

	Aggregator interface {
		//Collects incoming messages.
		Aggregate(interface{})

		//Flushes collected messages to sender.
		Flush(string)

		//Returns number of collected messages.
		Count(string) int
	}
)

const (
	AGGTYPE_STRING Type = iota
)
