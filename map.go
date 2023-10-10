package universalmap

import (
	"fmt"
	"strings"
)

type Hashcode uint

// Map uses custom hash & equal function
//
// Allow any key type such as slice
type Map[K, V any] struct {
	hash    func(K) Hashcode
	equal   func(K, K) bool
	hashMap map[Hashcode][]entry[K, V] // aux map
	size    int
}

type entry[K, V any] struct {
	k K
	v V
}

// New creates a new universal map. A user-defined hash and equal functions are required
//
// For example, a Map[[]int, int] can be created as follow
//
//	hash := func(k []int) Hashcode {
//		var s = 0
//		for _, e := range k {
//			s += e
//		}
//		return Hashcode(s)
//	}
//
//	equal := func(k1 []int, k2 []int) bool {
//		if len(k1) != len(k2) {
//			return false
//		} else {
//			for i, e := range k1 {
//				if e != k2[i] {
//					return false
//				}
//			}
//			return true
//		}
//	}
//
// m := New[[]int, int](hash, equal)
func New[K, V any](hash func(K) Hashcode, equal func(K, K) bool) *Map[K, V] {
	return &Map[K, V]{
		hash:    hash,
		equal:   equal,
		hashMap: map[Hashcode][]entry[K, V]{},
		size:    0,
	}
}

// Put sets the value for a key.
func (m *Map[K, V]) Put(key K, value V) {
	hash := m.hash(key)
	kvs := m.hashMap[hash]

	for i := range kvs {
		if m.equal(kvs[i].k, key) {
			kvs[i].v = value
			return // replace
		}
	}

	kvs = append(kvs, entry[K, V]{
		k: key,
		v: value,
	})

	m.hashMap[hash] = kvs
	m.size++
}

// Get returns the value stored in the map for a key, or nil if no value is present.
// The ok result indicates whether value was found in the map.
func (m *Map[K, V]) Get(key K) (value V, ok bool) {
	hash := m.hash(key)
	kvs := m.hashMap[hash]

	for _, kv := range kvs {
		if m.equal(kv.k, key) {
			return kv.v, true
		}
	}

	return // nil, false
}

// Delete deletes the value for a key.
func (m *Map[K, V]) Delete(key K) {
	hash := m.hash(key)
	kvs := m.hashMap[hash]

	index := -1
	for i, kv := range kvs {
		if m.equal(kv.k, key) {
			index = i
			break
		}
	}

	// not found
	if index == -1 {
		return
	}

	m.size--

	kvsLen := len(kvs)
	if kvsLen == 1 {
		// remove
		delete(m.hashMap, hash)
		return
	}

	// delete kv at index
	kvs[index] = kvs[kvsLen-1]
	kvs = kvs[:kvsLen-1]

	// save mem
	if cap(kvs) > len(kvs)*2 {
		kvs = append([]entry[K, V]{}, kvs...)
	}

	m.hashMap[hash] = kvs
	return
}

// Len returns number of values in the Map
func (m *Map[K, V]) Len() int {
	return m.size
}

// ForEach offers an iteration to the Map
func (m *Map[K, V]) ForEach(access func(K, V) (stop bool)) {
	for _, kvs := range m.hashMap {
		for _, kv := range kvs {
			if access(kv.k, kv.v) {
				return
			}
		}
	}
}

func (m *Map[K, V]) String() string {
	sb := strings.Builder{}
	sb.WriteByte('{')
	length := m.Len()
	m.ForEach(func(k K, v V) (stop bool) {
		sb.WriteString(fmt.Sprintf("%v", k))
		sb.WriteByte(':')
		sb.WriteString(fmt.Sprintf("%v", v))
		length--
		if length > 0 {
			sb.WriteString(", ")
		}
		return false
	})
	sb.WriteByte('}')
	return sb.String()
}
