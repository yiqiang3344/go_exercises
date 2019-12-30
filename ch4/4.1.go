// Server2 is a minimal "echo" and counter server.
package main

import (
	"crypto/sha256"
	"fmt"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}
func main() {
	v1 := sha256.Sum256([]byte("X"))
	v2 := sha256.Sum256([]byte("x"))
	fmt.Println(diffBitCount(v1, v2))
}

func diffBitCount(s [sha256.Size]byte, s1 [sha256.Size]byte) int {
	sum := 0
	for i, v := range s {
		v1 := s1[i]
		sum += int(pc[byte(v^v1)])
	}
	return sum
}
