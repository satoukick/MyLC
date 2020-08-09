package main

// Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func getCloneNode(visited map[*Node]*Node, node *Node) *Node {
	if node == nil {
		return nil
	}
	if _, ok := visited[node]; ok {
		return visited[node]
	}

	newNode := &Node{
		Val: node.Val,
	}
	visited[node] = newNode
	return newNode
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	visited := make(map[*Node]*Node)
	newNode := &Node{
		Val: head.Val,
	}
	oldNode := head
	visited[oldNode] = newNode
	for oldNode != nil {
		newNode.Random = getCloneNode(visited, oldNode.Random)
		newNode.Next = getCloneNode(visited, oldNode.Next)

		oldNode = oldNode.Next
		newNode = newNode.Next
	}

	return visited[head]
}
