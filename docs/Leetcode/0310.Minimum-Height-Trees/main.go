package minimumheighttrees

import "fmt"

func FindMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	graph := make(map[int][]int)
	degree := make(map[int]int, n) //出度: 出度是指從該節點（頂點）出發的邊的條數
	// 建立對應關係
	for _, v := range edges {
		if _, ok := graph[v[0]]; ok {
			graph[v[0]] = append(graph[v[0]], v[1])
		} else {
			// 沒找到
			graph[v[0]] = []int{v[1]}
		}

		if _, ok := graph[v[1]]; ok {
			graph[v[1]] = append(graph[v[1]], v[0])
		} else {
			// 沒找到
			graph[v[1]] = []int{v[0]}
		}
		degree[v[0]]++
		degree[v[1]]++
	}
	fmt.Printf("graph: %+v \n", graph)
	fmt.Printf("degree: %+v \n", degree)

	var queue []int
	for key, value := range degree { //先找到出度為1的點, 加到Queue
		if value == 1 {
			queue = append(queue, key)
		}
	}
	fmt.Printf("queue: %+v \n", queue)

	leaves := []int{}
	for len(queue) != 0 {
		leaves = leaves[0:0]
		size := len(queue)

		for i := 0; i < size; i++ {
			leaves = append(leaves, queue[i])
			for _, neighbor := range graph[queue[i]] { // 將該點鄰接出度-1, 等於1時加入Queue
				degree[neighbor]--
				if degree[neighbor] == 1 {
					queue = append(queue, neighbor)
				}
			}
		}
		queue = queue[size:]
	}
	return leaves
}

/*
n: 4,	edges: [][]int{{1, 0}, {1, 2}, {1, 3}},
graph: map[0:[1] 1:[0 2 3] 2:[1] 3:[1]]
           0 (degree: 1)
          /
         1 (3)
        / \
    2(1)   3(1)

degree: map[0:1 1:3 2:1 3:1]
queue: [2 3 0] (先degree=1)
*/
