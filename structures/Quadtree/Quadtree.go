package quadtree

// 定義一個二維點的結構
type Point struct {
	X, Y float64
}

// 定義一個矩形結構
type Rect struct {
	// 左上 (X1,Y1) ; 右下 (X2,Y2)
	X1, Y1, X2, Y2 float64
}

// 建立一個新的矩形
func NewRect(x1, y1, x2, y2 float64) Rect {
	return Rect{X1: x1, Y1: y1, X2: x2, Y2: y2}
}

// 計算矩形的寬度
func (r Rect) Width() float64 {
	return r.X2 - r.X1
}

// 計算矩形的高度
func (r Rect) Height() float64 {
	return r.Y2 - r.Y1
}

// 檢查矩形是否包含某一點
func (r Rect) Contains(point Point) bool {
	return point.X >= r.X1 && point.X <= r.X2 && point.Y >= r.Y1 && point.Y <= r.Y2
}

// 檢查矩形是否與另一個矩形相交
func (r Rect) Intersects(other Rect) bool {
	return r.X1 < other.X2 && r.X2 > other.X1 && r.Y1 < other.Y2 && r.Y2 > other.Y1
}

// 定義四叉樹節點結構
type QuadTreeNode struct {
	Bounds         Rect
	Points         []Point
	NW, NE, SW, SE *QuadTreeNode
}

// 創建一個新的四叉樹節點
func NewQuadTreeNode(bounds Rect) *QuadTreeNode {
	return &QuadTreeNode{
		Bounds: bounds,
		Points: make([]Point, 0),
	}
}

// 插入一個點到四叉樹
func (n *QuadTreeNode) Insert(point Point) {
	if !n.Bounds.Contains(point) {
		return
	}

	if len(n.Points) < 4 {
		n.Points = append(n.Points, point)
		return
	}

	if n.NW == nil {
		n.Subdivide()
	}

	n.NW.Insert(point)
	n.NE.Insert(point)
	n.SW.Insert(point)
	n.SE.Insert(point)
}

// 將節點分為四個子結點
// |NW|NE|
// |SW|SE|
func (n *QuadTreeNode) Subdivide() {
	xMid := (n.Bounds.X1 + n.Bounds.X2) / 2
	yMid := (n.Bounds.Y1 + n.Bounds.Y2) / 2

	n.NW = NewQuadTreeNode(Rect{n.Bounds.X1, n.Bounds.Y1, xMid, yMid})
	n.NE = NewQuadTreeNode(Rect{xMid, n.Bounds.Y1, n.Bounds.X2, yMid})
	n.SW = NewQuadTreeNode(Rect{n.Bounds.X1, yMid, xMid, n.Bounds.Y2})
	n.SE = NewQuadTreeNode(Rect{xMid, yMid, n.Bounds.X2, n.Bounds.Y2})
}

// 從節點中删除一個點
func (n *QuadTreeNode) Delete(point Point) {
	if !n.Bounds.Contains(point) {
		return
	}

	for i, p := range n.Points {
		if p == point {
			// 從節點中刪除點
			n.Points = append(n.Points[:i], n.Points[i+1:]...)
			return
		}
	}

	if n.NW == nil {
		return
	}

	n.NW.Delete(point)
	n.NE.Delete(point)
	n.SW.Delete(point)
	n.SE.Delete(point)
}

type QueryContext struct {
	Visited map[Point]bool
}

func NewQueryContext() *QueryContext {
	return &QueryContext{Visited: make(map[Point]bool)}
}

// 查詢一個範圍内的點
func (n *QuadTreeNode) QueryRange(rangeRect Rect) []Point {
	visited := NewQueryContext()
	return n.queryRange(rangeRect, visited)
}

// 查詢一個範圍内的點
func (n *QuadTreeNode) queryRange(rangeRect Rect, context *QueryContext) []Point {
	result := make([]Point, 0)

	// 判斷搜尋範圍(range)是否與當前領域有相交，沒有就直接跳掉，以節省效率
	if !n.Bounds.Intersects(rangeRect) {
		return result
	}

	// 如果搜尋範圍跟當前領域相交，檢查該領域有多少物體包含在搜尋範圍中，並增加至 result 中
	for _, point := range n.Points {
		if rangeRect.Contains(point) && !context.Visited[point] {
			result = append(result, point)
			context.Visited[point] = true
		}
	}

	if n.NW == nil {
		return result
	}

	result = append(result, n.NW.queryRange(rangeRect, context)...)
	result = append(result, n.NE.queryRange(rangeRect, context)...)
	result = append(result, n.SW.queryRange(rangeRect, context)...)
	result = append(result, n.SE.queryRange(rangeRect, context)...)

	return result
}
