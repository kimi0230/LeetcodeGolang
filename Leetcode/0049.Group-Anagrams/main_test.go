package groupanagrams

import (
	"reflect"
	"sort"
	"testing"
)

var tests = []struct {
	arg1 []string
	want [][]string
}{
	{
		[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
		[][]string{
			{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
	},
	{
		[]string{},
		[][]string{},
	}, {
		[]string{"a"},
		[][]string{
			{"a"}},
	},
}

// 創建一個輔助函數，將字符串數組排序，以便比較
func sortSubStrings(groups [][]string) {
	for _, group := range groups {
		sort.Strings(group)
	}
	// 排序整個切片
	sort.Slice(groups, func(i, j int) bool {
		return groups[i][0] < groups[j][0]
	})
}

func TestGroupAnagrams(t *testing.T) {
	for _, tt := range tests {
		result := GroupAnagrams(tt.arg1)
		// 對結果進行排序，以便比較
		sortSubStrings(result)

		// 對期望結果進行排序，以便比較
		sortSubStrings(tt.want)

		// 使用反射檢查結果和期望是否相同
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("got = %v, want = %v", result, tt.want)
		}
	}
}

func TestGroupAnagramsBySort(t *testing.T) {
	for _, tt := range tests {
		result := GroupAnagramsBySort(tt.arg1)
		// 對結果進行排序，以便比較
		sortSubStrings(result)

		// 對期望結果進行排序，以便比較
		sortSubStrings(tt.want)

		// 使用反射檢查結果和期望是否相同
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("got = %v, want = %v", result, tt.want)
		}
	}
}

func BenchmarkGroupAnagrams(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GroupAnagrams(tests[0].arg1)
	}
}

func BenchmarkGroupAnagramsBySort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GroupAnagramsBySort(tests[0].arg1)
	}
}

/*
go test -benchmem -run=none LeetcodeGolang/Leetcode/0049.Group-Anagrams -bench=.
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0049.Group-Anagrams
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkGroupAnagrams-8                  899431              1615 ns/op             968 B/op         12 allocs/op
BenchmarkGroupAnagramsBySort-8            592148              3566 ns/op             776 B/op         33 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0049.Group-Anagrams     3.611s
*/
