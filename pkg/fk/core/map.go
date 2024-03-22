package core

import "math/rand"

type RandMap struct {
	rand *rand.Rand
}

func NewRandMap(rand *rand.Rand) *RandMap {
	return &RandMap{
		rand,
	}
}

// returns a random key and value from the map
func (r *RandMap) KeyValue(m map[any]any) (any, any) {
	return GetRandomKeyValue(r, m)
}

// returns a random key and value which is an slice from the map
func (r *RandMap) KeySliceValue(m map[any][]any) (any, any) {
	return GetRandomKeyValue(r, m)
}

// returns a random key and value from the map
// keys and values can be any types
func GetRandomKeyValue[K comparable, V any](r *RandMap, m map[K]V) (K, V) {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	key := keys[r.rand.Intn(len(keys))]
	value := m[key]
	return key, value
}
