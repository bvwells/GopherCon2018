package src_test

import (
	"fmt"
	"strings"
	"testing"
)

// Write a test for strings.HasPrefix
// https://golang.org/pkg/strings/#HasPrefix
// Given the value "main.go", test that it has the prefix "main"
// Remember that your test has to start with the name `Test` and be in an `_test.go` file
func Test_HasPrefix(t *testing.T) {
	value := "main.go"
	prefix := "main"
	if !strings.HasPrefix(value, prefix) {
		t.Fatalf("expected %s to have prefix %s", value, prefix)
	}
}

// Write a table drive test for strings.Index
// https://golang.org/pkg/strings/#Index
// Use the following test conditions
// |------------------------------------------------|
// | Value                     | Substring | Answer |
// |===========================|===========|========|
// | "Gophers are amazing!"    | "are"     | 8      |
// | "Testing in Go is fun."   | "fun"     | 17     |
// | "The answer is 42."       | "is"      | 11     |
// |------------------------------------------------|
func Test_Index(t *testing.T) {
	tt := []struct {
		Value     string
		Substring string
		Answer    int
	}{
		{
			"Gophers are amazing!",
			"are",
			8,
		},
		{
			"Testing in Go is fun.",
			"fun",
			17,
		},
		{
			"The answer is 42.",
			"is",
			11,
		},
	}

	for _, test := range tt {
		if actual := strings.Index(test.Value, test.Substring); actual != test.Answer {
			t.Fatalf("expected index of substring '%s' in string '%s' to be %v", test.Substring, test.Value, test.Answer)
		}
	}
}

// Rewrite the above test for strings.Index using subtests
func Test_Index_WithSubTest(t *testing.T) {
	tt := []struct {
		Value     string
		Substring string
		Answer    int
	}{
		{
			"Gophers are amazing!",
			"are",
			8,
		},
		{
			"Testing in Go is fun.",
			"fun",
			17,
		},
		{
			"The answer is 42.",
			"is",
			11,
		},
	}

	for i, test := range tt {

		t.Run(fmt.Sprintf("sub test (%d)", i), func(st *testing.T) {
			if actual := strings.Index(test.Value, test.Substring); actual != test.Answer {
				st.Fatalf("expected index of substring '%s' in string '%s' to be %v", test.Substring, test.Value, test.Answer)
			}
		})
	}
}

// Here is a simple test that tests `strings.HasSuffix`.i
// https://golang.org/pkg/strings/#HasSuffix
func Test_HasSuffix(t *testing.T) {
	value := "main.go"
	if !strings.HasSuffix(value, ".go") {
		t.Fatalf("expected %s to have suffix %s", value, ".go")
	}
}
