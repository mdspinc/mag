package monitor

type (
	// Campaign info stored in monitor.
	storedCamp struct {
		frequencyKeyPresent int
		impressions         int
	}

	// Store is a struct for storing monitor data.
	// Map key is a capmaign ID.
	Store map[string]*storedCamp

	Storage interface {
		SumFKP() int
		Diff() int
	}
)

// NewStore initializes new Store struct.
func NewStore() Store {
	return make(Store)
}

func newStoredCamp(fkp, imp int) *storedCamp {
	return &storedCamp{
		frequencyKeyPresent: fkp,
		impressions:         imp,
	}
}

// SumFKP returns sum of all error of type FREQUNCY_KEY_PRESENT
// by all campaigns.
func (s Store) SumFKP() (r int) {
	for _, v := range s {
		r += v.frequencyKeyPresent
	}
	return r
}

// Diff returns difference between numbers of FREQUNCY_KEY_PRESENT errors
// between two Stores compared by keys which presented in both maps.
func (s Store) Diff(d Store) (r int) {
	for k, vs := range s {
		if vd, ok := d[k]; ok {
			r += vs.frequencyKeyPresent - vd.frequencyKeyPresent
		}
	}

	return r
}
