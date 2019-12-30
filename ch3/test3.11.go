package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	println(comma3(-100000000000.11111))
}

func comma3(s float64) string {
	var r bytes.Buffer
	s1 := strconv.FormatFloat(s, 'f', -1, 64)
	l := len(s1)
	decimal := false
	for i := l; i > 0; i-- {
		_s := s1[l-i : l-i+1]
		_sFirst := s1[:1]
		if _s == "." {
			decimal = true
		}
		if !decimal && i%3 == 0 && i != l && ((i > l-2 && _sFirst != "-") || i <= l-2) {
			r.WriteString(",")
		}
		fmt.Fprintf(&r, "%s", s1[l-i:l-i+1])
	}
	return r.String()
}
