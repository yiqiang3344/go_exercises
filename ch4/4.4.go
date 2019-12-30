// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"math"
)

func main() {
	var s = []int{1, 2, 3, 4, 5}
	rotate1(s)
	fmt.Printf("%x", s)
}

func rotate(s []int) {
	m := len(s) / 2
	var r = make([]int, m)
	copy(r, s[m:])
	copy(s[m:], s[:m])
	copy(s[:m], r)
}

func rotate1(s []int) {
	m := int(math.Ceil(float64(len(s)) / 2))
	n := m
	if m*2 != len(s) {
		n -= 1
	}
	for i := 0; i < n; i++ {
		s[i], s[i+m] = s[i+m], s[i]
	}
}
