// @file:        node_structures.go
// @version:     1.0
// @author:      aihaofeng
// @date:        2017.10.17
// @go version:  1.9
// @brief:       Interface test.

package main

import "fmt"

type Node struct {
	le   *Node
	data interface{}
	ri   *Node
}

func NewNode(left, right *Node) *Node {
	return &Node{left, nil, right}
}

func (n *Node) SetData(data interface{}) {
	n.data = data
}

func main() {
	root := NewNode(nil, nil)
	root.SetData("root node")

	// make child (leaf) nodes:
	a := NewNode(nil, nil)
	a.SetData("left node")
	b := NewNode(nil, nil)
	b.SetData("right node")

	root.le = a
	root.ri = b

	// Output: &{0x125275f0 root node 0x125275e0}
	fmt.Printf("%v\n", root)
}
