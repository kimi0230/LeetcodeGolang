package editdistance

import "testing"

var tests = []struct {
	arg1 string
	arg2 string
	want int
}{
	{
		"horse",
		"ros",
		3,
	},
	{
		"intention",
		"execution",
		5,
	},
}

func TestMinDistance(t *testing.T) {
	for _, tt := range tests {
		if got := MinDistance(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkMinDistance(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinDistance(tests[i%2].arg1, tests[i%2].arg2)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0072.Edit-Distance -bench=.
*/
