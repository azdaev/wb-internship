package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type sortStringTest struct {
	str      string
	expected string
}

type groupAnagramsTest struct {
	words    []string
	expected *map[string][]string
}

var sortStringTests = []sortStringTest{
	{"dabc", "abcd"},
	{"caaa", "aaac"},
	{"bike", "beik"},
}

var groupAnagramsTests = []groupAnagramsTest{
	{words: []string{"eat", "ate", "teA", "bike", "kibe", "cab", "bca", "Abc", "aBc", "a", "b"}, expected: &map[string][]string{"bike": []string{"bike", "kibe"}, "cab": []string{"abc", "bca", "cab"}, "eat": []string{"ate", "eat", "tea"}}},
	{words: []string{"eat", "eat", "eat", "ate", "teA", "bike", "kibe", "cab", "bca", "Abc", "aBc"}, expected: &map[string][]string{"bike": []string{"bike", "kibe"}, "cab": []string{"abc", "bca", "cab"}, "eat": []string{"ate", "eat", "tea"}}},
}

func TestSortString(t *testing.T) {
	for _, test := range sortStringTests {
		assert.Equal(t, test.expected, SortString(test.str))
	}
}

func TestGroupAnagrams(t *testing.T) {
	for _, test := range groupAnagramsTests {
		assert.Equal(t, test.expected, GroupAnagrams(test.words))
	}
}
