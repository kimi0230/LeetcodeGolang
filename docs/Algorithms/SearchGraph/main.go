package searchgraph

import (
	"LeetcodeGolang/Utility/crud"
	"LeetcodeGolang/structures"
	"fmt"
	"strconv"

	jsoniter "github.com/json-iterator/go"
)

func fetchNeighbours(node int) []int {
	crud := crud.NewCrud("https://hackbear.tv/graph/" + strconv.Itoa(node))
	var result = []int{}

	if got := crud.Get(); got.Error != nil {
		// fmt.Printf("got = %v", got)
	} else {
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		json.Unmarshal([]byte(got.Response), &result)
	}
	return result
}

/*
     2
   /   \
  /     \
 1 - 3 - 5 - 7
  \        /
   \      6
    \   /
      4
*/

// 1 -> [2, 3, 4]
// 2 -> [1, 5]
// 3 -> [1, 5]
// 4 -> [1, 6]
// 5 -> [2, 3, 7]
// 6 -> [4, 7]
// 7 -> [5, 6]
// print : 1,2,3,4,5,6,7
// Also this is valid : 1,4,3,2,6,5,7

// You will be working on this part
/*
時間複雜度是O(V+E)，其中V是圖中節點的數量，E是圖中邊的數量
*/
func SearchGraph(start int) {
	queue := structures.NewQueue()
	queue.Push(start)
	visit := make(map[int][]int)
	for queue.Len() > 0 {
		node := queue.Pop()
		if _, ok := visit[node]; !ok {
			fmt.Printf("%d ", node)
			neighours := fetchNeighbours(node)
			visit[node] = neighours
			for _, neighour := range neighours {
				if _, ok := visit[neighour]; !ok {
					queue.Push(neighour)
				}
			}
		}

	}
}

func SearchGraph2(start int) {
	queue := structures.NewQueue()
	queue.Push(start)
	visited := make(map[int]bool)
	for queue.Len() > 0 {
		node := queue.Pop()
		if !visited[node] {
			fmt.Printf("%d ", node)
			visited[node] = true
			neighbors := fetchNeighbours(node)
			for _, neighbor := range neighbors {
				if !visited[neighbor] {
					queue.Push(neighbor)
				}
			}
		}
	}
}
