package tree

import "fmt"

type Node struct {
	data  interface{}
	left  *Node
	right *Node
}

func NewNode(data interface{}) *Node {
	return &Node{data: data}
}

func (this *Node) String() string {
	return fmt.Sprintf("v:%+v, left:%+v, right:%+v", this.data, this.left, this.right)
	// return fmt.Sprintf("v:%+v, left:%+v, right:%+v", this.data, this.left, this.right)
}

// func main() {
// 	a := Node{"a", nil, nil}
// 	b := Node{"b", nil, nil}
// 	c := Node{"c", nil, nil}
// 	a.left = &b
// 	b.right = &c
// 	fmt.Println(a.String())
// }
