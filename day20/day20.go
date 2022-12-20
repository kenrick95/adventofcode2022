package main

import (
	"strconv"
	"utils"
)

type Node struct {
	originalIndex int
	value         int
	next          *Node
	prev          *Node
}

func main() {
	nodeMap := map[int]*Node{}
	numLength := 0
	i := 0
	lines := utils.ReadFileToLines("day20.in")

	// Part 1
	// decryptionKey := 1
	// mixCount := 1
	// Part 2
	decryptionKey := 811589153
	mixCount := 10

	var pointerToZero *Node = nil
	for _, line := range lines {
		if line == "" {
			continue
		}
		num, _ := strconv.Atoi(line)
		node := &Node{
			originalIndex: i,
			value:         num * decryptionKey,
			next:          nil,
			prev:          nil,
		}
		nodeMap[i] = node

		{
			prevNode, prevNodeExist := nodeMap[i-1]
			if prevNodeExist && prevNode != nil {
				prevNode.next = node
				node.prev = prevNode
			}
		}

		if node.value == 0 {
			pointerToZero = node
		}

		i += 1
		numLength = i
	}
	// To make the list circular
	firstNode := nodeMap[0]
	lastNode := nodeMap[numLength-1]
	firstNode.prev = lastNode
	lastNode.next = firstNode

	debug := func() {

		head := firstNode
		println("debug", head.originalIndex, head.value)
		for i := 0; i < numLength; i++ {
			print(head.value)
			if i != numLength-1 {
				print(", ")
			}
			head = head.next
		}
		println()

	}

	insertNodeAfter := func(node *Node, targetNode *Node) {
		if targetNode.next == node {
			// no change
			return
		}
		if node.next == targetNode {
			nodePrev := node.prev
			targetNodeNext := targetNode.next

			node.prev = targetNode
			node.next = targetNodeNext

			targetNode.prev = nodePrev
			targetNode.next = node

			targetNodeNext.prev = node
			nodePrev.next = targetNode
		} else {
			nodePrev := node.prev
			nodeNext := node.next
			nodePrev.next = nodeNext
			nodeNext.prev = nodePrev

			targetNodeNext := targetNode.next
			targetNodeNext.prev = node
			targetNode.next = node
			node.next = targetNodeNext
			node.prev = targetNode
		}
	}

	mix := func() {
		for i := 0; i < numLength; i++ {
			// println()
			// println()
			// println("i", i)
			// debug()
			node := nodeMap[i]
			moveCount := utils.Abs(node.value) % (numLength - 1)

			if moveCount == 0 {
				// nothing
			} else {
				targetNode := node.prev
				if node.value > 0 {
					targetNode = targetNode.next
				}
				for c := 0; c < moveCount; c++ {
					if node.value > 0 {
						targetNode = targetNode.next
					} else {
						targetNode = targetNode.prev
					}
					if targetNode == node {
						// Need to treat as if we removed node from the list
						c -= 1
					}
				}
				// println("insertNode", node.originalIndex, "(", node.value, ")", "after", targetNode.originalIndex, "(", targetNode.value, ")")
				insertNodeAfter(node, targetNode)
			}

			// debug()
		}
	}

	for count := 0; count < mixCount; count++ {
		mix()
		debug()
	}

	ans := 0
	{
		node := pointerToZero
		for i := 0; i <= 3005; i++ {
			if i == 1000 || i == 2000 || i == 3000 {
				println("i", i, node.value)
				ans += node.value
			}
			node = node.next
		}
	}
	println("ans", ans)

}
