package sort_test

import (
	"fmt"
	"sort"
	"testing"
)

//Sort a slice of ints, float64s or strings
func TestSortInt(t *testing.T){
	s:= []int{4,2,3,1}
	sort.Ints(s)
	fmt.Println(s) // [1 2 3 4]
}

//Sort with custom comparator
func TestSortCustom(t *testing.T){
	family := []struct {
		Name string
		Age  int
	}{
		{"Alice", 23},
		{"David", 2},
		{"Eve", 2},
		{"Bob", 25},
	}

	// Sort by age, keeping original order or equal elements.
	sort.SliceStable(family, func(i, j int) bool {
		return family[i].Age < family[j].Age
	})
	fmt.Println(family) // [{David 2} {Eve 2} {Alice 23} {Bob 25}]
}

//Sort custom data structures
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

type Person struct {
	Name string
	Age  int
}

// ByAge implements sort.Interface based on the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func TestCustomStructures(t *testing.T){
	family := []Person{
		{"Alice", 23},
		{"Eve", 2},
		{"Bob", 25},
	}
	sort.Sort(ByAge(family))
	fmt.Println(family) // [{Eve 2} {Alice 23} {Bob 25}]
}

//Bonus: Sort a map by key or value
func TestSortKeyVule(t *testing.T){
	m := map[string]int{"Alice": 2, "Cecil": 1, "Bob": 3}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, m[k])
	}
	// Output:
	// Alice 2
	// Bob 3
	// Cecil 1
}