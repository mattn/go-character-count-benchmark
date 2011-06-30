package count_test

import (
	. "count"
	. "utf8"
	"testing"
)

const benchMarkString = "Hello World, こんにちわ世界"

func BenchmarkLenViaIntArr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if LenViaIntArr(benchMarkString) != 20 {
			panic("unexpected result")
		}
	}
}

func BenchmarkRuneCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if RuneCountInString(benchMarkString) != 20 {
			panic("unexpected result")
		}
	}
}

func BenchmarkLenViaUTF8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if LenViaUTF8(benchMarkString) != 20 {
			panic("unexpected result")
		}
	}
}

