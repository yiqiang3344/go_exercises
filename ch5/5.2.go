package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

//编写函数，记录在HTML树中出现的同名元素的次数。
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	for k, v := range Visit1(make(map[string]int), doc) {
		fmt.Printf("%s:%d\n", k, v)
	}
}

func Visit1(count map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		count[n.Data]++
	}
	if n.FirstChild != nil {
		count = VisitNextSibling1(count, n.FirstChild)
	}
	return count
}

func VisitNextSibling1(count map[string]int, n *html.Node) map[string]int {
	count = Visit1(count, n)
	if n.NextSibling != nil {
		count = VisitNextSibling1(count, n.NextSibling)
	}
	return count
}
