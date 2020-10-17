package heapsort

// Sort a slice of int-s in place, ascending order
func SortAsc(a []int) {
	heapify(a)
	for i := len(a) - 1; i > 0; i-- {
		swap(0, i, a)
		down(0, a[:i])
	}
}

// Since the sort is ascending, build a max-heap
func heapify(a []int) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		down(i, a)
	}
}

func down(i int, a []int) {
	var swapWith int

	for ; i <= len(a)/2-1; {
		left := i*2 + 1

		swapWith = left
		if left < len(a)-1 {
			right := left + 1
			if a[left] < a[right] {
				swapWith = right
			}
		}

		if a[i] < a[swapWith] {
			swap(i, swapWith, a)
			i = swapWith
		} else {
			return
		}
	}
}

func swap(i, j int, a []int) {
	a[i], a[j] = a[j], a[i]
}
