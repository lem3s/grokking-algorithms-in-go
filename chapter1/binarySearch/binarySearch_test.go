package binarysearch

import (
	"testing"
)

func Test_BinarySearch(t *testing.T) {
	list := []int{3, 4, 9, 23, 44, 98, 243, 2000}
	expected := 1

	result := BinarySearch(list, 4)

	if result != expected {
		t.Errorf(
			"\"BinarySearch(%v, %v)\" FAILED, expected -> %v, got -> %v",
			list,
			expected,
			expected,
			result,
		)
	} else {

		t.Logf("\"BinarySearch(%v, %v)\" SUCCEDED, expected -> %v, got -> %v", list, expected, expected, result)
	}
}
