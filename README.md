## Leetcode in Golang
[Leetcode](https://leetcode.com/), [Codility](https://app.codility.com/programmers/) , [GeekforGeeks algorithms](https://www.geeksforgeeks.org/fundamentals-of-algorithms/?ref=shm) exercises written in Golang.

---
## [leetcode](https://leetcode.com/) Content
- [Leetcode in Golang](#leetcode-in-golang)
- [leetcode Content](#leetcode-content)
  - [Data Structure](#data-structure)
    - [Array](#array)
    - [Stack](#stack)
    - [Linked List](#linked-list)
  - [Algorithm](#algorithm)
    - [Sort](#sort)
    - [Backtracking (回溯法)](#backtracking-回溯法)
    - [Dynamic Programming](#dynamic-programming)
    - [Sliding Window](#sliding-window)
    - [Two Pointers](#two-pointers)
    - [Bit Manipulation](#bit-manipulation)
    - [Union Find](#union-find)
    - [Breadth First Search](#breadth-first-search)
- [GeeksforGeeks Content](#geeksforgeeks-content)
- [Codility Content](#codility-content)
- [Reference](#reference)
### Data Structure
#### Array
| No. | Title | Solution | Difficulty | Time | Space | Topic |
|--|:--:|:--:|--|--|--|--|
| 0001 | [Two Sum](https://leetcode.com/problems/two-sum/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0001.Two-Sum) | Easy | O(n) | O(n) | Array |
| 0003 | [Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0003.Longest-Substring-Without-Repeating-Characters) | Medium | O(n) | O(1) | Array |
| 0015 | [3 Sum](https://leetcode.com/problems/3sum/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0015.3Sum) | Medium | O(n^2) | O(n) | Array |
| 0027 | [Remove Element](https://leetcode.com/problems/remove-element/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0027.Remove-Element) | Easy | O(n) | O(1) | Array |
| 0035 | [Search Insert Position](https://leetcode.com/problems/search-insert-position/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0035.Search-Insert-Position) | Easy |  O(n), O(logn) | O(1) | Array |
| 0059 | [Spiral Matrix II](https://leetcode.com/problems/spiral-matrix-ii/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0059.Spiral-Matrix-II) | Medium | O(n) | O(n^2) | Array |
| 0088 | [Merge Sorted Array](https://leetcode.com/problems/merge-sorted-array/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0088.Merge-Sorted-Array) | Easy | O(n) | O(1) | Array |

---

#### Stack

| No. | Title | Solution | Difficulty | Time | Space | Topic |
|-----|:-----:|:--------:|------------|------|-------|-------|
| 0094 | [Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0094.Binary-Tree-Inorder-Traversal) | Medium | O(n) | O(1) | Stack |

---

#### Linked List

| No. | Title | Solution | Difficulty | Time | Space | Topic |
|-----|:-----:|:--------:|------------|------|-------|-------|
| 0141 | [Linked List Cycle](https://leetcode.com/problems/linked-list-cycle) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0141.Linked-List-Cycle) | Easy | O(n) | O(1) | Linked List, Two Pointers |
| 0142 | [Linked List Cycle II](https://leetcode.com/problems/linked-list-cycle-ii/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0142.Linked-List-CycleII) | Medium |  |  | Linked List, Two Pointers |
| 0876 | [Middle of the Linked List](https://leetcode.com/problems/middle-of-the-linked-list/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0876.Middle-of-the-Linked-List) | Easy |  |  | Linked List, Two Pointers |
| 0203 | [Remove Linked List Elements](https://leetcode.com/problems/remove-linked-list-elements/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0203.Remove-Linked-List-Elements) | Easy | O(n) | O(1) | Linked List |

---

### Algorithm
#### Sort
| No. | Title | Solution | Difficulty | Time | Space | Topic |
|-----|:-----:|:--------:|------------|------|-------|-------|
| 0075 | [Sort Colors](https://leetcode.com/problems/sort-colors/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0075.Sort-Color) | Medium | O(n) | O(1) | Sort |


---

#### Backtracking (回溯法)
DFS. 解決一個回溯問題, 實際上就是一個決策數的遍歷過程.
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

| No. | Title | Solution | Difficulty | Time | Space | Topic |
|-----|:-----:|:--------:|------------|------|-------|-------|
| 0046 | [Permutations (全排列)](https://leetcode.com/problems/permutations/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0046.Permutations) | Medium | O(n) | O(n) | Backtracking |

---

#### Dynamic Programming
動態規劃問題的一般形式就是求**最值**, **最長**遞增子序列, **最小**編輯距離等. 核心問題是窮舉
1. 重疊子問題
   1. memory table
   2. DP table
2. 最優子結構
3. 狀態轉移方程式
   1. 這問題的 base case (最簡單情況) 是什麼?
   2. 這問題有什麼**狀態**
   3. 對於每個狀態, 可以做出什麼**選擇**, 使得狀態發生改變
   4. 如何定義 dp 數組/函數的含義來表現**狀態**和**選擇**?

```python
# 初始化 base case
dp[0][0][...] = base
# 進行狀態轉移
for 狀態1 in 狀態1的所有取值：
    for 狀態2 in 狀態2的所有取值：
        for ...
            dp[狀態1][狀態2][...] = 求最值(選擇1，選擇2...)
```

| No. | Title | Solution | Difficulty | Time | Space | Topic |
|-----|:-----:|:--------:|------------|------|-------|-------|
| 0053 | [Maximum Subarray](https://leetcode.com/problems/maximum-subarray/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0053.Maximum-Subarray) | Easy | O(n) | O(n) | Dynamic Programming |
| 0322 | [Coin Change](https://leetcode.com/problems/coin-change/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0322.Coin-Change) | Medium | O(nm) | O(n) | Dynamic Programming |
| 0509 | [Fibonacci Number](https://leetcode.com/problems/fibonacci-number/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0509.Fibonacci-Number) | Easy | 很多解法 | 很多解法 | Dynamic Programming |

---

#### Sliding Window

| No. | Title | Solution | Difficulty | Time | Space | Topic |
|-----|:-----:|:--------:|------------|------|-------|-------|
| 0209 | [Minimum Size Subarray Sum](https://leetcode.com/problems/minimum-size-subarray-sum/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0209.Minimum-Size-Subarray-Sum) | Medium | O(n^2) / O(n) / O(nlog n) |  O(1) / O(1) / O(n) | Sliding Window |

---

#### Two Pointers
分為兩類 1. "快,慢指針" 2. "左,右指針"
1. 快,慢指針: 主要解決 linkedlist 問題, 典型的判斷 linkedlist 是否包含環
2. 左,右指針: 主要解決array(或 string)中的問題, 如二分搜尋.
https://labuladong.gitee.io/algo/2/20/56/

| No. | Title | Solution | Difficulty | Time | Space | Topic |
|-----|:-----:|:--------:|------------|------|-------|-------|
| 0141 | [Linked List Cycle](https://leetcode.com/problems/linked-list-cycle) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0141.Linked-List-Cycle) | Easy | O(n) | O(1) | Linked List, Two Pointers |
| 0142 | [Linked List Cycle II](https://leetcode.com/problems/linked-list-cycle-ii/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0142.Linked-List-CycleII) | Medium |  |  | Linked List, Two Pointers |
| 0344 | [Reverse String](https://leetcode.com/problems/reverse-string/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0344.Reverse-String) | Easy | O(n) |  O(1) | Two Pointers |
| 0876 | [Middle of the Linked List](https://leetcode.com/problems/middle-of-the-linked-list/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0876.Middle-of-the-Linked-List) | Easy |  |  | Linked List, Two Pointers |

---

#### Bit Manipulation

| No. | Title | Solution | Difficulty | Time | Space | Topic |
|-----|:-----:|:--------:|------------|------|-------|-------|
| 0693 | [Binary Number with Alternating Bits](https://leetcode.com/problems/binary-number-with-alternating-bits/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0693.Binary-Number-with-Alternating-Bits) | Easy | O(n)/ O(1) |  O(1) / O(1) | Bit Manipulation |

---

#### Union Find

| No. | Title | Solution | Difficulty | Time | Space | Topic |
|-----|:-----:|:--------:|------------|------|-------|-------|
| 0721 | [Accounts Merge](https://leetcode.com/problems/accounts-merge/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0721.Accounts-Merge) | Easy | O(n) / O(n log n) |  O(n) / O(n) | Union Find |

---

#### Breadth First Search
DFS 算法可以被認為是[回溯算法](#backtracking-回溯法), BFS算法都是用**Queue**這種數據結構, 每次將一個截短周圍的所有節點加入Queue.
BFS 找到的路徑一定是最短的, 但是代價是**空間複雜度**比DFS大
[BFS vs DFS](https://github.com/kimi0230/MyGoNote/blob/main/structures/BFS_vs_DFS.md)

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

| No. | Title | Solution | Difficulty | Time | Space | Topic |
|-----|:-----:|:--------:|------------|------|-------|-------|
| 0310 | [Minimum Height Trees](https://leetcode.com/problems/minimum-height-trees/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0310.Minimum-Height-Trees) | Medium |  |  | Breadth First Search |
| 0752 | [752. Open the Lock](https://leetcode.com/problems/open-the-lock/) | [Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/Leetcode/0752.Open-the-Lock) | Medium |  |  | Breadth First Search |

---

## [GeeksforGeeks](https://www.geeksforgeeks.org/) Content
| Topic    |  Title  | No. |  Solution  |  Difficulty |  TimeComplexity	| SpaceComplexity|
|:--------:|:--------:|:--------------------------------------------------------------|:--------:|:--------:|:--------:|:--------:|
|Sorting|Find Minimum Difference Between Any Two Elements|0031|[Go](https://github.com/kimi0230/LeetcodeGolang/tree/master/GeeksforGeeks/SortingAlgorithms/0031.Find-Minimum-Difference-Between-Any-Two-Elements)| Basic |O(n^2), O(n log n)|O(n), O(n)|

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
* [labuladong 算法小抄](https://github.com/labuladong/fucking-algorithm)