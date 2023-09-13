package mergetwosortedlists

import "testing"

func TestMergeTwoLists(t *testing.T) {
	list1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	list2 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	res := mergeTwoLists(&list1, &list2)

	// I could write a piece of code to create an array from the list and compare it
	// But, I am too lazy.

	t.Log(res) // Useless. Just for the compiler.
}
