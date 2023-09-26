package maximumdepthofnarytree

import "testing"

func TestMaxDepth(t *testing.T) {
	inp := &Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 3,
				Children: []*Node{
					{
						Val:      5,
						Children: nil,
					},
					{
						Val:      6,
						Children: nil,
					},
				},
			},
			{
				Val:      2,
				Children: nil,
			},
			{
				Val:      4,
				Children: nil,
			},
		},
	}

	res := maxDepth(inp)
	if res != 3 {
		t.Logf("Output: %d, Expected: 3", res)
		t.Fail()
	}

	inp = &Node{
		Val: 1,
		Children: []*Node{
			{
				Val:      2,
				Children: nil,
			},
			{
				Val: 3,
				Children: []*Node{
					{
						Val:      6,
						Children: nil,
					},
					{
						Val: 7,
						Children: []*Node{
							{
								Val: 11,
								Children: []*Node{
									{
										Val:      14,
										Children: nil,
									},
								},
							},
						},
					},
				},
			},
			{
				Val: 4,
				Children: []*Node{
					{
						Val: 8,
						Children: []*Node{
							{
								Val:      12,
								Children: nil,
							},
						},
					},
				},
			},
			{
				Val: 5,
				Children: []*Node{
					{
						Val: 9,
						Children: []*Node{
							{
								Val:      13,
								Children: nil,
							},
						},
					},
					{
						Val:      10,
						Children: nil,
					},
				},
			},
		},
	}

	res = maxDepth(inp)
	if res != 5 {
		t.Logf("Output: %d, Expected: 5", res)
		t.Fail()
	}
}
