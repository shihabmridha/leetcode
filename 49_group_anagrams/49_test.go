package groupanagrams

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	res := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})

	if !reflect.DeepEqual(res, [][]string{
		{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"},
	}) {
		t.Logf(`output: %v, expected: [["bat"],["nat","tan"],["ate","eat","tea"]]`, res)
		t.Fail()
	}

	res = groupAnagrams([]string{""})
	if !reflect.DeepEqual(res, [][]string{{""}}) {
		t.Logf(`output: %v, expected: [[""]]`, res)
		t.Fail()
	}

	res = groupAnagrams([]string{"a"})
	if !reflect.DeepEqual(res, [][]string{{"a"}}) {
		t.Logf(`output: %v, expected: [["a"]]`, res)
		t.Fail()
	}

	res = groupAnagrams([]string{"", "b"})
	if !reflect.DeepEqual(res, [][]string{{""}, {"b"}}) {
		t.Logf(`output: %v, expected: [[""],["b"]]`, res)
		t.Fail()
	}
}
