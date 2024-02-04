package insertdeletegetrandom

import (
	"reflect"
	"testing"
)

var tests = []struct {
	arg1 []string
	arg2 []int
	want []interface{}
}{
	{
		[]string{
			"RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom",
		},
		[]int{
			0, 1, 1, 2, 0, 1, 2, 0,
		},
		[]interface{}{
			nil, true, true, true, 2, false, false, 2,
		},
	},
}

func TestInsertdeletegetrandom(t *testing.T) {
	for _, tt := range tests {
		result := []interface{}{}
		var rSet RandomizedSet
		for i := 0; i < len(tt.arg1); i++ {
			switch tt.arg1[i] {
			case "RandomizedSet":
				rSet = Constructor()
				result = append(result, nil)
			case "insert":
				result = append(result, rSet.Insert(tt.arg2[i]))
			case "remove":
				result = append(result, rSet.Remove(tt.arg2[i]))
			case "getRandom":
				result = append(result, rSet.GetRandom())
			}
		}
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("got = %v, want = %v", result, tt.want)
		}
	}
}
