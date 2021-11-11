package main

import (
	"fmt"
	"sort"
)

// Find returns the smallest index i at which x == a[i],
// or len(a) if there is no such index.
func Find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}

// Contains tells whether a contains x.
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
func main(){


	a := []string{"A", "C", "C"}

	fmt.Println(sort.SearchStrings(a, "A")) // 0
	fmt.Println(sort.SearchStrings(a, "B")) // 1
	fmt.Println(sort.SearchStrings(a, "C")) // 1
	fmt.Println(sort.SearchStrings(a, "D")) // 3

	x := "C"

	i := sort.Search(len(a), func(i int) bool { return x <= a[i] })
	if i < len(a) && a[i] == x {
		fmt.Printf("Found %s at index %d in %v.\n", x, i, a)
	} else {
		fmt.Printf("Did not find %s in %v.\n", x, a)
	}
	// Output: Found C at index 1 in [A C C].
}