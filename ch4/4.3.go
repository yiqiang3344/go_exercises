// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
)

func main() {
	var s = [4]int{1, 2, 3, 4}
	reverse(&s)
	fmt.Println(s)
}

// reverse reverses a slice of ints in place.
func reverse(s *[4]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
