package timebasedkeyvaluestore

import "sort"

// 時間複雜 O(), 空間複雜 O()
type element struct {
	key       string
	value     string
	timestamp int
}

type TimeMap struct {
	elements map[string][]element
}

func Constructor() TimeMap {
	return TimeMap{elements: make(map[string][]element)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.elements[key] = append(this.elements[key], element{
		key:       key,
		value:     value,
		timestamp: timestamp,
	})
}

// O(log n)
func (this *TimeMap) Get(key string, timestamp int) string {
	if _, ok := this.elements[key]; !ok {
		return ""
	}
	// 二分法找出符合條件的最左邊
	index := sort.Search(len(this.elements[key]), func(i int) bool {
		return this.elements[key][i].timestamp > timestamp
	})
	if index == 0 {
		return ""
	}

	index--
	return this.elements[key][index].value
}
