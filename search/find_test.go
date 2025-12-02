package search

import "fmt"

func ExampleFind() {
	l1 := []int{1, 4, 3, 6, 21, 10, 38}

	fmt.Println(Find(l1, 10))
	fmt.Println(Find(l1, 6))
	fmt.Println(Find(l1, 3))
	fmt.Println(Find(l1, 400))

	// Output:
	// 5
	// 3
	// 2
	// -1
	}   
