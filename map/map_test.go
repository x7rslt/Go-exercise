package map_test

import (
	"fmt"
	"sort"
	"testing"
)

func TestMap(t *testing.T){
	//Create a new map
	var m map[string]int                // nil map of string-int pairs

	m1 := make(map[string]float64)      // Empty map of string-float64 pairs
	m2 := make(map[string]float64, 100) // Preallocate room for 100 entries

	m3 := map[string]float64{           // Map literal
		"e":  2.71828,
		"pi": 3.1416,
	}
	fmt.Printf("m:%v,%T;\n1:%v,%T;\nm2:%v,%T;\nm3:%v,%T;\n", m, m,m1,m1, m2, m2, m3,m3)
	fmt.Println(len(m),len(m1),len(m2),len(m3))                // Size of map: 0 0 0 2
}

func TestMapEdit(t *testing.T){
	//Add, update, get and delete keys/values
	m := make(map[string]float64)

	m["pi"] = 3.14             // Add a new key-value pair
	m["pi"] = 3.1416           // Update value
	fmt.Println(m)             // Print map: "map[pi:3.1416]"

	v := m["pi"]
	fmt.Println(v)// Get value: v == 3.1416
	v = m["pie"]               // Not found: v == 0 (zero value)
	fmt.Println(v)

	_, found := m["pi"]        // found == true
	_, found = m["pie"]        // found == false
	fmt.Println(found)

	if x, found := m["pi"]; found {
		fmt.Println(x)
	}                           // Prints "3.1416"

	delete(m, "pi")             // Delete a key-value pair
	fmt.Println(m)              // Print map: "map[]"
}

//Find a key in map
func TestMapKey(t *testing.T){
	m := map[string]float64{"pi": 3.14}
	v, found := m["pi"] // v == 3.14  found == true
	v, found = m["pie"] // v == 0.0   found == false
	_, found = m["pi"]  // found == true
	fmt.Println(v,found)

	m = map[string]float64{"pi": 3.14}
	if v, found := m["pi"]; found {
		fmt.Println(v)
	}
	// Output: 3.14

	m = map[string]float64{"pi": 3.14}

	v = m["pi"] // v == 3.14
	fmt.Println(v)
	v = m["pie"] // v == 0.0 (zero value)
	fmt.Println(v)
}

//Sort a map by key or value
func TestMapSort(t *testing.T){
	m := map[string]int{"Alice": 23, "Eve": 2, "Bob": 25}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	fmt.Printf("%T,%v\n",keys,keys)
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, m[k])
	}
}