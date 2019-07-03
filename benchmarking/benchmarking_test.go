package test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

const x = "hello"

var dumdum = ""

func withPlus() string {
	return x + x + x + x + x + x + x + x + x + x
}

func withSprintf() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", x, x, x, x, x, x, x, x, x, x)
}

func withBuffer() string {
	bb := &bytes.Buffer{}
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	return bb.String()
}

func withStringBuilder() string {
	bb := &strings.Builder{}
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	return bb.String()
}

func BenchmarkWithPlus(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dumdum = withPlus()
	}
}

func BenchmarkWithSprintf(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dumdum = withSprintf()
	}
}

func BenchmarkWithBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dumdum = withBuffer()
	}
}

func BenchmarkWithStringBuilder(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dumdum = withBuffer()
	}
}
