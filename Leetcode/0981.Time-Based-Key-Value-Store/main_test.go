package timebasedkeyvaluestore

import (
	"testing"
)

func TestTimeMap(t *testing.T) {
	timeMap := Constructor()
	timeMap.Set("foo", "bar", 1)
	timeMap.Set("foo", "bar2", 4)
	timeMap.Set("foo", "bar3", 6)

	tests := []struct {
		key       string
		timestamp int
		want      string
	}{
		{"foo", 1, "bar"},
		{"foo", 3, "bar"},
		{"foo", 4, "bar2"},
		{"foo", 5, "bar2"},
		{"foo", 6, "bar3"},
		{"foo", 7, "bar3"},
		{"foo2", 1, ""},
	}

	for _, tt := range tests {
		got := timeMap.Get(tt.key, tt.timestamp)
		if got != tt.want {
			t.Errorf("Get(%v, %v) = %v, want %v", tt.key, tt.timestamp, got, tt.want)
		}
	}
}

func BenchmarkTimeMap_Get(b *testing.B) {
	timeMap := Constructor()
	timeMap.Set("foo", "bar", 1)
	timeMap.Set("foo", "bar2", 4)
	timeMap.Set("foo", "bar3", 6)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		timeMap.Get("foo", 5)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0354.Russian-Doll-Envelopes -bench=.

*/
