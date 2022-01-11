package openthelock

import (
	"testing"
)

var tests = []struct {
	arg1 []string
	arg2 string
	want int
}{
	{
		[]string{"0201", "0101", "0102", "1212", "2002"},
		"0202",
		6,
	},
	{
		[]string{"8888"},
		"0009",
		1,
	},
	{
		[]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"},
		"8888",
		-1,
	},
}

func TestOpenLock(t *testing.T) {
	for _, tt := range tests {
		if got := OpenLock(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("got = %v \n want = %v \n", got, tt.want)
		}
	}
}
