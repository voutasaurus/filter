package filter

type Predicate func(int) bool

// Split returns in two separate slices: the ints of list that satisfy the
// predicate and the ints that do not. The second slice is not kept in a
// stable order relative to the original slice.
func Split(list []int, p Predicate) (in []int, not []int) {
	pivot := 0
	for i := 0; i < len(list); i++ {
		if p(list[i]) {
			list[i], list[pivot] = list[pivot], list[i]
			pivot++
		}
	}
	return list[:pivot], list[pivot:]
}

// Filter returns a slice of the ints of list that satisfy the predicate.
func Filter(list []int, p Predicate) []int {
	pivot := 0
	for i := 0; i < len(list); i++ {
		if p(list[i]) {
			list[pivot] = list[i]
			pivot++
		}
	}
	return list[:pivot]
}
