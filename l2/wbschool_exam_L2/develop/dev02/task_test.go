package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type repeatTest struct {
	char     rune
	n        int
	expected []rune
}

type parseTest struct {
	s              string
	expectedString string
	expectedError  bool
}

var repeatTests = []repeatTest{
	{'a', 4, []rune{'a', 'a', 'a', 'a'}},
	{'a', 1, []rune{'a'}},
	{'\\', 5, []rune{'\\', '\\', '\\', '\\', '\\'}},
}

var parseTests = []parseTest{
	{"a4bc2d5e", "aaaabccddddde", false},
	{"abcd", "abcd", false},
	{"45", "", true},
	{"", "", false},
	{"qwe\\4\\5", "qwe45", false},
	{"qwe\\45", "qwe44444", false},
	{"qwe\\\\5", "qwe\\\\\\\\\\", false},
}

func TestRepeat(t *testing.T) {
	for _, test := range repeatTests {
		output := Repeat(test.char, test.n)
		assert.Equal(t, test.expected, output, "they should be equal")
	}
}

func TestParse(t *testing.T) {
	for _, test := range parseTests {
		output, err := Parse(test.s)
		if test.expectedError {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}

		assert.Equal(t, test.expectedString, output)
	}
}
