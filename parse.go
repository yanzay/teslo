package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func parse(raw string) (*html.Node, error) {
	doc, err := html.Parse(strings.NewReader(raw))
	if err != nil {
		return nil, err
	}
	fmt.Println(doc.Parent)
	fmt.Println(doc.FirstChild)
	fmt.Println(doc.LastChild)
	fmt.Println(doc.PrevSibling)
	fmt.Println(doc.NextSibling)
	fmt.Println(doc.Type)
	fmt.Println(doc.DataAtom)
	fmt.Println(doc.Data)
	fmt.Println(doc.Namespace)
	fmt.Println(doc.Attr)
	return nil, nil
}
