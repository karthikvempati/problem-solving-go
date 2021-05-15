/*
Suppose you're given a binary tree represented as an array. For example, [3,6,2,9,-1,10] represents the following binary tree (where -1 is a non-existent node):

enter image description here

Write a function that determines whether the left or right branch of the tree is larger. The size of each branch is the sum of the node values. The function should return the string "Right" if the right side is larger and "Left" if the left side is larger. If the tree has 0 nodes or if the size of the branches are equal, return the empty string.

Example Input:

[3,6,2,9,-1,10]
Example Output:

Left
*/

package solution

import "fmt"

func Solution(arr []int64) string {
	// Type your solution here
	length := int64(len(arr))
	if length == 0 || length == 1 {
		return ""
	} else if length == 2 {
		return "Left"
	}
	hl := HeightOfTree(arr[1:])
	hr := HeightOfTree(arr[2:])
	if hl > hr {
		return "Left"
	} else if hr > hl {
		return "Right"
	} else {
		return ""
	}
}

func HeightOfTree(arr []int64) int64 {
	length := int64(len(arr))
	if length == 0 {
		return 0
	} else if length == 1 {
		return 1
	}
	l := LeftChild(0, arr)
	r := RightChild(0, arr)

	if l == -1 && r == -1 {
		return 1
	} else if l == -1 && r < length {
		return 1 + HeightOfTree(arr[r:])
	} else if r == -1 && l < length {
		return 1 + HeightOfTree(arr[l:])
	} else if l < length && r < length {
		fmt.Println(l, r, arr)
		return 1 + Max(HeightOfTree(arr[r:]), HeightOfTree(arr[l:]))
	}

	return 1
}

func Max(i, j int64) int64 {
	if i >= j {
		return i
	}

	return j
}

func LeftChild(i int64, arr []int64) int64 {
	var left int64
	left = 2*i + 1
	length := int64(len(arr))
	if i < 0 || i > length || left > length || arr[left] == -1 {
		return -1
	}
	return left
}

func RightChild(i int64, arr []int64) int64 {
	var right int64
	length := int64(len(arr))
	right = 2*i + 2
	if i < 0 || i > length || right > length || arr[right] == -1 {
		return -1
	}
	return right
}
