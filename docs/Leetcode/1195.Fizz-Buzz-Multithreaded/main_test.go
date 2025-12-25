package fizzbuzzmultithreaded

import "testing"

var tests = []struct {
	arg1 int
	want []interface{}
}{
	{
		15,
		[]interface{}{1, 2, "fizz", 4, "buzz", "fizz", 7, 8, "fizz", "buzz", 11, "fizz", 13, 14, "fizzbuzz"},
	},
}

func TestFizzBuzz(t *testing.T) {
	for _, tt := range tests {
		// if got := ReverseList(tt.arg1); !reflect.DeepEqual(got, tt.want) {
		// if got := FizzBuzz(tt.arg1); got != tt.want {
		// 	t.Errorf("got = %v, want = %v", got, tt.want)
		// }
		FizzBuzz(tt.arg1)
	}
}

func BenchmarkFizzBuzz(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FizzBuzz(tests[0].arg1)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0354.Russian-Doll-Envelopes -bench=.

*/
