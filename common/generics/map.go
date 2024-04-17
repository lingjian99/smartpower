package generics

import "sync"

type MutexMap[K comparable, V any] struct {
	l sync.RWMutex
	v map[K]V
}

func MakeMutexMap[K comparable, V any]() *MutexMap[K, V] {
	return &MutexMap[K, V]{
		v: make(map[K]V),
	}
}

func (m *MutexMap[K, V]) Get(k K) (V, bool) {
	m.l.RLock()
	defer m.l.RUnlock()
	v, ok := m.v[k]
	return v, ok
}

func (m *MutexMap[K, V]) Set(k K, v V) {
	m.l.Lock()
	defer m.l.Unlock()
	m.v[k] = v
}

func (m *MutexMap[K, V]) Del(k K) {
	m.l.Lock()
	defer m.l.Unlock()
	delete(m.v, k)
}

func (m *MutexMap[K, V]) Keys() []K {
	m.l.Lock()
	defer m.l.Unlock()
	var keys []K
	for k := range m.v {
		keys = append(keys, k)
	}
	return keys
}

func (m *MutexMap[K, V]) Range(fn func(k K, v V) bool) {
	m.l.Lock()
	defer m.l.Unlock()
	for k, v := range m.v {
		if !fn(k, v) {
			break
		}
	}
}
func (m *MutexMap[K, V]) Len() int {
	return len(m.v)
}
