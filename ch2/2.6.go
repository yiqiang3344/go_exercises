// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"gopl.io/ch2/tempconv"
)

func main() {
	fmt.Println(tempconv.CToK(tempconv.BoilingC))
}
