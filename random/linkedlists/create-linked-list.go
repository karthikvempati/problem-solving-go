package linkedlists

import (
	"fmt"

	"github.com/karthikvempati/data-structures-go/linkedlist"
)

func CreateLinkedList(nums []int) linkedlist.LinkedList {
	l := linkedlist.LinkedList{}

	for i := len(nums) - 1; i >= 0; i-- {
		l.Prepend(nums[i])
	}

	return l
}

func PrintLinkedList(l linkedlist.LinkedList) {
	node := *l.Head
	for ok := true; ok; ok = (node.Next != nil) {
		fmt.Println(node.Val)
		node = *node.Next
	}
}
