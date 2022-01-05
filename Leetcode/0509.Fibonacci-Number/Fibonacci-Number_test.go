package fibonaccinumber

import (
	"testing"
)

func TestFib(t *testing.T) {
	tests := []struct {
		name string
		in   int
		want int
	}{
		{
			name: "FibIterative",
			in:   39,
			want: 63245986,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FibIterative(tt.in); got != tt.want {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestFibrecur(t *testing.T) {
	tests := []struct {
		name string
		in   int
		want int
	}{
		{
			name: "Fibrecur",
			in:   39,
			want: 63245986,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fibrecur(tt.in); got != tt.want {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestFibDP(t *testing.T) {
	tests := []struct {
		name string
		in   int
		want int
	}{
		{
			name: "FibDP",
			in:   39,
			want: 63245986,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FibDP(tt.in); got != tt.want {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestFibDPStateCompression(t *testing.T) {
	tests := []struct {
		name string
		in   int
		want int
	}{
		{
			name: "FibDPStateCompression",
			in:   39,
			want: 63245986,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FibDPStateCompression(tt.in); got != tt.want {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestFibmem(t *testing.T) {
	tests := []struct {
		name string
		in   int
		want int
	}{
		{
			name: "Fibmem",
			in:   39,
			want: 63245986,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fibmem(tt.in); got != tt.want {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestFib5(t *testing.T) {
	tests := []struct {
		name string
		in   int
		want int
	}{
		{
			name: "Fib5",
			in:   39,
			want: 63245986,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fib5(tt.in); got != tt.want {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestFib6(t *testing.T) {
	tests := []struct {
		name string
		in   int
		want int
	}{
		{
			name: "公式法",
			in:   39,
			want: 63245986,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fib6(tt.in); got != tt.want {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestFib7(t *testing.T) {
	tests := []struct {
		name string
		in   int
		want int
	}{
		{
			name: "協程版",
			in:   39,
			want: 63245986,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fib7(tt.in); got != tt.want {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFib(b *testing.B) {
	input := 39
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FibIterative(input)
	}
}

func BenchmarkFibrecur(b *testing.B) {
	input := 39
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Fibrecur(input)
	}
}

func BenchmarkFibmem(b *testing.B) {
	input := 39
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Fibmem(input)
	}
}

func BenchmarkFibDP(b *testing.B) {
	input := 39
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FibDP(input)
	}
}

func BenchmarkFibDPStateCompression(b *testing.B) {
	input := 39
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FibDPStateCompression(input)
	}
}

// go test -benchmem -run=none MyGoNote/algorithms/Fibonacci -bench=.
/*
goos: darwin
goarch: amd64
pkg: MyGoNote/algorithms/Fibonacci
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkFib-8                           3409812               435.8 ns/op           992 B/op          5 allocs/op
BenchmarkFibrecur-8                            3         391624567 ns/op               0 B/op          0 allocs/op
BenchmarkFibmem-8                       106473912               11.56 ns/op            0 B/op          0 allocs/op
BenchmarkFibmem2-8                      28739854                49.79 ns/op           16 B/op          1 allocs/op
BenchmarkFibmemStateCompression-8       123085564               10.72 ns/op            0 B/op          0 allocs/op
PASS
ok      MyGoNote/algorithms/Fibonacci   12.090s
*/
