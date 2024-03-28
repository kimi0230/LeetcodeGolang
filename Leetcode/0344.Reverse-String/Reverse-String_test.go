package reversestring

import (
	"reflect"
	"testing"
)

var tests = []struct {
	arg1 []byte
	want []byte
}{
	{
		[]byte{'h', 'e', 'l', 'l', 'o'},
		[]byte{'o', 'l', 'l', 'e', 'h'},
	},
	{
		[]byte{'H', 'a', 'n', 'n', 'a', 'h'},
		[]byte{'h', 'a', 'n', 'n', 'a', 'H'},
	},
}

func TestReverseString(t *testing.T) {
	for _, tt := range tests {
		ReverseString(tt.arg1)
		if !reflect.DeepEqual(tt.arg1, tt.want) {
			t.Errorf("got = %v \n want = %v \n", tt.arg1, tt.want)
		}
	}
}
