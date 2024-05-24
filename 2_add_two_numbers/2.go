package addtwonumbers

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	var res *ListNode = nil
	var last *ListNode = nil

	for {
		val := 0
		sum := (l1.Val + l2.Val) + carry

		if sum > 9 {
			val = sum % 10
			carry = sum / 10
		} else {
			val = sum
			carry = 0
		}

		if res == nil {
			res = &ListNode{
				Val: val,
			}
			last = res
		} else {
			last.Next = &ListNode{
				Val: val,
			}
			last = last.Next
		}

		if l1.Next == nil && l2.Next == nil && carry == 0 {
			break
		}

		if l1.Next != nil {
			l1 = l1.Next
		} else {
			l1.Val = 0
		}

		if l2.Next != nil {
			l2 = l2.Next
		} else {
			l2.Val = 0
		}
	}

	return res
}
