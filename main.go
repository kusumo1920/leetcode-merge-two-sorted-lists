package main

import "fmt"

func main() {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}
	output := mergeTwoListsSolution2(list1, list2)
	traverseListNode(output)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoListsSolution1(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	pointer1 := list1
	pointer2 := list2
	resultHead := ListNode{}
	pointerResult := &resultHead

	for pointer1 != nil && pointer2 != nil {
		if pointer1.Val <= pointer2.Val {
			pointerResult.Val = pointer1.Val
			pointer1 = pointer1.Next
		} else {
			pointerResult.Val = pointer2.Val
			pointer2 = pointer2.Next
		}

		if pointer1 != nil && pointer2 != nil {
			pointerResult.Next = &ListNode{}
			pointerResult = pointerResult.Next
		}
	}

	if pointer1 == nil {
		pointerResult.Next = pointer2
	} else {
		pointerResult.Next = pointer1
	}

	return &resultHead
}

func mergeTwoListsSolution2(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val <= list2.Val {
		list1.Next = mergeTwoListsSolution2(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoListsSolution2(list1, list2.Next)
		return list2
	}
}

func traverseListNode(node *ListNode) {
	fmt.Print("[")
	for node != nil {
		fmt.Printf("%d", node.Val)
		if node.Next != nil {
			fmt.Print(" ")
		}

		node = node.Next
	}
	fmt.Println("]")
}
