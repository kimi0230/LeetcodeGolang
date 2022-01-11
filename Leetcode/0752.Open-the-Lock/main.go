package openthelock

// 方法ㄧ: 單向BFS
func OpenLock(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}
	targetNum := strToInt(target)

	// 紀錄已窮舉過的密碼, 防止走回頭路
	visited := make([]bool, 10000)
	visited[0] = true
	for _, deadend := range deadends {
		num := strToInt(deadend)
		if num == 0 {
			return -1
		}
		visited[num] = true
	}

	depth, curDepth, nextDepth := 0, []int16{0}, make([]int16, 0)
	var nextNum int16

	for len(curDepth) > 0 {
		nextDepth = nextDepth[0:0]
		// 當前Queue中所有的節點向外擴散
		for _, curNum := range curDepth {
			// 遍歷八種組合
			for incrementer := int16(1000); incrementer > 0; incrementer /= 10 {
				nextNum = PlusOne(curNum, incrementer)
				if nextNum == targetNum {
					return depth + 1
				}
				if !visited[nextNum] {
					visited[nextNum] = true
					nextDepth = append(nextDepth, nextNum)
				}

				nextNum = MinusOne(curNum, incrementer)
				if nextNum == targetNum {
					return depth + 1
				}
				if !visited[nextNum] {
					visited[nextNum] = true
					nextDepth = append(nextDepth, nextNum)
				}
			}
		}
		curDepth, nextDepth = nextDepth, curDepth
		// 增加步數
		depth++
	}

	return -1
}

/* Note: Golang set
type void struct{}
var member void

set := make(map[string]void) // New empty set
set["Foo"] = member          // Add
for k := range set {         // Loop
    fmt.Println(k)
}
delete(set, "Foo")      // Delete
size := len(set)        // Size
_, exists := set["Foo"] // Membership
*/

// 方法二: 雙向BFS. 不在使用 Queue而是用 hashset, 快速判斷兩者是否交集
// 從起點跟終點開始擴散, 當兩邊有交集時停止
func OpenLockBiDirection(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}
	targetNum := strToInt(target)

	// 紀錄已窮舉過的密碼, 防止走回頭路
	visited := make([]bool, 10000)
	for _, deadend := range deadends {
		num := strToInt(deadend)
		if num == 0 {
			return -1
		}
		visited[num] = true
	}

	depth := 0

	// 起點跟終點初始化
	curDepth := make(map[int16]struct{})
	nextDepth := make(map[int16]struct{})
	curDepth[0] = struct{}{}
	nextDepth[targetNum] = struct{}{}

	var nextNum int16

	for len(curDepth) != 0 && len(nextDepth) != 0 {
		// 儲存 curDepth 的擴散結果
		tmp := make(map[int16]struct{})

		// curDepth的節點向外擴散
		for curNum := range curDepth {
			// 判斷是否達到終點
			if visited[curNum] {
				continue
			}
			_, exists := nextDepth[curNum]
			if exists {
				return depth
			}
			visited[curNum] = true

			// 遍歷八種組合
			for incrementer := int16(1000); incrementer > 0; incrementer /= 10 {
				nextNum = PlusOne(curNum, incrementer)
				if !visited[nextNum] {
					tmp[nextNum] = struct{}{}
				}

				nextNum = MinusOne(curNum, incrementer)
				if !visited[nextNum] {
					tmp[nextNum] = struct{}{}
				}
			}
		}

		// 小技巧, 這裏交換 curDepth, nextDepth .
		// 下一輪 whihe會擴散 nextDepth.
		// 所以只要默認擴散curDepth, 就相當於輪流擴散curDepth, nextDepth
		curDepth, nextDepth = nextDepth, tmp
		// 增加步數
		depth++
	}

	return -1
}

func PlusOne(curNum int16, incrementer int16) (nextNum int16) {
	digit := (curNum / incrementer) % 10
	if digit == 9 {
		nextNum = curNum - 9*incrementer
	} else {
		nextNum = curNum + incrementer
	}
	return nextNum
}

func MinusOne(curNum int16, incrementer int16) (nextNum int16) {
	digit := (curNum / incrementer) % 10
	if digit == 0 {
		nextNum = curNum + 9*incrementer
	} else {
		nextNum = curNum - incrementer
	}
	return nextNum
}

func strToInt(str string) int16 {
	return int16(str[0]-'0')*1000 + int16(str[1]-'0')*100 + int16(str[2]-'0')*10 + int16(str[3]-'0')
}
