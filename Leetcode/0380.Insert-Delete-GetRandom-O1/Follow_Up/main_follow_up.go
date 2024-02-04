package insertdeletegetrandomfollowup

import "math/rand"

// 時間複雜 O(1), 空間複雜 O(n)
type RandomizedSet struct {
	arr  []int
	set  map[int][]int
	size int
}

func Constructor() RandomizedSet {
	arr := make([]int, 0)
	set := make(map[int][]int)
	size := 0
	return RandomizedSet{arr, set, size}
}

func (this *RandomizedSet) Insert(val int) bool {
	this.arr = append(this.arr, val)
	this.set[val] = append(this.set[val], this.size)
	this.size++
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	indexArry, ok := this.set[val]
	if !ok {
		return false
	}
	index := indexArry[0]

	// swapping
	lastValue := this.arr[this.size-1]
	this.arr[index] = lastValue
	this.set[lastValue] = append(this.set[lastValue], index)

	// Remove last value
	this.arr = this.arr[:this.size-1]
	this.set[val] = this.set[val][1:]
	this.size--
	return true
}

func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(this.size)
	return this.arr[index]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
