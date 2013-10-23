package filter

import "fmt"

type Predicate func(int) bool

func Filter(list []int, p Predicate) ([]int, []int) {
	pivot := 0
	for i := 0; i < len(list); i++ {
		if p(list[i]) {
			// swap with pivot
			list[i], list[pivot] = list[pivot], list[i]
			pivot++
		} // otherwise leave it 
	}
	return list[:pivot], list[pivot:]
}

func Filter_destructive(list []int, p Predicate) ([]int) {
	pivot := 0
	for i := 0; i < len(list); i++ {
		if p(list[i]) {
			// swap with pivot
			list[pivot] = list[i]
			pivot++
		} // otherwise leave it 
	}
	return list[:pivot]
}

func even(i int) bool {
	return i % 2 == 0
}

func truth(i int) bool {
	return true
}

// Write a proper test
func FilterTest() {
	c := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a, b := Filter(c, even)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(even(1))
	
	d := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	e := Filter_destructive(d, even)
	fmt.Println(e)
}