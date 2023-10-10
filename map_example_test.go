package universalmap

import (
	"testing"
)

func TestMap(t *testing.T) {
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

	m.Put(nil, -1)
	m.Put([]int{}, 0)

	m.Put([]int{1}, 1)
	m.Put([]int{2}, 2)
	m.Put([]int{3}, 3)

	m.Put([]int{1, 1}, 11)
	m.Put([]int{1, 1, 1}, 111)

	// m.Len() = 6
	println("m.Len() =", m.Len())

	// {[3]:3, [1 1 1]:111, []:0, [1]:1, [2]:2, [1 1]:11}
	println(m.String())
}
