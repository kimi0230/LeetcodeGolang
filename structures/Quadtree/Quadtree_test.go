package quadtree

import (
	"fmt"
	"testing"
)

func TestQuadtree(t *testing.T) {
	// 建立一個root node，表示整整個空間
	rootNode := NewQuadTreeNode(Rect{0, 0, 100, 100})

	// 插入一些點
	points := []Point{
		{20, 20},
		{30, 30},
		{40, 40},
		{50, 50},
		{60, 60},
	}

	for _, point := range points {
		rootNode.Insert(point)
	}

	// 查詢一個範圍內的點
	queryRect := Rect{25, 25, 55, 55}
	result := rootNode.QueryRange(queryRect)

	fmt.Println("Points within the query range:")
	for _, point := range result {
		fmt.Printf("X: %.2f, Y: %.2f\n", point.X, point.Y)
	}
}

func TestQuadtree2(t *testing.T) {
	// 創建一個簡單的四叉樹
	rootBounds := NewRect(0, 0, 10, 10)
	quadtree := NewQuadTreeNode(rootBounds)

	// 插入一些點
	points := []Point{
		{1, 1},
		{2, 2},
		{7, 7},
		{8, 8},
		{5, 5},
	}

	for _, point := range points {
		quadtree.Insert(point)
	}

	// 測試點是否存在
	for _, point := range points {
		if !quadtree.Bounds.Contains(point) {
			t.Errorf("Point should be inside the bounds")
		}
	}

	// 測試點是否可以查詢到
	queryRect := NewRect(1, 1, 6, 6)
	queryResult := quadtree.QueryRange(queryRect)
	fmt.Println(queryResult)

	if len(queryResult) != 3 {
		t.Errorf("Query should return 3 points")
	}

	// 測試刪除點
	pointToDelete := Point{2, 2}
	quadtree.Delete(pointToDelete)

	// 再次查詢
	queryResult = quadtree.QueryRange(queryRect)

	if len(queryResult) != 2 {
		t.Errorf("Query should return 2 points after deletion")
	}
}
