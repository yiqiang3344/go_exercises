// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
)

func main() {
	var s = []int{1, 2, 3, 3, 4, 2, 3, 3, 4, 5, 5}
	s = uniq(s)
	fmt.Printf("%x", s)
}

func uniq(s []int) []int {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			copy(s[i:len(s)-1], s[i+1:])
			s = s[:len(s)-1]
			i += 1
		}
	}
	return s
}
