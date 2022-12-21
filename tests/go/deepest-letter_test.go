package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetDeepestLetter(t *testing.T) {
	testCases := []letterTestCase{
		{input: "a(b)c", expected: 'b'},
		{input: "((a))(((M)))(c)(D)(e)(((f))(((G))))h(i)", expected: 'g'},
		{input: "((A)(b)c", expected: '?'},
		{input: "(a)((G)c)", expected: 'g'},
		{input: "(8)", expected: '?'},
		{input: "(!)", expected: '?'},
	}

	for _, test := range testCases {
		t.Run(test.input, func(t *testing.T) {
			actual := getDeepestLetter(test.input)
			if !cmp.Equal(test.expected, actual) {
				t.Log(cmp.Diff(test.expected, actual))
				t.Fail()
			}
		})
	}
}
