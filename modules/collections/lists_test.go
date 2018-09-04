package collections

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListContains(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		description string
		list        []string
		element     string
		expected    bool
	}{
		{"empty list, empty element", []string{}, "", false},
		{"empty list, non-empty element", []string{}, "foo", false},
		{"list with 1 item, element matches", []string{"foo"}, "foo", true},
		{"list with 1 item, element doesn't match", []string{"bar"}, "foo", false},
		{"list with 3 items, element matches", []string{"bar", "foo", "baz"}, "foo", true},
		{"list with 3 items, element doesn't match", []string{"bar", "foo", "baz"}, "nope", false},
		{"list with 3 items, empty element", []string{"bar", "foo", "baz"}, "", false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actual := ListContains(testCase.list, testCase.element)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		description string
		list1       []string
		list2       []string
		expected    []string
	}{
		{"empty list, empty list", []string{}, []string{}, []string{}},
		{"empty list, non-empty list", []string{}, []string{"foo"}, []string{}},
		{"non-empty list, empty list", []string{"foo"}, []string{}, []string{"foo"}},
		{"list with 1 item, list with no matches", []string{"foo"}, []string{"bar"}, []string{"foo"}},
		{"list with 1 item, list with 1 match", []string{"foo"}, []string{"foo"}, []string{}},
		{"list with 1 item, list with multiple matches and non-matches", []string{"foo"}, []string{"foo", "bar", "foo"}, []string{}},
		{"list with multiple items, list with no matches", []string{"foo", "bar", "baz"}, []string{"abc", "def"}, []string{"foo", "bar", "baz"}},
		{"list with multiple items, list with 1 match", []string{"foo", "bar", "baz"}, []string{"abc", "foo", "def"}, []string{"bar", "baz"}},
		{"list with multiple items, list with multiple matches", []string{"foo", "bar", "baz", "foo", "bar", "baz"}, []string{"abc", "foo", "baz"}, []string{"bar", "bar"}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actual := ListSubtract(testCase.list1, testCase.list2)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}

func TestGetElement(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		description string
		list     []string
		index    int
		expected string
	}{
		{"full list, first element", []string{"A", "B", "C"}, 0, "A"},
		{"full list, last element", []string{"A", "B", "C"}, 2, "C"},
		{"single-element list, first element", []string{"A"}, 0, "A"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actual := GetElement(t, testCase.list, testCase.index)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}

func TestGetElementWithFailure(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		description string
		list     []string
		index    int
		expected string
	}{
		{"full list, element too large", []string{"A", "B", "C"}, 4, "A"},
		//{"full list, element too small", []string{"A", "B", "C"}, -1, "C"},
	}

	for _, testCase := range testCases {
		GetElementThatPanics(t, testCase.description, testCase.list, testCase.index)
	}
}

func GetElementThatPanics(t *testing.T, description string, list []string, index int) {
	assert.Panics(
		t,
		func() { GetElement(t, list, index) },
		fmt.Sprintf("Test Case: %s\nExpected code to fail with invalid index.", description),
	)
}
