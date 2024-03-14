package copylistwithrandompointer

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

var cacheMap map[*Node]*Node

// 時間複雜 O(n), 空間複雜 O(n)
func copyRandomList(head *Node) *Node {
	cacheMap = make(map[*Node]*Node)
	return deepCopy(head)
}

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}
	if n, ok := cacheMap[node]; ok {
		return n
	}
	newNode := &Node{Val: node.Val}
	cacheMap[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode
}

func copyRandomList2(head *Node) *Node {
	cMap := make(map[*Node]*Node)
	cur := head
	// 第一次遍歷專門clone節點，藉助 Hash表把原始節點和clone節點的映射存儲起來;
	for cur != nil {
		newNode := &Node{Val: cur.Val}
		cMap[cur] = newNode
		cur = cur.Next
	}
	// 第二次專門組裝節點，照著原數據結構的樣子，把clone節點的指標組裝起來。
	newHead := cMap[head]
	for head != nil {
		node := cMap[head]
		node.Next = cMap[head.Next]
		node.Random = cMap[head.Random]
		head = head.Next
	}
	return newHead
}
