# [752. Open the Lock](https://leetcode.com/problems/open-the-lock/) 
###### tags: `Medium` `Leetcode`

You have a lock in front of you with 4 circular wheels. Each wheel has 10 slots: '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'. The wheels can rotate freely and wrap around: for example we can turn '9' to be '0', or '0' to be '9'. Each move consists of turning one wheel one slot.

The lock initially starts at '0000', a string representing the state of the 4 wheels.

You are given a list of deadends dead ends, meaning if the lock displays any of these codes, the wheels of the lock will stop turning and you will be unable to open it.

Given a target representing the value of the wheels that will unlock the lock, return the minimum total number of turns required to open the lock, or -1 if it is impossible.


Example 1:
```
Input: deadends = ["0201","0101","0102","1212","2002"], target = "0202"
Output: 6
Explanation: 
A sequence of valid moves would be "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202".
Note that a sequence like "0000" -> "0001" -> "0002" -> "0102" -> "0202" would be invalid,
because the wheels of the lock become stuck after the display becomes the dead end "0102".
```

Example 2:
```
Input: deadends = ["8888"], target = "0009"
Output: 1
Explanation: We can turn the last wheel in reverse to move from "0000" -> "0009".
```

Example 3:
```
Input: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"], target = "8888"
Output: -1
Explanation: We cannot reach the target without getting stuck.
```

Constraints:
* 1 <= deadends.length <= 500
* deadends[i].length == 4
* target.length == 4
* target will not be in the list deadends.
* target and deadends[i] consist of digits only.

## 題目大意
有一個四個圓形播盤的轉盤鎖, 每個播盤有0~9共10個數字, 每個播盤上下旋轉可以把 0變成9 或 9變成0. **每次只能旋轉一個播盤**
初始直接為0.且有一組 deadends 數組. 不能接數字轉到其中任一組. 如果無法得到 target 回傳 -1


## 解題思路
從 0000開始轉, 轉一次可窮舉出 "1000", "9000", "0100", "0900", "0010", "0090", "0001", "0009". **8** 總可能,
再以這八種密碼為基礎, 對每總密碼再轉一下, 窮舉出每個可能
可以抽象成一副圖, 每個節點有8個相鄰的節點, 讓你求出最短距離


## 來源
* https://leetcode.com/problems/open-the-lock/
* https://books.halfrost.com/leetcode/ChapterFour/0700~0799/0752.Open-the-Lock/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0310.Minimum-Height-Trees/main.go

```go
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
```