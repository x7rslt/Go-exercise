package main

import (
	"crypto/sha256"
	"fmt"
)

func ncount(s1, s2 string) int {
	c1 := sha256.Sum256([]byte(s1))
	c2 := sha256.Sum256([]byte(s2))
	var count int
	for i, v := range c1 {
		if v != c2[i] {
			count++
		}
	}
	fmt.Printf("c1 Sha256:%x\nc2 Sha256:%x\n", c1, c2)
	return count
}

func main() {
	s1 := "x"
	s2 := "X"
	fmt.Println(ncount(s1, s2))

}
