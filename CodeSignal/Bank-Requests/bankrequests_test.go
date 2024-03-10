package bankrequests

import (
	"reflect"
	"testing"
)

var tests = []struct {
	arg1 []string
	arg2 []int
	want []int
}{
	{
		[]string{
			"transfer 1 4 10",
			"deposit 3 10",
			"withdraw 5 15",
		},
		[]int{20, 30, 10, 90, 60},
		[]int{10, 30, 20, 100, 45},
	},
	{
		[]string{
			"transfer 1 4 40",
			"deposit 3 10",
			"withdraw 5 65",
		},
		[]int{20, 30, 10, 90, 60},
		[]int{0},
	},
}

func TestBankRequests(t *testing.T) {
	for _, tt := range tests {
		if got := bankRequests(tt.arg1, tt.arg2); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}
