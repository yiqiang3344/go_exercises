// Server2 is a minimal "echo" and counter server.
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	var _type string
	_var := os.Args[1]
	if len(os.Args) > 2 {
		_type = os.Args[2]
	}

	switch _type {
	case "512":
		fmt.Printf("%x\n", sha512.Sum512([]byte(_var)))
	case "318":
		fmt.Printf("%x\n", sha512.Sum384([]byte(_var)))
	default:
		fmt.Printf("%x\n", sha256.Sum256([]byte(_var)))
	}
}
