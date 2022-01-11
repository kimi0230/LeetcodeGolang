package openthelock

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
