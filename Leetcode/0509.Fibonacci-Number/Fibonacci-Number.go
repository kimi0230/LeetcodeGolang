// --- Directions
// Print out the n-th entry in the fibonacci series.
// The fibonacci series is an ordering of numbers where
// each number is the sum of the preceeding two.
// For example, the sequence
//  [0, 1, 1, 2, 3, 5, 8, 13, 21, 34]
// forms the first ten entries of the fibonacci series.
// Example:
//   fib(4) === 3

package fibonaccinumber

import "math"

// Fib : iterative 迴圈  O(n) . 空間複雜  O(n). 自底向上的記憶化搜索
func FibIterative(n int) int {
	var result = []int{0, 1}

	for i := 2; i <= n; i++ {
		a := result[i-1]
		b := result[i-2]
		result = append(result, a+b)
	}
	return result[n]
}

// Fibrecur : recursive 遞迴 O(2^n) . 空間複雜  O(n)
func Fibrecur(n int) int {
	if n < 2 {
		return n
	}

	return Fibrecur(n-1) + Fibrecur(n-2)
}

// the memo table
var memo = map[int]int{}

// FibDP : memoization 備忘錄的遞迴
// 比較快!
func FibDP(n int) int {
	// 如果有cached 直接回傳
	if v, vok := memo[n]; vok == true {
		// cached
		return v
	}

	if _, vok := memo[0]; vok == false {
		// 沒cache
		memo[0] = 0
	}

	if _, vok := memo[1]; vok == false {
		// 沒cache
		memo[1] = 1
	}

	for i := 2; i <= n; i++ {
		if _, vok := memo[i]; vok == false {
			// 沒cache
			a := memo[i-1]
			b := memo[i-2]
			memo[i] = a + b
		}
		// fmt.Println(memo)
	}
	return memo[n]
	// return FibDP(n-1) + FibDP(n-2)
}

// 狀態壓縮
func FibDPStateCompression(n int) int {
	if v, vok := memo[n]; vok == true {
		// cached
		return v
	}

	if _, vok := memo[0]; vok == false {
		// 沒cache
		memo[0] = 0
	}

	if _, vok := memo[1]; vok == false {
		// 沒cache
		memo[1] = 1
	}

	prev := 1
	curr := 1

	for i := 2; i <= n; i++ {
		sum := prev + curr
		prev = curr
		curr = sum
	}
	return curr
}

var cache = map[int]int{}

func memoize(fn func(int) int) func(int) int {
	return func(args int) int {
		if _, vok := cache[args]; vok == true {
			return cache[args]
		}

		result := fn(args) // 跑 Fibrecur(args)
		cache[args] = result
		return result
	}
}

// Fibmem : memoization 拉出遞迴function
func Fibmem(n int) int {
	result := memoize(Fibrecur)
	return result(n)
}

// 解法五 矩阵快速幂 时间复杂度 O(log n)，空间复杂度 O(log n)
// | 1 1 | ^ n   = | F(n+1) F(n)   |
// | 1 0 |		   | F(n)	F(n-1) |
func Fib5(N int) int {
	if N <= 1 {
		return N
	}
	var A = [2][2]int{
		{1, 1},
		{1, 0},
	}
	A = matrixPower(A, N-1)
	return A[0][0]
}

func matrixPower(A [2][2]int, N int) [2][2]int {
	if N <= 1 {
		return A
	}
	A = matrixPower(A, N/2)
	A = multiply(A, A)

	var B = [2][2]int{
		{1, 1},
		{1, 0},
	}
	if N%2 != 0 {
		A = multiply(A, B)
	}

	return A
}

func multiply(A [2][2]int, B [2][2]int) [2][2]int {
	x := A[0][0]*B[0][0] + A[0][1]*B[1][0]
	y := A[0][0]*B[0][1] + A[0][1]*B[1][1]
	z := A[1][0]*B[0][0] + A[1][1]*B[1][0]
	w := A[1][0]*B[0][1] + A[1][1]*B[1][1]
	A[0][0] = x
	A[0][1] = y
	A[1][0] = z
	A[1][1] = w
	return A
}

// 解法六 公式法 f(n)=(1/√5)*{[(1+√5)/2]^n -[(1-√5)/2]^n}，用 时间复杂度在 O(log n) 和 O(n) 之间，空间复杂度 O(1)
// 经过实际测试，会发现 pow() 系统函数比快速幂慢，说明 pow() 比 O(log n) 慢
// 斐波那契数列是一个自然数的数列，通项公式却是用无理数来表达的。而且当 n 趋向于无穷大时，前一项与后一项的比值越来越逼近黄金分割 0.618（或者说后一项与前一项的比值小数部分越来越逼近 0.618）。
// 斐波那契数列用计算机计算的时候可以直接用四舍五入函数 Round 来计算。
func Fib6(N int) int {
	var goldenRatio float64 = float64((1 + math.Sqrt(5)) / 2)
	return int(math.Round(math.Pow(goldenRatio, float64(N)) / math.Sqrt(5)))
}

// 解法七 協程版，但是时间特别慢，不推荐，放在这里只是告诉大家，写 LeetCode 算法题的时候，启动 goroutine 特别慢
func Fib7(N int) int {
	return <-fibb(N)
}

func fibb(n int) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)

		if n <= 1 {
			result <- n
			return
		}
		result <- <-fibb(n-1) + <-fibb(n-2)
	}()
	return result
}
