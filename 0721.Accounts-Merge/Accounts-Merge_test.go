package accountsmerge

import (
	"reflect"
	"testing"
)

var tests = []struct {
	arg1 [][]string
	want [][]string
}{
	{
		[][]string{
			{"John", "johnsmith@mail.com", "john00@mail.com"},
			{"John", "johnnybravo@mail.com"},
			{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
			{"Mary", "mary@mail.com"},
		},
		[][]string{
			{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"},
			{"John", "johnnybravo@mail.com"},
			{"Mary", "mary@mail.com"},
		},
	},
	{
		[][]string{
			{"Alex", "Alex5@m.co", "Alex4@m.co", "Alex0@m.co"},
			{"Ethan", "Ethan3@m.co", "Ethan3@m.co", "Ethan0@m.co"},
			{"Kevin", "Kevin4@m.co", "Kevin2@m.co", "Kevin2@m.co"},
			{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe2@m.co"},
			{"Gabe", "Gabe3@m.co", "Gabe4@m.co", "Gabe2@m.co"},
		},
		[][]string{
			{"Alex", "Alex0@m.co", "Alex4@m.co", "Alex5@m.co"},
			{"Ethan", "Ethan0@m.co", "Ethan3@m.co"},
			{"Kevin", "Kevin2@m.co", "Kevin4@m.co"},
			{"Gabe", "Gabe0@m.co", "Gabe2@m.co", "Gabe3@m.co", "Gabe4@m.co"},
		},
	},
	{
		[][]string{
			{"David", "David0@m.co", "David4@m.co", "David3@m.co"},
			{"David", "David5@m.co", "David5@m.co", "David0@m.co"},
			{"David", "David1@m.co", "David4@m.co", "David0@m.co"},
			{"David", "David0@m.co", "David1@m.co", "David3@m.co"},
			{"David", "David4@m.co", "David1@m.co", "David3@m.co"},
		},
		[][]string{
			{"David", "David0@m.co", "David1@m.co", "David3@m.co", "David4@m.co", "David5@m.co"},
		},
	},
	{
		[][]string{},
		[][]string{},
	},
}

func TestAccountsMergeBurst(t *testing.T) {
	for _, tt := range tests {
		if got := AccountsMergeBurst(tt.arg1); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v \n want = %v \n", tt.arg1, tt.want)
		}
	}
}

func TestAccountsMerge(t *testing.T) {
	for _, tt := range tests {
		if got := accountsMerge(tt.arg1); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v \n want = %v \n", tt.arg1, tt.want)
		}
	}
}
