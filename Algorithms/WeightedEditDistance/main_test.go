package weightededitdistance

import (
	"testing"
)

func TestWeightedEditDistance(t *testing.T) {
	tests := []struct {
		word1    string
		word2    string
		weights  map[rune]int
		expected int
	}{
		{"abc", "", map[rune]int{'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5, 'f': 6}, 6},
		{"abc", "def", map[rune]int{'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5, 'f': 6}, 21},
		{"kitten", "sitting", map[rune]int{'k': 1, 'i': 2, 't': 3, 'e': 4, 'n': 5, 's': 6, 'g': 7}, 20},
	}

	for _, test := range tests {
		actual := WeightedEditDistance(test.word1, test.word2, test.weights)
		if actual != test.expected {
			t.Errorf("WeightedEditDistance(%s, %s, %v) = %d, expected %d", test.word1, test.word2, test.weights, actual, test.expected)
		}
	}
}

/*
{"abc", "def", map[rune]int{'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5, 'f': 6}, 21},
a 與 d 的替換代價：1 + 4 = 5
b 與 e 的替換代價：2 + 5 = 7
c 與 f 的替換代價：3 + 6 = 9
因此，最小操作代價是 5 + 7 + 9 = 21

{"kitten", "sitting", map[rune]int{'k': 1, 'i': 2, 't': 3, 'e': 4, 'n': 5, 's': 6, 'g': 7}, 14},
k 转换为 s 的替换代价：1 + 6 = 7
i 转换为 i 的替换代价：0
t 转换为 t 的替换代价：0
t 转换为 t 的替换代价：0
e 转换为 i 的替换代价：4 + 2 = 6
n 转换为 n 的替换代价：5 + 5 = 10
因此，最小操作代价是 7 + 4 + 6 + 6 + 6 + 10 = 39

*/
