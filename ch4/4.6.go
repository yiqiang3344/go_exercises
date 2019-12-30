// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "test 啊  test   1"
	fmt.Println(uniqSpace(s))
}

func uniqSpace(s string) string {
	var str string
	preIsSpace := false
	for _, v := range s {
		//fmt.Println(1,v)
		if !unicode.IsSpace(v) && preIsSpace {
			//fmt.Println(2,v)
			str += " "
			preIsSpace = false
		}
		if !unicode.IsSpace(v) {
			str += string(v)
		}else{
			preIsSpace = true
		}
	}
	return str
}
