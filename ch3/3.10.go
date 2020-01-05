package main

import (
	"bytes"
	"fmt"
)

func main() {
	//println(comma("100000000000"))
	//println(comma1("100000000000"))
	println(comma2("100000000000"))
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// comma inserts commas in a non-negative decimal integer string.
func comma1(s string) string {
	var r string
	l := len(s)
	for i := l; i > 0; i-- {
		if i%3 == 0 && i != l {
			r += ","
		}
		r += s[l-i : l-i+1]
	}
	return r
}

func comma2(s string) string {
	var r bytes.Buffer
	l := len(s)
	for i := l; i > 0; i-- {
		if i%3 == 0 && i != l {
			r.WriteString(",")
		}
		fmt.Fprintf(&r, "%s", s[l-i:l-i+1])
	}
	return r.String()
}
