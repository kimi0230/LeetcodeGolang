package a1b2c3

import "testing"

func TestChannelWBuffer(t *testing.T) {
	ChannelWBuffer()
}

func TestChannelWOBuffer(t *testing.T) {
	ChannelWOBuffer()
}
func TestWGLock(t *testing.T) {
	WGLock()
}

func BenchmarkChannelWBuffer(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChannelWBuffer()
	}
}

func BenchmarkChannelWOBuffer(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChannelWOBuffer()
	}
}

func BenchmarkWGLock(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		WGLock()
	}
}

// go test -benchmem -run=none LeetcodeGolang/Algorithms/A1B2C3 -bench=.
/*
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Algorithms/A1B2C3
cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
BenchmarkChannelWBuffer-4          79159             14612 ns/op             344 B/op          5 allocs/op
BenchmarkChannelWOBuffer-4         83068             14451 ns/op             344 B/op          5 allocs/op
BenchmarkWGLock-4                  51303             23072 ns/op              96 B/op          5 allocs/op
PASS
ok      LeetcodeGolang/Algorithms/A1B2C3        4.092s
*/
