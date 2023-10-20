package syncx

import "sync"

type Map[k comparable, v any] struct {
	m sync.Map
}

func (m *Map[K, V]) Delete(key K) {
	m.m.Delete(key)
}

/*func (m *Map[K, V]) Get(key K) V {
	v, ok := m.m.Load(key)

	if !ok {
		panic("Access to undefined index.")
	}

	return v.(V)
}*/

func (m *Map[K, V]) Load(key K) (V, bool) {
	v, ok := m.m.Load(key)

	if !ok {
		return v.(V), ok
	}

	return v.(V), ok
}

func (m *Map[K, V]) Range(f func(key K, value V) bool) {
	m.m.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

func (m *Map[K, V]) Store(key K, value V) {
	m.m.Store(key, value)
}

func (m *Map[K, V]) Len() int {
	var output = 0

	m.Range(func(_ K, _ V) bool {
		output += 1
		return true
	})

	return output
}
