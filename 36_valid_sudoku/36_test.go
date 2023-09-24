package validsudoku

import "testing"

func TestIsValidSudoku(t *testing.T) {

	input := [][]byte{
		[]byte(string([]rune{'5', '3', '.', '.', '7', '.', '.', '.', '.'})),
		[]byte(string([]rune{'6', '.', '.', '1', '9', '5', '.', '.', '.'})),
		[]byte(string([]rune{'.', '9', '8', '.', '.', '.', '.', '6', '.'})),
		[]byte(string([]rune{'8', '.', '.', '.', '6', '.', '.', '.', '3'})),
		[]byte(string([]rune{'4', '.', '.', '8', '.', '3', '.', '.', '1'})),
		[]byte(string([]rune{'7', '.', '.', '.', '2', '.', '.', '.', '6'})),
		[]byte(string([]rune{'.', '6', '.', '.', '.', '.', '2', '8', '.'})),
		[]byte(string([]rune{'.', '.', '.', '4', '1', '9', '.', '.', '5'})),
		[]byte(string([]rune{'.', '.', '.', '.', '8', '.', '.', '7', '9'})),
	}

	res := isValidSudoku(input)
	if !res {
		t.Logf("Output: %t, Expected: 1", res)
		t.Fail()
	}

	input = [][]byte{
		[]byte(string([]rune{'8', '3', '.', '.', '7', '.', '.', '.', '.'})),
		[]byte(string([]rune{'6', '.', '.', '1', '9', '5', '.', '.', '.'})),
		[]byte(string([]rune{'.', '9', '8', '.', '.', '.', '.', '6', '.'})),
		[]byte(string([]rune{'8', '.', '.', '.', '6', '.', '.', '.', '3'})),
		[]byte(string([]rune{'4', '.', '.', '8', '.', '3', '.', '.', '1'})),
		[]byte(string([]rune{'7', '.', '.', '.', '2', '.', '.', '.', '6'})),
		[]byte(string([]rune{'.', '6', '.', '.', '.', '.', '2', '8', '.'})),
		[]byte(string([]rune{'.', '.', '.', '4', '1', '9', '.', '.', '5'})),
		[]byte(string([]rune{'.', '.', '.', '.', '8', '.', '.', '7', '9'})),
	}

	res = isValidSudoku(input)
	if res {
		t.Logf("Output: %t, Expected: 0", res)
		t.Fail()
	}
}
