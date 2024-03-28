package fizzbuzz

import (
	"reflect"
	"testing"
)

var tests = []struct {
	arg1 int
	want []string
}{
	{
		3,
		[]string{"1", "2", "Fizz"},
	},
}

func TestFizzBuzz(t *testing.T) {
	for _, tt := range tests {
		if got := fizzBuzz(tt.arg1); !reflect.DeepEqual(got, tt.want) {
			// if got := fizzBuzz(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkFizzBuzz(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fizzBuzz(tests[0].arg1)
	}
}

func TestFizzBuzz2(t *testing.T) {
	for _, tt := range tests {
		if got := fizzBuzz2(tt.arg1); !reflect.DeepEqual(got, tt.want) {
			// if got := fizzBuzz2(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkFizzBuzz2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fizzBuzz2(tests[0].arg1)
	}
}

func TestFizzBuzz3(t *testing.T) {
	for _, tt := range tests {
		if got := fizzBuzz3(tt.arg1); !reflect.DeepEqual(got, tt.want) {
			// if got := fizzBuzz3(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkFizzBuzz3(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fizzBuzz3(tests[0].arg1)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0412.Fizz-Buzz -bench=.

*/
