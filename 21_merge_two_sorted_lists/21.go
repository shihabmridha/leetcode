package mergetwosortedlists

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	list := []int{}

	for list1 != nil {
		list = append(list, list1.Val)
		list1 = list1.Next
	}

	for list2 != nil {
		list = append(list, list2.Val)
		list2 = list2.Next
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})

	var head *ListNode = nil
	var result *ListNode = nil

	for _, e := range list {
		if result == nil {
			head = &ListNode{Val: e}
			result = head
		} else {
			result.Next = &ListNode{Val: e}
			result = result.Next
		}
	}

	return head
}
