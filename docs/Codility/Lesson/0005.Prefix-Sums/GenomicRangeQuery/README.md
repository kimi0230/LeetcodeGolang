# [GenomicRangeQuery](https://app.codility.com/programmers/lessons/5-prefix_sums/genomic_range_query/)
Find the minimal nucleotide(核苷酸) from a range of sequence DNA.

A DNA sequence can be represented as a string consisting of the letters A, C, G and T, which correspond to the types of successive nucleotides in the sequence. Each nucleotide has an impact factor, which is an integer. Nucleotides of types A, C, G and T have impact factors of 1, 2, 3 and 4, respectively. You are going to answer several queries of the form: What is the minimal impact factor of nucleotides contained in a particular part of the given DNA sequence?

The DNA sequence is given as a non-empty string S = S[0]S[1]...S[N-1] consisting of N characters. There are M queries, which are given in non-empty arrays P and Q, each consisting of M integers. The K-th query (0 ≤ K < M) requires you to find the minimal impact factor of nucleotides contained in the DNA sequence between positions P[K] and Q[K] (inclusive).

For example, consider string S = CAGCCTA and arrays P, Q such that:

    P[0] = 2    Q[0] = 4
    P[1] = 5    Q[1] = 5
    P[2] = 0    Q[2] = 6
The answers to these M = 3 queries are as follows:

The part of the DNA between positions 2 and 4 contains nucleotides G and C (twice), whose impact factors are 3 and 2 respectively, so the answer is 2.
The part between positions 5 and 5 contains a single nucleotide T, whose impact factor is 4, so the answer is 4.
The part between positions 0 and 6 (the whole string) contains all nucleotides, in particular nucleotide A whose impact factor is 1, so the answer is 1.
Write a function:

func Solution(S string, P []int, Q []int) []int

that, given a non-empty string S consisting of N characters and two non-empty arrays P and Q consisting of M integers, returns an array consisting of M integers specifying the consecutive answers to all queries.

Result array should be returned as an array of integers.

For example, given the string S = CAGCCTA and arrays P, Q such that:

    P[0] = 2    Q[0] = 4
    P[1] = 5    Q[1] = 5
    P[2] = 0    Q[2] = 6
the function should return the values [2, 4, 1], as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
M is an integer within the range [1..50,000];
each element of arrays P, Q is an integer within the range [0..N − 1];
P[K] ≤ Q[K], where 0 ≤ K < M;
string S consists only of upper-case English letters A, C, G, T.
Copyright 2009–2021 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.

## 題目大意
CAGCCTA
A 代表1, C 代表2, G 代表3 ,T 代表4
K=0: P[0]=2, Q[0]=4 之間的核苷酸 DNA(GCC)因素分別是3和2, 最小的就是2.
K=1: P[1]=5, Q[1]=5 DNA(T),最小的是4.
K=2: P[2]=0, Q[2]=6 DNA(CAGCCTA),最小的是1.

## 解題思路

## 來源
https://app.codility.com/programmers/lessons/5-prefix_sums/genomic_range_query/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Codility/Lesson/0005.Prefix-Sums/GenomicRangeQuery/GenomicRangeQuery.go


```go
package GenomicRangeQuery

func Solution(S string, P []int, Q []int) []int {
	A, C, G, T := prefixSums(S)
	result := make([]int, len(P))
	/*
		// fmt.Println("A: ", A)
		// fmt.Println("C: ", C)
		// fmt.Println("G: ", G)
		// fmt.Println("T: ", T)
		idx  0 1 2 3 4 5 6 7
		S:  [C A G C C T A]
		A:  [0 0 1 1 1 1 1 2]
		C:  [0 1 1 1 2 3 3 3]
		G:  [0 0 0 1 1 1 1 1]
		T:  [0 0 0 0 0 0 1 1]
		P:  [2 5 0]
		Q:  [4 5 6]
	*/

	for k, _ := range P {
		// 判斷 A[end of slice]-A[Begin of Slice]是否大於零 即可判斷是否 A 出現過
		if A[Q[k]+1]-A[P[k]] > 0 {
			result[k] = 1
		} else if C[Q[k]+1]-C[P[k]] > 0 {
			result[k] = 2
		} else if G[Q[k]+1]-G[P[k]] > 0 {
			result[k] = 3
		} else if T[Q[k]+1]-T[P[k]] > 0 {
			result[k] = 4
		}
	}

	return result
}

// 數算從開始到每個固定索引的A,C,G,T個數. 開頭插入0
func prefixSums(S string) ([]int, []int, []int, []int) {
	n := len(S)
	A := make([]int, n+1)
	C := make([]int, n+1)
	G := make([]int, n+1)
	T := make([]int, n+1)

	for i := 1; i < n+1; i++ {
		s := string(S[i-1])
		A[i] = A[i-1]
		C[i] = C[i-1]
		G[i] = G[i-1]
		T[i] = T[i-1]

		switch s {
		case "A":
			A[i]++
		case "C":
			C[i]++
		case "G":
			G[i]++
		case "T":
			T[i]++
		}

	}

	return A, C, G, T
}

func inLoop(arr string, s string) bool {
	for _, v := range arr {
		if string(v) == s {
			return true
		}
	}
	return false
}

// O(N * M)
func SolutionBurst(S string, P []int, Q []int) []int {
	result := make([]int, len(P))
	for i := 0; i < len(P); i++ {
		tmp := S[P[i] : Q[i]+1]
		if inLoop(tmp, "A") {
			result[i] = 1
		} else if inLoop(tmp, "C") {
			result[i] = 2
		} else if inLoop(tmp, "G") {
			result[i] = 3
		} else if inLoop(tmp, "T") {
			result[i] = 4
		}
	}
	return result
}

/*
def solutionBySlice(S, P, Q):
  result = []
  length = len(P)
  for i in range(length):
    temp = (S[P[i]:Q[i]+1])
    if "A" in temp:
      result.append(1)
    elif "C" in temp:
      result.append(2)
    elif "G" in temp:
      result.append(3)
    elif "T" in temp:
      result.append(4)
  return result
*/
```