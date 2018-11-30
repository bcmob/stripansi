package stripansi

import (
	"strings"
	"testing"
)

const testStr = "\u001B[4m<cake>\u001B[0m\u001B[4m<cake>\u001B[0m\u001B[4m<cake>\u001B[0m\u001B[4m<cake>\u001B[0m\u001B[4m<cake>\u001B[0m\u001B[4m<cake>\u001B[0m\u001B[4m<cake>\u001B[0m\u001B[4m<cake>\u001B[0m\u001B[4m<cake>\u001B[0m\u001B[4m<cake>\u001B[0m\u001B[4m<cake>\u001B[0m\u001B[4m<cake>\u001B[0m\u001B[4m<cake>\u001B[0m\u001B[4m<cake>\u001B[0m\u001B[4m<cake>\u001B[0m\u001B[4m<cake>\u001B[0m"

func BenchmarkStrip(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := Strip(testStr)
		if strings.Contains(s, "<") {
			b.Fatalf("Unexpected string: %s", s)
		}
	}
}
func BenchmarkFullR(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := StripFullR(testStr)
		if strings.Contains(s, "<") {
			b.Fatalf("Unexpected string: %s", s)
		}
	}
}

func BenchmarkFull(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := StripFull(testStr)
		if strings.Contains(s, "<") {
			b.Fatalf("Unexpected string: %s", s)
		}
	}
}
