package searchgraph

import "testing"

func TestSearchGraph(t *testing.T) {
	SearchGraph(1)
}

func TestSearchGraph2(t *testing.T) {
	SearchGraph2(1)
}
func BenchmarkSearchGraph(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SearchGraph(1)
	}
}

func BenchmarkSearchGraph2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SearchGraph2(1)
	}
}

// go test -benchmem -run=none LeetcodeGolang/Algorithms/SearchGraph -bench=.
/*

 */
