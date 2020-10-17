package heapsort

import (
	"math/rand"
	"testing"
	"time"
)

func TestHeapify(t *testing.T) {
	a, seed := genTestArray(8)
	was := make([]int, len(a))
	copy(was, a)
	heapify(a)
	if !isMaxHeap(a) {
		t.Errorf("heapify() didn't produce a heap for %v: got %v (seed: %d)", was, a, seed)
	}
}

func TestSortAsc(t *testing.T) {
	a, seed := genTestArray(12)
	was := make([]int, len(a))
	copy(was, a)
	SortAsc(a)
	if !isSortedAsc(a) {
		t.Errorf("SortAsc() didn't produce a sorted array for %v: got %v (seed: %d)", was, a, seed)
	}
}

func genTestArray(n int) ([]int, int64) {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = i
	}
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	rand.Shuffle(n, func(i, j int) { a[i], a[j] = a[j], a[i] })
	return a, seed
}

func isSortedAsc(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

func isMaxHeap(a []int) bool {
	for i := 0; i <= len(a)/2-1; i++ {
		left := i*2 + 1
		if a[i] < a[left] {
			return false
		}
		right := left + 1
		if right <= len(a)-1 {
			if a[i] < a[right] {
				return false
			}
		}
	}
	return true
}
