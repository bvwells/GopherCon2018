package test

import (
	"strings"
	"testing"
)

var dumdum = ""

func BenchmarkWithPlus(b *testing.B) {
	str := "Education is what remains after one has forgotten what one has learned in school."
	for n := 0; n < b.N; n++ {
		dumdum = strings.ToUpper(str)
	}
}
