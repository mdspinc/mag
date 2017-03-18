package monitor

type Store map[string]*storedCamp

func NewStore() Store {
	return make(Store)
}

func (s Store) SumFKP() (r int) {
	for _, v := range s {
		r += v.frequencyKeyPresent
	}
	return r
}
