## Leetcode in Golang
[![Build latest tag](https://github.com/kimi0230/LeetcodeGolang/actions/workflows/releace.yml/badge.svg)](https://github.com/kimi0230/LeetcodeGolang/actions/workflows/releace.yml) [![Build my gitbook and deploy to gh-pages](https://github.com/kimi0230/LeetcodeGolang/actions/workflows/build.yml/badge.svg)](https://github.com/kimi0230/LeetcodeGolang/actions/workflows/build.yml) [![pages-build-deployment](https://github.com/kimi0230/LeetcodeGolang/actions/workflows/pages/pages-build-deployment/badge.svg)](https://github.com/kimi0230/LeetcodeGolang/actions/workflows/pages/pages-build-deployment) ![](https://steel-quark-crabapple.glitch.me/badge?page_id=kimi0230.LeetCodeGolang)

[Leetcode](https://leetcode.com/), [Codility](https://app.codility.com/programmers/) , [GeekforGeeks algorithms](https://www.geeksforgeeks.org/fundamentals-of-algorithms/?ref=shm) exercises written in Golang.

https://kimi0230.github.io/LeetcodeGolang/

---
## [leetcode](https://leetcode.com/) Content
- [Leetcode in Golang](#leetcode-in-golang)
- [leetcode Content](#leetcode-content)
  - [Data Structure](#data-structure)
    - [Array \& String](#array--string)
    - [Matrix](#matrix)
    - [Linked List](#linked-list)
    - [HashSet \& HashMap](#hashset--hashmap)
    - [Stack \& Queue](#stack--queue)
    - [Heap \& Priority Queue](#heap--priority-queue)
      - [Heap 有幾個特色：](#heap-有幾個特色)
      - [heap sort](#heap-sort)
    - [Disjoint Set Union](#disjoint-set-union)
    - [Trie](#trie)
    - [Binary Indexed Tree](#binary-indexed-tree)
    - [Design Data Structures](#design-data-structures)
  - [Algorithm](#algorithm)
    - [Greedy](#greedy)
    - [Sort](#sort)
    - [Multiple Pointers](#multiple-pointers)
    - [Backtracking (回溯法). DFS](#backtracking-回溯法-dfs)
    - [DFS \& BFS](#dfs--bfs)
    - [Dynamic Programming](#dynamic-programming)
    - [Sliding Window](#sliding-window)
    - [Sweep Line](#sweep-line)
    - [Rolling Sum](#rolling-sum)
    - [Two Pointers](#two-pointers)
    - [Bit Manipulation∂](#bit-manipulation)
    - [Union Find](#union-find)
    - [Breadth First Search](#breadth-first-search)
    - [Binary Search](#binary-search)
    - [Minimax](#minimax)
  - [Graph](#graph)
    - [Graph](#graph-1)
    - [Topological Sort](#topological-sort)
    - [Tree](#tree)
    - [Tree Traversal](#tree-traversal)
    - [Binary Search Tree](#binary-search-tree)
    - [Compputational Geometry](#compputational-geometry)
  - [Selected Topics](#selected-topics)
    - [Mathematics](#mathematics)
    - [Random](#random)
    - [Bitwise Manipulation](#bitwise-manipulation)
- [GeeksforGeeks Content](#geeksforgeeks-content)
- [Codility Content](#codility-content)
- [Reference](#reference)

### Data Structure

#### Array & String

| No.                                                                                                            |                                                                Title                                                                |                                                         Solution                                                          | Difficulty | Time          | Space  | Topic                         |
|----------------------------------------------------------------------------------------------------------------|:-----------------------------------------------------------------------------------------------------------------------------------:|:-------------------------------------------------------------------------------------------------------------------------:|------------|---------------|--------|-------------------------------|
| [0001](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0001.Two-Sum)                                        |                                          [Two Sum](https://leetcode.com/problems/two-sum/)                                          |                    [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0001.Two-Sum)                     | Easy       | O(n)          | O(n)   | Array                         |
| [0003](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0003.Longest-Substring-Without-Repeating-Characters) |   [Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)   | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0003.Longest-Substring-Without-Repeating-Characters) | Medium     | O(n)          | O(1)   | Array, Sliding Window         |
| [0015](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0015.3Sum)                                           |                                            [3 Sum](https://leetcode.com/problems/3sum/)                                             |                      [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0015.3Sum)                      | Medium     | O(n^2)        | O(n)   | Array                         |
| [0027](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0027.Remove-Element)                                 |                                   [Remove Element](https://leetcode.com/problems/remove-element/)                                   |                 [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0027.Remove-Element)                 | Easy       | O(n)          | O(1)   | Array                         |
| [0035](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0035.Search-Insert-Position)                         |                           [Search Insert Position](https://leetcode.com/problems/search-insert-position/)                           |             [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0035.Search-Insert-Position)             | Easy       | O(n), O(logn) | O(1)   | Array                         |
| [0049](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0049.Group-Anagrams)                                 |                               [Search Insert Position](https://leetcode.com/problems/group-anagrams/)                               |                 [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0049.Group-Anagrams)                 | Medium     | O(kn)         | O(kn)  | Array                         |
| [0059](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0059.Spiral-Matrix-II)                               |                                 [Spiral Matrix II](https://leetcode.com/problems/spiral-matrix-ii/)                                 |                [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0059.Spiral-Matrix-II)                | Medium     | O(n)          | O(n^2) | Array                         |
| [0088](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0088.Merge-Sorted-Array)                             |                               [Merge Sorted Array](https://leetcode.com/problems/merge-sorted-array/)                               |               [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0088.Merge-Sorted-Array)               | Easy       | O(n)          | O(1)   | Array                         |
| [0217](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0217.Contains-Duplicate)                             |                      [0217.Contains Duplicate](https://leetcode.com/problems/contains-duplicate/description/)                       |               [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0217.Contains-Duplicate)               | Easy       | O(n)          | O(n)   | Array                         |
| [0242](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0242.Valid-Anagram)                                  |                           [0242.Valid Anagram](https://leetcode.com/problems/valid-anagram/description/)                            |                 [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0242.Valid-Anagram)                  | Easy       | O(n)          | O(n)   | Array                         |
| [0409](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0409.Longest-Palindrome)                             |                            [409. Longest Palindrome](https://leetcode.com/problems/longest-palindrome/)                             |               [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0409.Longest-Palindrome)               | Easy       | O(n)          | O(1)   | Array                         |
| [0380](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0380.Insert-Delete-GetRandom-O1)                     |             [0380.Insert Delete GetRandom O(1)](https://leetcode.com/problems/insert-delete-getrandom-o1/description/)              |           [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0380.Insert-Delete-GetRandom-O1)           | Medium     | O(1)          | O(n)   | Array                         |
| [0381](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0381.Insert-Delete-GetRandom-O1-Duplicates-allowed)  | [0381.Insert Delete GetRandom O(1) Duplicates allowed](https://leetcode.com/problems/insert-delete-getrandom-o1-duplicates-allowed) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0381.Insert-Delete-GetRandom-O1-Duplicates-allowed)  | Medium     | O(1)          | O(n)   | Array                         |
| [0412](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0412.Fizz-Buzz)                                      |                               [0412.Fizz Buzz](https://leetcode.com/problems/fizz-buzz/description/)                                |                   [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0412.Fizz-Buzz)                    | Easy       | O(n)          | O(n)   | Array, string                 |
| [1195](https://kimi0230.github.io/LeetcodeGolang/Leetcode/1195.Fizz-Buzz-Multithreaded)                        |                 [1195.Fizz Buzz Multithreaded](https://leetcode.com/problems/fizz-buzz-multithreaded/description/)                  |            [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/1195.Fizz-Buzz-Multithreaded)             | Medium     | O(n)          |        | Array, string,Concurrency     |
| [0238](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0238.Product-of-Array-Except-Self)                   |            [238. Product of Array Except Self](https://leetcode.com/problems/product-of-array-except-self/description/)             |          [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0238.Product-of-Array-Except-Self)          | Medium     | O(n)          |        | Array, string, Prefix Sum     |
| [0128](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0128.Longest-Consecutive-Sequence)                   |            [128. Longest Consecutive Sequence](https://leetcode.com/problems/longest-consecutive-sequence/description/)             |          [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0128.Longest-Consecutive-Sequence)          | Medium     | O(n)          | O(n)   | Array, Hash Table, Union Find |

---

#### Matrix

---
#### Linked List

| No.                                                                                               |                                                Title                                                |                                                  Solution                                                   | Difficulty | Time        | Space | Topic                     |
|---------------------------------------------------------------------------------------------------|:---------------------------------------------------------------------------------------------------:|:-----------------------------------------------------------------------------------------------------------:|------------|-------------|-------|---------------------------|
| [0019](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0019.Remove-Nth-Node-From-End-of-List/) | [Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0019.Remove-Nth-Node-From-End-of-List) | Medium     | O(n)        | O(1)  | Linked List, Two Pointers |
| [0141](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0141.Linked-List-Cycle/)                |                [Linked List Cycle](https://leetcode.com/problems/linked-list-cycle)                 |        [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0141.Linked-List-Cycle)         | Easy       | O(n)        | O(1)  | Linked List, Two Pointers |
| [0142](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0142.Linked-List-CycleII/)              |             [Linked List Cycle II](https://leetcode.com/problems/linked-list-cycle-ii/)             |       [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0142.Linked-List-CycleII)        | Medium     | O(n)        | O(1)  | Linked List, Two Pointers |
| [0203](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0203.Remove-Linked-List-Elements)       |      [Remove Linked List Elements](https://leetcode.com/problems/remove-linked-list-elements/)      |   [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0203.Remove-Linked-List-Elements)    | Easy       | O(n)        | O(1)  | Linked List               |
| [0206](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0206.Reverse-Linked-List)               |              [Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/)              |              [Go](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0206.Reverse-Linked-List)              | Easy       | O(n)        | O(1)  | Linked List               |
| [0876](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0876.Middle-of-the-Linked-List/)        |        [Middle of the Linked List](https://leetcode.com/problems/middle-of-the-linked-list/)        |    [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0876.Middle-of-the-Linked-List)     | Easy       |             |       | Linked List, Two Pointers |
| [0021](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0021.Merge-Two-Sorted-Lists/)           |     [Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/description/)     |      [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0021.Merge-Two-Sorted-Lists)      | Easy       | O(log n)    | O(1)  | Linked List               |
| [0002](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0002.Add-Two-Numbers/)                  |                  [Add Two Number](https://leetcode.com/problems/add-two-numbers//)                  |         [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0002.Add-Two-Numbers)          | Medium     | O(max(m,n)) | O(1)  | Linked List               |

---

#### HashSet & HashMap

---

#### Stack & Queue

| No.                                                                                            |                                             Title                                             |                                                 Solution                                                 | Difficulty | Time | Space | Topic |
|------------------------------------------------------------------------------------------------|:---------------------------------------------------------------------------------------------:|:--------------------------------------------------------------------------------------------------------:|------------|------|-------|-------|
| [0020](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0020.Valid-Parentheses/)             |             [Valid Parentheses](https://leetcode.com/problems/valid-parentheses/)             |       [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0020.Valid-Parentheses)       | Easy       | O(n) | O(n)  | Stack |
| [0094](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0094.Binary-Tree-Inorder-Traversal/) | [Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0094.Binary-Tree-Inorder-Traversal) | Medium     | O(n) | O(1)  | Stack |


#### Heap & Priority Queue

Heap 總是能讓整棵樹當中最大或最小值維持在root節點上

##### Heap 有幾個特色：
* 常見架構是像 binary tree 那樣
* 保持 balanced
* max heap 的 root 是最大值；min heap 的 root 則是最小值
* 雖然是 tree，卻很適合放在 array 中處理

##### heap sort
根據定義 heap 的 root 一定是最大(假設是 max heap)，也就是說，無序數列經過 heapify 再作 n 次 root deletion 取出最大值，就可以得到排序的結果。
最後就得到 heap sort 的 worst case 時間複雜度 O(nlogn) 的結果。
可是 quick sort 的 worst case 時間複雜度是 O(n²)，怎麼 quick sort 的時間複雜度比較糟糕卻比較受歡迎？
google 的結果是說 heap sort 比較不利於 caching 對於 spatial locality 機制，蠻有道理的阿。
* https://www.zhihu.com/question/23873747
* https://rust-algo.club/sorting/heapsort/index.html

| No.                                                                                              |                                               Title                                               |                                                  Solution                                                  | Difficulty | Time             | Space | Topic                            |
|--------------------------------------------------------------------------------------------------|:-------------------------------------------------------------------------------------------------:|:----------------------------------------------------------------------------------------------------------:|------------|------------------|-------|----------------------------------|
| [0703](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0703.Kth-Largest-Element-in-a-Stream/) | [Kth Largest Element in a Stream](https://leetcode.com/problems/kth-largest-element-in-a-stream/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0703.Kth-Largest-Element-in-a-Stream) | Easy       | O(K + (N-K)logK) | O(k)  | Heap, Priority Queue             |
| [1046](https://kimi0230.github.io/LeetcodeGolang/Leetcode/1046.Last-Stone-Weight/)               |               [Last Stone Weight](https://leetcode.com/problems/last-stone-weight/)               |        [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/1046.Last-Stone-Weight)        | Easy       | O(nlogn)         | O(n)  | Heap, Priority Queue             |
| [0347](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0347.Top-K-Frequent-Elements/)         |   [Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/description/)   |     [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0347.Top-K-Frequent-Elements)     | Medium     | O(Nlog⁡k)         | O(n)  | Heap, Priority Queue, Quick Sort |


#### Disjoint Set Union

#### Trie

#### Binary Indexed Tree
 
#### Design Data Structures

---

### Algorithm

#### Greedy

---

#### Sort
| No.                                                                                              |                                               Title                                               |                                                  Solution                                                  | Difficulty | Time | Space   | Topic |
|--------------------------------------------------------------------------------------------------|:-------------------------------------------------------------------------------------------------:|:----------------------------------------------------------------------------------------------------------:|------------|------|---------|-------|
| [0075](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0075.Sort-Colors/)                     |                     [Sort Colors](https://leetcode.com/problems/sort-colors/)                     |           [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0075.Sort-Colors)           | Medium     | O(n) | O(1)    | Sort  |
| [0215](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0215.Kth-Largest-Element-in-an-Array/) | [Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0215.Kth-Largest-Element-in-an-Array) | Medium     | O(n) | O(logn) | Sort  |

#### Multiple Pointers

---

#### Backtracking (回溯法). DFS
DFS. 解決一個回溯問題, 實際上就是一個**決策樹**的遍歷過程.
算是一個暴力的窮舉算法
1. 路徑：也就是已經做出的選擇。
2. 選擇列表：也就是你當前可以做的選擇。
3. 結束條件：也就是到達決策樹底層，無法再做選擇的條件。
4. https://www.bilibili.com/video/BV1P5411N7Xc


```python
result = []
def backtrack(路徑, 選擇列表):
    if 滿足結束條件:
        result.add(路徑)
        return
    
    for 選擇 in 選擇列表:
        做選擇(前序)
        backtrack(路徑, 選擇列表)
        撤銷選擇(後序)
```

| No.                                                                           |                                Title                                 |                                        Solution                                         | Difficulty | Time   | Space | Topic        |
|-------------------------------------------------------------------------------|:--------------------------------------------------------------------:|:---------------------------------------------------------------------------------------:|------------|--------|-------|--------------|
| [0046](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0046.Permutations/) | [Permutations (全排列)](https://leetcode.com/problems/permutations/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0046.Permutations) | Medium     | O(n)   | O(n)  | Backtracking |
| [0078](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0078.Subsets/)      |          [Subsets](https://leetcode.com/problems/subsets/)           |   [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0078.Subsets)    | Medium     | O(n^2) | O(n)  | Backtracking |

---

#### DFS & BFS

找最短路徑用**BFS**, 
其他時用**DFS**用得多一些, 因為遞迴較好寫

假設有棵滿的二叉樹,節點數為 N. 對DFS來說空間複雜度就是遞迴, 最壞的情況就是樹的高度 O(log N)
BFS算法, Queue每次都會存二叉樹一層的節點, 最壞的情況下空間複雜度應該就是樹的最下層的數量, 也就是 N/2. 空間複雜度 O(N)

DFS（深度優先搜索）通常使用堆棧（Stack）來實現。在DFS中，您首先處理一個節點，然後將其子節點按某種順序推入堆棧中，接著繼續處理堆棧頂部的節點，直到堆棧為空。
BFS（廣度優先搜索）則使用隊列（Queue）來實現。在BFS中，您首先處理一個節點，然後將其子節點按某種順序排隊，接著繼續處理隊列的前端節點，直到隊列為空。

| No.                                                                                 |                                  Title                                  |                                           Solution                                            | Difficulty | Time   | Space  | Topic     |
|-------------------------------------------------------------------------------------|:-----------------------------------------------------------------------:|:---------------------------------------------------------------------------------------------:|------------|--------|--------|-----------|
| [0695](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0695.Max-Area-of-Island/) | [Max Area of Island](https://leetcode.com/problems/max-area-of-island/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0695.Max-Area-of-Island) | Medium     | O(m*n) | O(m*n) | DFS & BFS |
| [0733](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0733.Flood-Fill/)         |         [Flood Fill](https://leetcode.com/problems/flood-fill)          |     [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0733.Flood-Fill)     | Easy       | O(m*n) | O(m*n) | DFS & BFS |

---
#### Dynamic Programming
動態規劃問題的一般形式就是**求最值**, **最長**遞增子序列, **最小**編輯距離等. 核心問題是窮舉
1. 重疊子問題
   1. memory table
   2. DP table
2. 最優子結構
3. 狀態轉移方程式
   1. 這問題的 base case (最簡單情況) 是什麼?
   2. 這問題有什麼**狀態**
   3. 對於每個狀態, 可以做出什麼**選擇**, 使得狀態發生改變
   4. 如何定義 dp 數組/函數的含義來表現**狀態**和**選擇**?

| 替換 /跳過 <br> dp[i-1][j-1] | 刪除 <br> dp[i-1][j] |
|------------------------------|----------------------|
| 插入 <br> dp[i][j-1]         | dp[i][j]             |

```python
# 初始化 base case
dp[0][0][...] = base
# 進行狀態轉移
for 狀態1 in 狀態1的所有取值：
    for 狀態2 in 狀態2的所有取值：
        for ...
            dp[狀態1][狀態2][...] = 求最值(選擇1，選擇2...)
```

| No.                                                                                             |                                              Title                                              |                                                 Solution                                                  | Difficulty | Time                          | Space    | Topic               |
|-------------------------------------------------------------------------------------------------|:-----------------------------------------------------------------------------------------------:|:---------------------------------------------------------------------------------------------------------:|------------|-------------------------------|----------|---------------------|
| [0053](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0053.Maximum-Subarray/)               |               [Maximum Subarray](https://leetcode.com/problems/maximum-subarray/)               |        [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0053.Maximum-Subarray)        | Easy       | O(n)                          | O(n)     | Dynamic Programming |
| [0072](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0072.Edit-Distance/)                  |               [0072. Edit Distance](https://leetcode.com/problems/edit-distance/)               |         [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0072.Edit-Distance)          | Hard       |                               |          | Dynamic Programming |
| [0300](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0300.Longest-Increasing-Subsequence/) | [Longest-Increasing-Subsequence](https://leetcode.com/problems/longest-increasing-subsequence/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0300.Longest-Increasing-Subsequence) | Medium     | 方法一:O(n^2) 方法二:O(nlogn) | O(n)     | Dynamic Programming |
| [0322](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0322.Coin-Change/)                    |                    [Coin Change](https://leetcode.com/problems/coin-change/)                    |          [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0322.Coin-Change)           | Medium     | O(nm)                         | O(n)     | Dynamic Programming |
| [0354](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0354.Russian-Doll-Envelopes/)         |         [Russian Doll Envelope](https://leetcode.com/problems/russian-doll-envelopes/)          |     [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0354.Russian-Doll-Envelopes)     | Hard       |                               |          | Dynamic Programming |
| [0509](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0509.Fibonacci-Number/)               |               [Fibonacci Number](https://leetcode.com/problems/fibonacci-number/)               |        [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0509.Fibonacci-Number)        | Easy       | 很多解法                      | 很多解法 | Dynamic Programming |
| [0070](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0070.Climbing-Stairs/)                |             [0070.Climbing Stairs](https://leetcode.com/problems/climbing-stairs/)              |        [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0070.Climbing-Stairs)         | Easy       | O(n)                          | O(n)     | Dynamic Programming |
| [0746](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0746.Min-Cost-Climbing-Stairs/)       |    [0746.Min Cost Climbing Stairs](https://leetcode.com/problems/min-cost-climbing-stairs/)     |    [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0746.Min-Cost-Climbing-Stairs)    | Easy       | O(n)                          | O(1)     | Dynamic Programming |

---

#### Sliding Window
維護一個窗口, 不斷滑動
```c++
void slidingWindow(string s, string t){
    unordered map<char,int>need, window;
    for (char c:t) need[c++]
    int left = 0 , right = 0
    int valid = 0

    // 先移動 right 再移動 left. 直到right到達 string的末端
    while(right < s.size()){
        // c是將移入窗口的字符
        char c = s[right]
        // 右移窗口 
        right++
        // 進行窗口內數據的一系列更新
        // ...

        /*** 用來debug 輸出位置 ***/
        printf("window: [%d, %d)\n",left,right)
        /************************/

        // 判斷左側窗口是否收縮
        while(window needs shrink){
           // d是將移出窗口的字符
           // 左移窗口 
            left++
            // 進行窗口內數據的一系列更新
            // ...
        }
    }
}
```
| No.                                                                                            |                                             Title                                             |                                                 Solution                                                 | Difficulty | Time                      | Space              | Topic          |
|------------------------------------------------------------------------------------------------|:---------------------------------------------------------------------------------------------:|:--------------------------------------------------------------------------------------------------------:|------------|---------------------------|--------------------|----------------|
| [0209](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0209.Minimum-Size-Subarray-Sum/)     |     [Minimum Size Subarray Sum](https://leetcode.com/problems/minimum-size-subarray-sum/)     |   [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0209.Minimum-Size-Subarray-Sum)   | Medium     | O(n^2) / O(n) / O(nlog n) | O(1) / O(1) / O(n) | Sliding Window |
| [0438](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0438.Find-All-Anagrams-in-a-String/) | [Find All Anagrams in a String](https://leetcode.com/problems/find-all-anagrams-in-a-string/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0438.Find-All-Anagrams-in-a-String) | Medium     | O(n)                      | O(1)               | Sliding Window |
| [0567](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0567.Permutation-in-String/)         |         [Permutation in String](https://leetcode.com/problems/permutation-in-string/)         |     [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0567.Permutation-in-String)     | Medium     | O(n)                      | O(1)               | Sliding Window |

---
#### Sweep Line


---
#### Rolling Sum

---
#### Two Pointers
只要**array有序**, 就應該想到雙指針技巧
分為兩類 1. "快,慢指針" 2. "左,右指針"
   1. 快,慢指針: 主要解決 linkedlist 問題, 典型的判斷 linkedlist 是否包含環
   2. 左,右指針: 主要解決array(或 string)中的問題, 如二分搜尋.

https://labuladong.gitee.io/algo/2/21/57/

| No.                                                                                               |                                            Title                                             |                                                  Solution                                                   | Difficulty | Time | Space | Topic                       |
|---------------------------------------------------------------------------------------------------|:--------------------------------------------------------------------------------------------:|:-----------------------------------------------------------------------------------------------------------:|------------|------|-------|-----------------------------|
| [0019](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0019.Remove-Nth-Node-From-End-of-List/) | [Remove Nth Node From End of List](https://leetcode.com/problems/middle-of-the-linked-list/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0019.Remove-Nth-Node-From-End-of-List) | Medium     | O(n) | O(1)  | Linked List, Two Pointers   |
| [0141](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0141.Linked-List-Cycle/)                |             [Linked List Cycle](https://leetcode.com/problems/linked-list-cycle)             |        [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0141.Linked-List-Cycle)         | Easy       | O(n) | O(1)  | Linked List, Two Pointers   |
| [0283](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0283.Move-Zeroes/)                      |                  [Move Zeroes](https://leetcode.com/problems/move-zeroes/)                   |           [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0283.Move-Zeroes)            | Easy       | O(n) | O(1)  | Two Pointers                |
| [0142](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0142.Linked-List-CycleII/)              |         [Linked List Cycle II](https://leetcode.com/problems/linked-list-cycle-ii/)          |       [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0142.Linked-List-CycleII)        | Medium     | O(n) | O(1)  | Linked List, Two Pointers   |
| [0344](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0344.Reverse-String/)                   |               [Reverse String](https://leetcode.com/problems/reverse-string/)                |          [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0344.Reverse-String)          | Easy       | O(n) | O(1)  | Two Pointers                |
| [0876](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0876.Middle-of-the-Linked-List/)        |    [Middle of the Linked List](https://leetcode.com/problems/middle-of-the-linked-list/)     |    [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0876.Middle-of-the-Linked-List)     | Easy       |      |       | Linked List, Two Pointers   |
| [0011](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0011.Container-With-Most-Water/)        |                 [Container With Most Water](0011.Container-With-Most-Water)                  |    [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0011.Container-With-Most-Water)     | Medium     |      |       | Two Pointers                |
| [0074](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0074.Search-a-2D-Matrix/)               |           [Search a 2D Matrix](https://leetcode.com/problems/search-a-2d-matrix/)            |        [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0074.Search-a-2D-Matrix)        | Medium     |      |       | Binary Search, Two Pointers |

---

#### Bit Manipulation∂

| No.                                                                                                  |                                                   Title                                                   |                                                    Solution                                                    | Difficulty | Time       | Space       | Topic            |
|------------------------------------------------------------------------------------------------------|:---------------------------------------------------------------------------------------------------------:|:--------------------------------------------------------------------------------------------------------------:|------------|------------|-------------|------------------|
| [0693](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0693.Binary-Number-with-Alternating-Bits/) | [Binary Number with Alternating Bits](https://leetcode.com/problems/binary-number-with-alternating-bits/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0693.Binary-Number-with-Alternating-Bits) | Easy       | O(n)/ O(1) | O(1) / O(1) | Bit Manipulation |

---

#### Union Find

| No.                                                                             |                              Title                              |                                         Solution                                          | Difficulty | Time              | Space       | Topic      |
|---------------------------------------------------------------------------------|:---------------------------------------------------------------:|:-----------------------------------------------------------------------------------------:|------------|-------------------|-------------|------------|
| [0721](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0721.Accounts-Merge/) | [Accounts Merge](https://leetcode.com/problems/accounts-merge/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0721.Accounts-Merge) | Easy       | O(n) / O(n log n) | O(n) / O(n) | Union Find |

---

#### Breadth First Search
* DFS 算法可以被認為是[回溯算法](#backtracking-回溯法), BFS算法都是用**Queue**這種數據結構, 每次將一個截短周圍的所有節點加入Queue.
* BFS 找到的路徑一定是最短的, 但是代價是**空間複雜度**比DFS大. [BFS vs DFS](https://github.com/kimi0230/MyGoNote/blob/main/structures/BFS_vs_DFS.md)
* 優化: 雙向 BFS 優化, 在 while 開始時做一個判斷. 讓每次都選擇較小的集合進行擴散,
那麼佔用的空間增長速度就會慢一些, 盡可能以最小的空間代價產生 curDepth 和 nextDepth 的交集
無論單向的 BFS 或是 雙向BFS, 優化過的BFS 空間複雜度都是一樣的

```c++
// 計算從起點 start 到 終點 target 的最點距離
int BFS(Node start, Node targe){
    Queue<Node> q; // 核心數據結構
    Set<Node> visited; // 避免走回頭路
    
    q.offer(start); // 將起點加入 Queue
    visited.add(start);
    int step = 0; // 紀錄擴散的步數

    while(q not empty) {
        int sz = q.size();
        // 當前 Queue 中的所有節點向四周擴散
        for (int i = 0 ; i < sz; i++) {
            Node cur = q.poll();
            // 這裡判斷是否到達終點
            if (cur is target) {
                return step;
            }

            // 將cur 的相鄰節點加入 Queue
            for (Node x : cur.adj()) {
                if (x not in visited) {
                    q.offer(x);
                    visited.add(x);
                }
            }
        }
        // 在這裡更新步數
        step++
    }
}
```

| No.                                                                                   |                                    Title                                    |                                            Solution                                             | Difficulty | Time | Space | Topic                |
|---------------------------------------------------------------------------------------|:---------------------------------------------------------------------------:|:-----------------------------------------------------------------------------------------------:|------------|------|-------|----------------------|
| [0310](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0310.Minimum-Height-Trees/) | [Minimum Height Trees](https://leetcode.com/problems/minimum-height-trees/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0310.Minimum-Height-Trees) | Medium     |      |       | Breadth First Search |
| [0752](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0752.Open-the-Lock/)        |     [752. Open the Lock](https://leetcode.com/problems/open-the-lock/)      |    [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0752.Open-the-Lock)     | Medium     |      |       | Breadth First Search |

---
#### Binary Search
分析二分搜尋技巧: 不要出現 **else**, 而是把所有情況用 **else if** 寫清楚.
計算 mid 時需要防止溢出
```c++
int binarySearch(int[] nums, int target){
    int left = 0 , right = ...;
    while(...) {
        int mid = left + (right - left)/2
        if (nums[mid] == target){
            ...
        } else if (nums[mid] < target){
            left = ...
        } else if (nums[mid] > target){
            right = ...
        }
    }
    return ...;
}
```
| No.                                                                                  |                                           Title                                            |                                            Solution                                            | Difficulty | Time                                  | Space                         | Topic         |
|--------------------------------------------------------------------------------------|:------------------------------------------------------------------------------------------:|:----------------------------------------------------------------------------------------------:|------------|---------------------------------------|-------------------------------|---------------|
| [0704](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0704.Binary-Search/)       |             [704. Binary Search](https://leetcode.com/problems/binary-search/)             |    [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0704.Binary-Search)    | Easy       | 最差:O(long n)<br> 最佳O(1)剛好在中間 | 迭代: O(1) <br/> 遞迴O(log n) | Binary Search |
| [0875](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0875.Koko-Eating-Bananas/) | [875. Koko Eating Bananas](https://leetcode.com/problems/koko-eating-bananas/description/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0875.Koko-Eating-Bananas) | Medium     | O（n log m）                          | O(1)                          | Binary Search |


---
#### Minimax

---
### Graph

#### Graph

---

#### Topological Sort

---

#### Tree
| No.                                                                                                                       |                                                                               Title                                                                               |                                                               Solution                                                               | Difficulty | Time | Space           | Topic     |
|---------------------------------------------------------------------------------------------------------------------------|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------:|:------------------------------------------------------------------------------------------------------------------------------------:|------------|------|-----------------|-----------|
| [0226](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0226.Invert-Binary-Tree)                                        |                                              [Invert Binary Tree](https://leetcode.com/problems/invert-binary-tree/)                                              |                    [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0226.Invert-Binary-Tree)                     | Easy       | O(n) | O(1)            | Tree      |
| [0104](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0104.Maximum-Depth-of-Binary-Tree/)                             |                             [Maximum Depth of Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree//description/)                              |               [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0104.Maximum-Depth-of-Binary-Tree)                | Easy       | O(n) | O(1)            | Tree      |
| [0543](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0543.Diameter-of-Binary-Tree)                                   |                                         [Diameter of Binary Tree](https://leetcode.com/problems/diameter-of-binary-tree/)                                         |                  [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0543.Diameter-of-Binary-Tree)                  | Easy       | O(n) | O(n), O(log(n)) | Tree, DFS |
| [0110](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0110.Balanced-Binary-Tree)                                      |                                            [Balanced Binary Tree](https://leetcode.com/problems/balanced-binary-tree/)                                            |                   [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0110.Balanced-Binary-Tree)                    | Easy       | O(n) | O(1)            | Tree, DFS |
| [0100](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0100.Same-Tree)                                                 |                                                       [Same Tree](https://leetcode.com/problems/same-tree/)                                                       |                         [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0100.Same-Tree)                         | Easy       | O(n) | O(1)            | Tree      |
| [0105](https://kimi0230.github.io/LeetcodeGolang/Leetcode/0105.Construct-Binary-Tree-from-Preorder-and-Inorder-Traversal) | [Construct Binary Tree from Preorder and Inorder Traversal](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0105.Construct-Binary-Tree-from-Preorder-and-Inorder-Traversal) | Medium     | O(n) | O(n)            | Array     |

---

#### Tree Traversal

---

#### Binary Search Tree

---
#### Compputational Geometry

---
### Selected Topics
#### Mathematics
---
#### Random
---
#### Bitwise Manipulation
---
## [GeeksforGeeks](https://www.geeksforgeeks.org/) Content
|  Topic  |                      Title                       | No.  |                                                                      Solution                                                                      | Difficulty |   TimeComplexity   | SpaceComplexity |
|:-------:|:------------------------------------------------:|:-----|:--------------------------------------------------------------------------------------------------------------------------------------------------:|:----------:|:------------------:|:---------------:|
| Sorting | Find Minimum Difference Between Any Two Elements | 0031 | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/GeeksforGeeks/SortingAlgorithms/0031.Find-Minimum-Difference-Between-Any-Two-Elements) |   Basic    | O(n^2), O(n log n) |   O(n), O(n)    |

---
## [Codility](https://app.codility.com/programmers/) Content
<table cellspacing="1" cellpadding="1" frame="solid"  align='border_left'>
    <tr>   
        <th colspan="2">  Topic </th>
        <th> Title</th>
        <th> Solution</th>
        <th> Difficulty</th>
        <th> TimeComplexity</th>
        <th> SpaceComplexity</th>
    </tr>  
    <!-- lesson 1  -->
    <tr>   
        <th> Lesson 1</th>
        <th>
            <a href="https://app.codility.com/programmers/lessons/1-iterations/"> Iterations <br/></a>
        </th>
        <th> 
            <a href="https://app.codility.com/programmers/lessons/1-iterations/binary_gap/"> Binary Gap <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0001.Iterations/Binary-Gap"> Go <br/></a>
        </th>
        <th> Painless </th>
        <th> O(log n) </th>
        <th> O(1) </th>
    </tr> 
    <!-- lesson 2  -->
    <tr>   
        <th rowspan="2"> Lesson 2</th>
        <th rowspan="2">
            <a href="https://app.codility.com/programmers/lessons/2-array/"> Array <br/></a>
        </th>
        <th> 
            <a href="https://app.codility.com/programmers/lessons/2-arrays/cyclic_rotation/"> Cyclic Rotation <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0002.Array/CyclicRotation"> Go <br/></a>
        </th>
        <th> Painless </th>
        <th> O(1) </th>
        <th> O(1) </th>
    </tr> 
    <tr>   
        <th> 
            <a href="https://app.codility.com/programmers/lessons/2-arrays/odd_occurrences_in_array/"> Odd Occurrences In Array <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0002.Array/OddOccurrencesInArray"> Go <br/></a>
        </th>
        <th> Painless </th>
        <th> O(n), O(n) </th>
        <th> O(n), O(1) </th>
    </tr>
    <!-- lesson 3  -->
     <tr>   
        <th rowspan="3"> Lesson 3</th>
        <th rowspan="3">
            <a href="https://app.codility.com/programmers/lessons/3-time_complexity/"> Time Complexity <br/></a>
        </th>
        <th> 
            <a href="https://app.codility.com/programmers/lessons/3-time_complexity/frog_jmp/"> Frog Jmp <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0003.Time-Complexity/FrogJmp"> Go <br/></a>
        </th>
        <th> Painless </th>
        <th> O(1) </th>
        <th> O(1) </th>
    </tr> 
    <tr>   
        <th> 
            <a href="https://app.codility.com/programmers/lessons/3-time_complexity/perm_missing_elem/"> Perm Missing Elem <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0003.Time-Complexity/PermMissingElem"> Go <br/></a>
        </th>
        <th> Painless </th>
        <th> O(n) </th>
        <th> O(1) </th>
    </tr>
    <tr>   
        <th> 
            <a href="https://app.codility.com/programmers/lessons/3-time_complexity/tape_equilibrium/"> Tape Equilibrium <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0003.Time-Complexity/TapeEquilibrium"> Go <br/></a>
        </th>
        <th> Painless </th>
        <th> O(n) </th>
        <th> O(n) </th>
    </tr>
    <!-- lesson 4  -->
     <tr>   
        <th rowspan="4"> Lesson 4</th>
        <th rowspan="4">
            <a href="https://app.codility.com/programmers/lessons/4-counting_elements/"> Counting Elements <br/></a>
        </th>
        <th> 
            <a href="https://app.codility.com/programmers/lessons/4-counting_elements/frog_river_one/"> Frog River One <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0004.Counting-Elements/FrogRiverOne"> Go <br/></a>
        </th>
        <th> Painless </th>
        <th> O(n) </th>
        <th> O(n) </th>
    </tr> 
    <tr>   
        <th> 
            <a href="https://app.codility.com/programmers/lessons/4-counting_elements/max_counters/"> Max Counters <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0004.Counting-Elements/MaxCounters"> Go <br/></a>
        </th>
        <th style="color:#5EC0C6;"> Respectable </th>
        <th> O(n+m) </th>
        <th> O(n) </th>
    </tr>
    <tr>   
        <th> 
            <a href="https://app.codility.com/programmers/lessons/4-counting_elements/missing_integer/"> Missing Integer <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0004.Counting-Elements/MissingInteger"> Go <br/></a>
        </th>
        <th style="color:#5EC0C6;"> Respectable </th>
        <th> O(n) </th>
        <th> O(n) </th>
    </tr>
     <tr>   
        <th> 
            <a href="https://app.codility.com/programmers/lessons/4-counting_elements/perm_check/"> Perm Check <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0004.Counting-Elements/PermCheck"> Go <br/></a>
        </th>
        <th> Painless </th>
        <th> O(n) </th>
        <th> O(n) </th>
    </tr>
    <!-- lesson 5  -->
    <tr>   
        <th rowspan="4"> Lesson 5</th>
        <th rowspan="4">
            <a href="https://app.codility.com/programmers/lessons/5-prefix_sums/"> Prefix Sums <br/></a>
        </th>
        <th> 
            <a href="https://app.codility.com/programmers/lessons/5-prefix_sums/count_div/"> Count Div <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0005.Prefix-Sums/CountDive"> Go <br/></a>
        </th>
        <th style="color:#5EC0C6;"> Respectable </th>
        <th> O(1) </th>
        <th> O(1) </th>
    </tr> 
    <tr>   
        <th> 
            <a href="https://app.codility.com/programmers/lessons/5-prefix_sums/genomic_range_query/"> Genomic Range Query <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0005.Prefix-Sums/GenomicRangeQuery"> Go <br/></a>
        </th>
        <th style="color:#5EC0C6;"> Respectable </th>
        <th> O(n+m) </th>
        <th> O(n) </th>
    </tr>
    <tr>   
        <th> 
            <a href="https://app.codility.com/programmers/lessons/5-prefix_sums/min_avg_two_slice/"> MinAvg Two Slice <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0005.Prefix-Sums/MinAvgTwoSlice"> Go <br/></a>
        </th>
        <th style="color:#5EC0C6;"> Respectable </th>
        <th> O(n) </th>
        <th> O(n) </th>
    </tr>
    <tr>   
        <th> 
            <a href="https://app.codility.com/programmers/lessons/5-prefix_sums/passing_cars/"> Passing Cars <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0005.Prefix-Sums/PassingCars"> Go <br/></a>
        </th>
        <th > Painless </th>
        <th> O(n) </th>
        <th> O(1) </th>
    </tr>
    <!-- lesson 6  -->
    <tr>   
        <th rowspan="4"> Lesson 6</th>
        <th rowspan="4">
            <a href="https://app.codility.com/programmers/lessons/6-sorting/"> Sorting <br/></a>
        </th>
        <th> 
            <a href="https://app.codility.com/programmers/lessons/6-sorting/distinct/"> Distinct <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0006.Sorting/Distinct"> Go <br/></a>
        </th>
        <th> Painless </th>
        <th> O(nlogn) </th>
        <th> O(n) </th>
    </tr> 
    <tr>   
        <th> 
            <a href="https://app.codility.com/programmers/lessons/6-sorting/max_product_of_three/"> Max Product of Three <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0006.Sorting/MaxProductOfThree"> Go <br/></a>
        </th>
        <th> Painless </th>
        <th> O(nlogn) </th>
        <th> O(1) </th>
    </tr>
    <tr>   
        <th> 
            <a href="https://app.codility.com/programmers/lessons/6-sorting/number_of_disc_intersections/"> Number Of Disc Intersections <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0006.Sorting/NumberOfDiscIntersections"> Go <br/></a>
        </th>
        <th style="color:#5EC0C6;"> Respectable </th>
        <th> O(nlogn) </th>
        <th> O(n) </th>
    </tr>
     <tr>   
        <th> 
            <a href="https://app.codility.com/programmers/lessons/6-sorting/triangle/"> Triangle <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0006.Sorting/Triangle"> Go <br/></a>
        </th>
        <th > Painless </th>
        <th> O(nlogn) </th>
        <th> O(n) </th>
    </tr>
    <!-- lesson 7  -->
    <tr>   
        <th rowspan="4"> Lesson 7</th>
        <th rowspan="4">
            <a href="https://app.codility.com/programmers/lessons/7-stacks_and_queues/"> Stacks and Queues <br/></a>
        </th>
        <th> 
            <a href="https://app.codility.com/programmers/lessons/7-stacks_and_queues/brackets/"> Brackets <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0007.Stacks-and-Queues/Brackets"> Go <br/></a>
        </th>
        <th> Painless </th>
        <th> O(n) </th>
        <th> O(n) </th>
    </tr> 
    <tr>   
        <th> 
            <a href="https://app.codility.com/programmers/lessons/7-stacks_and_queues/fish/"> Fish <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0007.Stacks-and-Queues/Fish"> Go <br/></a>
        </th>
        <th> Painless </th>
        <th> O(n) </th>
        <th> O(n) </th>
    </tr>
    <tr>   
        <th> 
            <a href="https://app.codility.com/programmers/lessons/7-stacks_and_queues/nesting/"> Nesting <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0007.Stacks-and-Queues/Nesting"> Go <br/></a>
        </th>
        <th> Painless </th>
        <th> O(n) </th>
        <th> O(1) </th>
    </tr>
     <tr>   
        <th> 
            <a href="https://app.codility.com/programmers/lessons/7-stacks_and_queues/stone_wall/"> Stone Wall <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0007.Stacks-and-Queues/StoneWall"> Go <br/></a>
        </th>
        <th > Painless </th>
        <th> O(n) </th>
        <th> O(n) </th>
    </tr>
    <!-- lesson 8  -->
    <tr>   
        <th rowspan="2"> Lesson 8</th>
        <th rowspan="2">
            <a href="https://app.codility.com/programmers/lessons/8-leader/"> Leader <br/></a>
        </th>
        <th> 
            <a href="https://app.codility.com/programmers/lessons/8-leader/dominator/"> Dominator <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0008.Leader/Dominator"> Go <br/></a>
        </th>
        <th> Painless </th>
        <th> O(n) </th>
        <th> O(1) </th>
    </tr> 
    <tr>   
        <th> 
            <a href="https://app.codility.com/programmers/lessons/8-leader/equi_leader/"> EquiLeader <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0008.Leader/EquiLeader"> Go <br/></a>
        </th>
        <th> Painless </th>
        <th> O(n) </th>
        <th> O(n) </th>
    </tr>
    <!-- lesson 9  -->
    <tr>   
        <th rowspan="3"> Lesson 9</th>
        <th rowspan="3">
            <a href="https://app.codility.com/programmers/lessons/9-maximum_slice_problem/"> Maximum slice problem <br/></a>
        </th>
        <th> 
            <a href="https://app.codility.com/programmers/lessons/9-maximum_slice_problem/max_profit/"> Max Profit <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0009.Maximum-Slice-Problem/MaxProfit"> Go <br/></a>
        </th>
        <th> Painless </th>
        <th> O(n) </th>
        <th> O(1) </th>
    </tr> 
    <tr>   
        <th> 
            <a href="https://app.codility.com/programmers/lessons/9-maximum_slice_problem/max_slice_sum/"> Max Slice Sum <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0009.Maximum-Slice-Problem/MaxSliceSum"> Go <br/></a>
        </th>
        <th> Painless </th>
        <th> O(n) </th>
        <th> O(n) </th>
    </tr>
    <tr>   
        <th> 
            <a href="https://app.codility.com/programmers/lessons/9-maximum_slice_problem/max_double_slice_sum/"> Max Double Slice Sum <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0009.Maximum-Slice-Problem/MaxDoubleSliceSum"> Go <br/></a>
        </th>
        <th style="color:#5EC0C6;"> Respectable </th>
        <th> O(n) </th>
        <th> O(n) </th>
    </tr>
    <!-- lesson 10  -->
    <tr>   
        <th rowspan="4"> Lesson 10</th>
        <th rowspan="4">
            <a href="https://app.codility.com/programmers/lessons/10-prime_and_composite_numbers/"> Prime and composite numbers <br/></a>
        </th>
        <th> 
            <a href="https://app.codility.com/programmers/lessons/10-prime_and_composite_numbers/count_factors/"> Count Factors <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0010.Prime-And-Composite-Numbers/CountFactors"> Go <br/></a>
        </th>
        <th> Painless </th>
        <th> O(sqrt(n)) </th>
        <th> O(1) </th>
    </tr> 
    <tr>   
        <th> 
            <a href="https://app.codility.com/programmers/lessons/10-prime_and_composite_numbers/flags/"> Flags <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0010.Prime-And-Composite-Numbers/Flags"> Go <br/></a>
        </th>
        <th style="color:#5EC0C6;"> Respectable </th>
        <th> O(n) </th>
        <th> O(n) </th>
    </tr>
    <tr>   
        <th> 
            <a href="https://app.codility.com/programmers/lessons/10-prime_and_composite_numbers/min_perimeter_rectangle/"> MinPerimeterRectangle <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0010.Prime-And-Composite-Numbers/MinPerimeterRectangle"> Go <br/></a>
        </th>
        <th> Painless </th>
        <th> O(sqrt(n))) </th>
        <th> O(1) </th>
    </tr>
    <tr>   
        <th> 
            <a href="https://app.codility.com/programmers/lessons/10-prime_and_composite_numbers/peaks/"> Peaks <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0010.Prime-And-Composite-Numbers/Peaks"> Go <br/></a>
        </th>
        <th style="color:#5EC0C6;"> Respectable </th>
        <th> O( n*log( log(n) )) </th>
        <th> O(n) </th>
    </tr>
    <!-- lesson 11  -->
    <tr>   
        <th rowspan="2"> Lesson 11</th>
        <th rowspan="2">
            <a href="https://app.codility.com/programmers/lessons/11-sieve_of_eratosthenes/"> Sieve of Eratosthenes <br/> (質數篩) </a>
        </th>
        <th> 
            <a href="https://app.codility.com/programmers/lessons/11-sieve_of_eratosthenes/count_non_divisible/"> Count Non Divisible <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0011.Sieve-of-Eratosthenes/CountNonDivisible"> Go <br/></a>
        </th>
        <th style="color:#5EC0C6;"> Respectable </th>
        <th> O(N * log(N)) </th>
        <th> O(n) </th>
    </tr> 
    <tr>   
        <th> 
            <a href="https://app.codility.com/programmers/lessons/11-sieve_of_eratosthenes/count_semiprimes/"> Count Semiprimes <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0011.Sieve-of-Eratosthenes/CountSemiprimes"> Go <br/></a>
        </th>
        <th style="color:#5EC0C6;"> Respectable </th>
        <th> O(N*log(log(N))+M) </th>
        <th> O(N+M) </th>
    </tr>
    <!-- lesson 12  -->
    <tr>   
        <th rowspan="2"> Lesson 12</th>
        <th rowspan="2">
            <a href="https://app.codility.com/programmers/lessons/12-euclidean_algorithm/"> Euclidean algorithm <br/> (輾轉相除法 or 歐幾里得算法) </a>
        </th>
        <th> 
            <a href="https://app.codility.com/programmers/lessons/12-euclidean_algorithm/chocolates_by_numbers/"> Chocolates By Numbers <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0012.Euclidean-Algorithm/ChocolatesByNumbers"> Go <br/></a>
        </th>
        <th> Painless </th>
        <th> O(log(N + M)) </th>
        <th> O(1) </th>
    </tr> 
    <tr>   
        <th> 
            <a href="https://app.codility.com/programmers/lessons/12-euclidean_algorithm/common_prime_divisors/"> Common Prime Divisors <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0012.Euclidean-Algorithm/CommonPrimeDivisors"> Go <br/></a>
        </th>
        <th style="color:#5EC0C6;"> Respectable </th>
        <th> O(Z * log(max(A) + max(B))**2) </th>
        <th> O(1) </th>
    </tr>
    <!-- lesson 13  -->
    <tr>   
        <th rowspan="2"> Lesson 13</th>
        <th rowspan="2">
            <a href="https://app.codility.com/programmers/lessons/13-fibonacci_numbers/"> Fibonacci numbers <br/> </a>
        </th>
        <th> 
            <a href="https://app.codility.com/programmers/lessons/13-fibonacci_numbers/fib_frog/"> FibFrog <br/></a>
        </th>
        <th>
            <a href="https://github.com/kimi0230/LeetcodeGolang/tree/master/Codility/Lesson/0013.Fibonacci-Numbers/FibFrog"> Go <br/></a>
        </th>
        <th style="color:#5EC0C6;"> Respectable </th>
        <th> O(N * log(N)) </th>
        <th> O(N) </th>
    </tr> 
    <tr>   
        <th> 
            <a href="https://app.codility.com/programmers/lessons/13-fibonacci_numbers/ladder/"> Ladder <br/></a>
        </th>
        <th>
            <a href="">  <br/></a>
        </th>
        <th style="color:#5EC0C6;"> Respectable </th>
        <th>  </th>
        <th>  </th>
    </tr>
    <!-- lesson 14  -->
    <tr>   
        <th rowspan="2"> Lesson 14</th>
        <th rowspan="2">
            <a href="https://app.codility.com/programmers/lessons/14-binary_search_algorithm/"> Binary search algorithm <br/> </a>
        </th>
        <th> 
            <a href="https://app.codility.com/programmers/lessons/14-binary_search_algorithm/min_max_division/"> MinMaxDivision <br/></a>
        </th>
        <th>
            <a href="">  <br/></a>
        </th>
        <th style="color:#5EC0C6;"> Respectable </th>
        <th>  </th>
        <th>  </th>
    </tr> 
    <tr>   
        <th> 
            <a href="https://app.codility.com/programmers/lessons/14-binary_search_algorithm/nailing_planks/"> NailingPlanks <br/></a>
        </th>
        <th>
            <a href="">  <br/></a>
        </th>
        <th style="color:#5EC0C6;"> Respectable </th>
        <th>  </th>
        <th>  </th>
    </tr>
    <!-- lesson 15  -->
    <tr>   
        <th rowspan="4"> Lesson 15</th>
        <th rowspan="4">
            <a href="https://app.codility.com/programmers/lessons/15-caterpillar_method/"> Caterpillar method <br/> </a>
        </th>
        <th> 
            <a href="https://app.codility.com/programmers/lessons/15-caterpillar_method/abs_distinct/"> AbsDistinct <br/></a>
        </th>
        <th>
            <a href="">  <br/></a>
        </th>
        <th> Painless </th>
        <th>  </th>
        <th>  </th>
    </tr> 
    <tr>   
        <th> 
            <a href="https://app.codility.com/programmers/lessons/15-caterpillar_method/count_distinct_slices/"> CountDistinctSlices <br/></a>
        </th>
        <th>
            <a href="">  <br/></a>
        </th>
        <th> Painless </th>
        <th>  </th>
        <th>  </th>
    </tr>
    <tr>   
        <th> 
            <a href="https://app.codility.com/programmers/lessons/15-caterpillar_method/count_triangles/"> CountTriangles <br/></a>
        </th>
        <th>
            <a href="">  <br/></a>
        </th>
        <th> Painless </th>
        <th>  </th>
        <th>  </th>
    </tr>
    <tr>   
        <th> 
            <a href="https://app.codility.com/programmers/lessons/15-caterpillar_method/min_abs_sum_of_two/"> MinAbsSumOfTwo <br/></a>
        </th>
        <th>
            <a href="">  <br/></a>
        </th>
        <th style="color:#5EC0C6;"> Respectable </th>
        <th>  </th>
        <th>  </th>
    </tr>
    <!-- lesson 16  -->
    <tr>   
        <th rowspan="2"> Lesson 16</th>
        <th rowspan="2">
            <a href="https://app.codility.com/programmers/lessons/16-greedy_algorithms/"> Greedy algorithms <br/> </a>
        </th>
        <th> 
            <a href="https://app.codility.com/programmers/lessons/16-greedy_algorithms/max_nonoverlapping_segments/"> MaxNonoverlappingSegments <br/></a>
        </th>
        <th>
            <a href="">  <br/></a>
        </th>
        <th> Painless </th>
        <th>  </th>
        <th>  </th>
    </tr> 
    <tr>   
        <th> 
            <a href="https://app.codility.com/programmers/lessons/16-greedy_algorithms/tie_ropes/"> TieRopes <br/></a>
        </th>
        <th>
            <a href="">  <br/></a>
        </th>
        <th> Painless </th>
        <th>  </th>
        <th>  </th>
    </tr>
    <!-- lesson 17  -->
    <tr>   
        <th rowspan="2"> Lesson 17</th>
        <th rowspan="2">
            <a href="https://app.codility.com/programmers/lessons/17-dynamic_programming/"> Dynamic programming <br/> </a>
        </th>
        <th> 
            <a href="https://app.codility.com/programmers/lessons/17-dynamic_programming/min_abs_sum/"> MinAbsSum <br/></a>
        </th>
        <th>
            <a href="">  <br/></a>
        </th>
        <th style="color:#4BA0A5;" > Ambitious </th>
        <th>  </th>
        <th>  </th>
    </tr> 
    <tr>   
        <th> 
            <a href="https://app.codility.com/programmers/lessons/17-dynamic_programming/number_solitaire/"> NumberSolitaire <br/></a>
        </th>
        <th>
            <a href="">  <br/></a>
        </th>
        <th style="color:#5EC0C6" > Respectable </th>
        <th>  </th>
        <th>  </th>
    </tr>
</table>

## Reference
* [leetcode](https://leetcode.com/)
* [leetcode-cn](https://leetcode-cn.com/)
* [halfrost](https://books.halfrost.com/leetcode/)
* [anakin](https://github.com/anakin/golang-leetcode)
* [wufenggirl](https://github.com/wufenggirl/LeetCode-in-Golang)
* [GeeksforGeeks](https://www.geeksforgeeks.org/)
* [Codility](https://app.codility.com/programmers/)
* [GitHub: labuladong/fucking-algorithm(labuladong 算法小抄)](https://github.com/labuladong/fucking-algorithm)
* https://cses.fi/
* https://github.com/neetcode-gh/leetcode