package insertdeletegetrandomo1duplicatesallowed

import (
	"math/rand"
)

// 時間複雜 O(1), 空間複雜 O(n)
type RandomizedCollection struct {
	arr  []int
	set  map[int]map[int]struct{}
	size int
}

func Constructor() RandomizedCollection {
	arr := make([]int, 0)
	set := make(map[int]map[int]struct{})
	size := 0
	return RandomizedCollection{arr, set, size}
}

func (this *RandomizedCollection) Insert(val int) bool {
	ids, ok := this.set[val]
	if !ok {
		// 沒有此數字了, 建立一個index的map
		ids = map[int]struct{}{}
		this.set[val] = ids
	}
	// index的map, key為當前最後的index, value為空的struct{}{}
	ids[this.size] = struct{}{}
	this.arr = append(this.arr, val)
	this.size++
	return !ok
}

func (this *RandomizedCollection) Remove(val int) bool {
	ids, ok := this.set[val]
	if !ok {
		return false
	}

	// 找出此val的的index, 並且把最後一個index的map刪除, 並且把此index的value設為空的struct{}{}
	var i int
	for id := range ids {
		i = id
		break
	}

	// 將最後一筆移到要替換的id
	this.arr[i] = this.arr[this.size-1]
	delete(ids, i)
	// 因為把最後一個元素移到前面了
	delete(this.set[this.arr[i]], this.size-1)

	if i < this.size-1 {
		// fmt.Printf("i =%d, this.size =%d\t", i, this.size)
		// 將最後一個元素的index的map中, 最後一個元素的index改為i
		this.set[this.arr[i]][i] = struct{}{}
	}
	if len(ids) == 0 {
		// 沒有此數字了
		delete(this.set, val)
	}

	this.arr = this.arr[:this.size-1]
	this.size--
	return true
}

func (this *RandomizedCollection) GetRandom() int {
	index := rand.Intn(this.size)
	return this.arr[index]
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
