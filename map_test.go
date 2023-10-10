package universalmap

import (
	"fmt"
	"testing"
)

func TestMap_Put(t *testing.T) {
	hash := func(k []int) Hashcode {
		var s = 0
		for _, e := range k {
			s += e
		}
		return Hashcode(s)
	}
	equal := func(k1 []int, k2 []int) bool {
		if len(k1) != len(k2) {
			return false
		} else {
			for i, e := range k1 {
				if e != k2[i] {
					return false
				}
			}
			return true
		}
	}
	m := New[[]int, int](hash, equal)

	m.Put([]int{1}, 10)
	m.Put([]int{2}, 2)
	m.Put([]int{3}, 3)
	m.Put([]int{1}, 1)

	assert(m.Len() == 3, m.Len())

	m.ForEach(func(k []int, v int) bool {
		fmt.Println(k, v)
		assert(k[0] == v, k, v)
		return true
	})
}

func TestMap_Put2(t *testing.T) {
	hash := func(k []int) Hashcode {
		return 0 // same slot for test
	}
	equal := func(k1 []int, k2 []int) bool {
		if len(k1) != len(k2) {
			return false
		} else {
			for i, e := range k1 {
				if e != k2[i] {
					return false
				}
			}
			return true
		}
	}
	m := New[[]int, int](hash, equal)

	m.Put([]int{1}, 10)
	m.Put([]int{2}, 2)
	m.Put([]int{3}, 3)
	m.Put([]int{1}, 1)

	assert(m.Len() == 3, m.Len())

	m.ForEach(func(k []int, v int) bool {
		fmt.Println(k, v)
		assert(k[0] == v, k, v)
		return true
	})
}

func assert(b bool, massages ...any) {
	if !b {
		panic(fmt.Sprint(massages...))
	}
}

func TestMap_Delete(t *testing.T) {
	hash := func(k []int) Hashcode {
		var s = 0
		for _, e := range k {
			s += e
		}
		return Hashcode(s)
	}
	equal := func(k1 []int, k2 []int) bool {
		if len(k1) != len(k2) {
			return false
		} else {
			for i, e := range k1 {
				if e != k2[i] {
					return false
				}
			}
			return true
		}
	}
	m := New[[]int, int](hash, equal)

	m.Put([]int{1}, 10)

	m.Delete([]int{1})
	m.Delete([]int{1})

	assert(m.Len() == 0, m.Len())

	m.Put([]int{2}, 2)
	m.Put([]int{3}, 3)
	m.Put([]int{1}, 1)

	m.Delete([]int{2})

	assert(m.Len() == 2, m.Len())

	m.ForEach(func(k []int, v int) bool {
		fmt.Println(k, v)
		assert(k[0] == v, k, v)
		return true
	})
}

func TestMap_Delete2(t *testing.T) {
	hash := func(k []int) Hashcode {
		return 0
	}
	equal := func(k1 []int, k2 []int) bool {
		if len(k1) != len(k2) {
			return false
		} else {
			for i, e := range k1 {
				if e != k2[i] {
					return false
				}
			}
			return true
		}
	}
	m := New[[]int, int](hash, equal)

	m.Put([]int{1}, 10)

	m.Delete([]int{1})
	m.Delete([]int{1})

	assert(m.Len() == 0, m.Len())

	m.Put([]int{2}, 2)
	m.Put([]int{3}, 3)
	m.Put([]int{1}, 1)

	m.Delete([]int{2})

	assert(m.Len() == 2, m.Len())

	m.ForEach(func(k []int, v int) bool {
		fmt.Println(k, v)
		assert(k[0] == v, k, v)
		return true
	})
}
