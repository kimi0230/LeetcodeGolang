package insertdeletegetrandomo1duplicatesallowed

import (
	"fmt"
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
			"RandomizedCollection", "insert", "insert", "insert", "remove", "getRandom",
		},
		[]int{
			0, 1, 1, 2, 2, 1,
		},
		[]interface{}{
			nil, true, false, true, true, 1,
		},
	},
	// {
	// 	[]string{
	// 		"RandomizedCollection", "insert", "insert", "insert", "getRandom", "remove", "getRandom",
	// 	},
	// 	[]int{
	// 		0, 1, 1, 2, 0, 1, 0,
	// 	},
	// 	[]interface{}{
	// 		nil, true, false, true, 2, true, false, 2,
	// 	},
	// },
}

func TestInsertdeletegetrandom(t *testing.T) {
	for _, tt := range tests {
		result := []interface{}{}
		var rSet RandomizedCollection
		for i := 0; i < len(tt.arg1); i++ {
			switch tt.arg1[i] {
			case "RandomizedCollection":
				rSet = Constructor()
				result = append(result, nil)
			case "insert":
				result = append(result, rSet.Insert(tt.arg2[i]))
			case "remove":
				result = append(result, rSet.Remove(tt.arg2[i]))
			case "getRandom":
				result = append(result, rSet.GetRandom())
			}
			fmt.Println(rSet.set)
		}
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("got = %v, want = %v", result, tt.want)
		}
	}
}
