package validparentheses

import "testing"

var tests = []struct {
	arg1 string
	want bool
}{
	{
		"()",
		true,
	},
	{
		"()[]{}",
		true,
	},
	{
		"(]",
		false,
	},
	{
		"[",
		false,
	},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		// if got := ReverseList(tt.arg1); !reflect.DeepEqual(got, tt.want) {
		if got := IsValid(tt.arg1); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkIsValid(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsValid(tests[0].arg1)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0354.Russian-Doll-Envelopes -bench=.

*/
