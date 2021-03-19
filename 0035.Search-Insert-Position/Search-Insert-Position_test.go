package searchinsertposition

import "testing"

var tests = []struct {
	arg1 []int
	arg2 int
	want int
}{
	{
		arg1: []int{1, 3, 5, 6},
		arg2: 5,
		want: 2,
	},
	{
		arg1: []int{1, 3, 5, 6},
		arg2: 2,
		want: 1,
	},
	{
		arg1: []int{1, 3, 5, 6},
		arg2: 7,
		want: 4,
	},
	{
		arg1: []int{1, 3, 5, 6},
		arg2: 0,
		want: 0,
	},
}

func TestSearchInsertBurst(t *testing.T) {
	for _, tt := range tests {
		if got := SearchInsertBurst(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestSearchInsertBisection(t *testing.T) {
	for _, tt := range tests {
		if got := SearchInsertBisection(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func TestSearchInsertBisection2(t *testing.T) {
	for _, tt := range tests {
		if got := SearchInsertBisection2(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}

func BenchmarkSearchInsertBurst(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SearchInsertBurst(tests[0].arg1, tests[0].arg2)
	}
}

func BenchmarkSearchInsertBisection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SearchInsertBisection(tests[0].arg1, tests[0].arg2)
	}
}

func BenchmarkSearchInsertBisection2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SearchInsertBisection2(tests[0].arg1, tests[0].arg2)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/0035.Search-Insert-Position -bench=.

*/
