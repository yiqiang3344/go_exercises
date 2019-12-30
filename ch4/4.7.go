// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "测试2一下a哈1哈哟"
	fmt.Println(reverseByte(s))
}

func reverseByte(s string) string {
	var list []string
	for i, j := 0, 0; i < len(s); {
		_, size := utf8.DecodeRuneInString(s[i:])
		list = append(list,s[i:i+size])
		j++
		i += size
	}
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i],list[j] = list[j],list[i]
	}
	return strings.Join(list,"")
}
