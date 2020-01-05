package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

//编写函数输出所有text结点的内容。注意不要访问<script>和<style>元素,因为这些元素对浏览者是不可见的。
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, v := range Visit2(nil, doc) {
		fmt.Printf("%#v\n", v)
	}
}

type Element struct {
	Type    string
	Content string
}

func Visit2(lists []Element, n *html.Node) []Element {
	var element Element
	if n.Type == html.ElementNode && n.Data != "script" && n.Data != "style" && n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
		element = Element{n.Data, n.FirstChild.Data}
		lists = append(lists, element)
	}

	if n.FirstChild != nil {
		lists = VisitNextSibling2(lists, n.FirstChild)
	}
	return lists
}

func VisitNextSibling2(lists []Element, n *html.Node) []Element {
	lists = Visit2(lists, n)
	if n.NextSibling != nil {
		lists = VisitNextSibling2(lists, n.NextSibling)
	}
	return lists
}
