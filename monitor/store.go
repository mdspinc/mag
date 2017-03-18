package monitor

// Store is a struct for storing monitor data.
type Store map[string]*storedCamp

// NewStore initializes new Store struct.
func NewStore() Store {
	return make(Store)
}

// SumFKP returns sum of all error of type FREQUNCY_KEY_PRESENT
// by all campaigns.
func (s Store) SumFKP() (r int) {
	for _, v := range s {
		r += v.frequencyKeyPresent
	}
	return r
}
