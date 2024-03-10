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
		[]int{-1},
	},
	{
		[]string{
			"withdraw 2 10",
			"transfer 5 1 20",
			"deposit 5 20",
			"transfer 3 4 15",
		},
		[]int{10, 100, 20, 50, 30},
		[]int{30, 90, 5, 65, 30},
	},
	{
		[]string{
			"deposit 3 400",
			"transfer 1 2 30",
			"withdraw 4 50",
		},
		[]int{20, 1000, 500, 40, 90},
		[]int{-2},
	},
}

func TestBankRequests(t *testing.T) {
	for _, tt := range tests {
		if got := bankRequests(tt.arg1, tt.arg2); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}
