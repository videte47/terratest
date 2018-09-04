package collections

import (
	"testing"
)

// ListSubtract removes all the items in list2 from list1.
func ListSubtract(list1 []string, list2 []string) []string {
	out := []string{}

	for _, item := range list1 {
		if !ListContains(list2, item) {
			out = append(out, item)
		}
	}

	return out
}

// ListContains returns true if the given list of strings (haystack) contains the given string (needle).
func ListContains(haystack []string, needle string) bool {
	for _, str := range haystack {
		if needle == str {
			return true
		}
	}

	return false
}

func GetElement(t *testing.T, list[] string, index int) string {
	if index >= len(list) {
		t.Fatalf("Element %d does not exist in the given list of length %d", index, len(list))
	}

	return list[index]
}