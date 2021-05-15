package linkedlists

import "fmt"

func ReverseLinkedList() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	l := CreateLinkedList(nums)
	prev := l.Head
	prev = nil
	current := l.Head
	next := l.Head.Next
	//Reverse Singly liked list
	for ok := true; ok; ok = next != nil {
		fmt.Println(prev, current, next)
		current.Next = prev
		prev = current
		current = next
		next = current.Next
		if next == nil {
			current.Next = prev
		}
		fmt.Println(prev, current, next, "\n")
		// fmt.Println(prev, current, next)
		// current.Next = prev
		// prev = current
		// current = next
		// next = next.Next
		// fmt.Println(prev, current, next, "\n")
	}

	l.Head = current

	PrintLinkedList(l)
}
