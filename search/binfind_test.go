package search

import "fmt"

func ExampleFindSorted() {
	l1 := []int{1, 3, 4, 6, 10, 21, 38}
	l2 := []int{}

	fmt.Println(FindSorted(l1, 10))
	fmt.Println(FindSorted(l1, 38))
	fmt.Println(FindSorted(l1, 3))
	fmt.Println(FindSorted(l1, 15))
	fmt.Println(FindSorted(l2, 2))
	fmt.Println(FindSorted(l1, 2))

	// Output:
	// 4
	// 6
	// 1
	// -1
	// -1
	// -1

}
