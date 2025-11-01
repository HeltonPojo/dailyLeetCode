package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	modifiedList([]int{1, 2, 3}, nil)
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	mapOccur := make(map[int]bool)

	for _, num := range nums {
		mapOccur[num] = true
	}

	currNode := head
	var parentNode *ListNode
	fmt.Print(mapOccur, parentNode)
	for currNode != nil {
		if mapOccur[currNode.Val] {
			if parentNode == nil {
				head = currNode.Next
				currNode = currNode.Next
			} else {
				parentNode.Next = currNode.Next
				currNode = currNode.Next
			}
		} else {
			parentNode = currNode
			currNode = currNode.Next
		}
	}

	return head
}
