package count_test

import (
	. "count"
	"testing"
)

const benchMarkString = "Hello World, こんにちわ世界"
const benchMarkCount = 1000000

func BenchmarkCharacterCount1(b *testing.B) {
	for i := 0; i < benchMarkCount; i++ {
		if Count1(benchMarkString) != 20 {
			panic("unexpected result")
		}
	}
}

func BenchmarkCharacterCount2(b *testing.B) {
	for i := 0; i < benchMarkCount; i++ {
		if Count2(benchMarkString) != 20 {
			panic("unexpected result")
		}
	}
}

