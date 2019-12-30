// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"gopl.io/ch2/popcount"
	"time"
)

func main() {
	start := time.Now()
	fmt.Printf("1 %d\n",popcount.PopCount(255))
	secs := time.Since(start).Nanoseconds()
	fmt.Printf("1 %d\n",secs)
	start = time.Now()
	fmt.Printf("2 %d\n",popcount.PopCount1(255))
	secs = time.Since(start).Nanoseconds()
	fmt.Printf("2 %d\n",secs)
	start = time.Now()
	fmt.Printf("3 %d\n",popcount.PopCount(255))
	secs = time.Since(start).Nanoseconds()
	fmt.Printf("3 %d\n",secs)
}