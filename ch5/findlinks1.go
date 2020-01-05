package main

import (
	"fmt"
	"gopl.io/ch5/example"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range example.Visit(nil, doc) {
		fmt.Println(link)
	}
}
