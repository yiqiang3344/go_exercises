package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

//修改findlinks代码中遍历n.FirstChild链表的部分，将循环调用visit，改成递归调用。
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range Visit(nil, doc) {
		fmt.Println(link)
	}
}

func Visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	if n.FirstChild != nil {
		links = VisitNextSibling(links, n.FirstChild)
	}
	return links
}

func VisitNextSibling(links []string, n *html.Node) []string {
	links = Visit(links, n)
	if n.NextSibling != nil {
		links = VisitNextSibling(links, n.NextSibling)
	}
	return links
}
