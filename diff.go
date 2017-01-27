package main

import "golang.org/x/net/html"

/*
Possible actions:
- appendChild
  param: id of parent node, child content
- insertAt
  param: id of parent node, position to insert, node content
- repalceNode
  param: id of node, new content
- deleteNode
  param: id of node
*/

type Diff struct {
	action  string
	content string
}

const (
	ActionNop         = ""
	ActionAppendChild = "appendChild"
	ActionReplaceNode = "replaceNode"
)

func diff(initial string, updated string) *Diff {
	initialNode, err := parse(initial)
	updatedNode, err := parse(updated)
	return &Diff{action: ActionAppendChild, content: "<li>third</li>"}
}

func nodeDiff(initial *html.Node, updated *html.Node) {
	if initial.Type != updated.Type {
		return &Diff{action: ActionReplaceNode}
	}
	if initial.Type == html.TextNode {
		if initial.Data == updated.Data {
			return &Diff{action: ActionNop}
		}
		return &Diff{action: ActionReplaceNode}
	}
  if initial.FirstChild != nil && updated.FirstChild != nil {
		return diff(inital.FirstChild, updated.FirstChild)
}
