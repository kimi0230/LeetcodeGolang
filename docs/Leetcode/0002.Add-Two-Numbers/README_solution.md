---
title: UMPIRE 2.Add-Two-Numbers
subtitle: "https://leetcode.com/problems/add-two-numbers/description/"
date: 2024-03-02T15:57:00+08:00
lastmod: 2024-03-02T15:57:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "UMPIRE Solution for 0002.Add-Two-Numbers"
license: ""
images: []

tags:
  - LeetCode
  - Go
  - Medium
  - Add Two Numbers
  - Linked List
  - Math
  - Recursion
  - Amazon
  - Apple
  - Facebook
  - Microsoft
  - Bloomberg
categories: [LeetCode]

featuredImage: ""
featuredImagePreview: ""

hiddenFromHomePage: false
hiddenFromSearch: false
twemoji: false
lightgallery: true
ruby: true
fraction: true
fontawesome: true
linkToMarkdown: false
rssFullText: false

toc:
  enable: true
  auto: true
code:
  copy: true
  maxShownLines: 200
math:
  enable: false
share:
  enable: true
comment:
  enable: true
---

# Solution - UMPIRE

## Output 1: UMPIRE 解題（完整思考版）

### U – Understand（理解題目）
- **題目描述**：給予兩個非空的鏈結串列，分別代表兩個非負整數。數字是以「逆序」儲存的，意即鏈結串列的第一個節點代表個位數。請將這兩個數字相加，並以鏈結串列的形式回傳總和。
- **關鍵限制**：
    - 每個節點只包含一個數字（0-9）。
    - 除非是數字 0 本身，否則這兩個數字都不會有前導零。
    - 兩個鏈結串列的長度可能不同。
- **澄清與細節**：
    - 題目保證輸入串列非空，且內容皆為非負整數。
    - 需要精確處理「進位（Carry）」邏輯，尤其是最後一位產生的進位。
- **測試案例定義**：
    - **Happy Path**: 
        - $L1: [2, 4, 3]$ (342), $L2: [5, 6, 4]$ (465) $\rightarrow$ 結果: $[7, 0, 8]$ (807)
    - **Edge Cases**:
        - **長度不同**: $L1: [9, 9, 9]$, $L2: [1] \rightarrow$ 結果: $[0, 0, 0, 1]$ (處理不同步結束與連續進位)。
        - **最後進位**: $L1: [5]$, $L2: [5] \rightarrow$ 結果: $[0, 1]$ (最後迴圈結束後仍有進位)。
        - **值為零**: $L1: [0]$, $L2: [0] \rightarrow$ 結果: $[0]$。

### M – Match（匹配知識）
- **主要演算法/資料結構**：鏈結串列遍歷（Linked List Traversal）與基礎算術（Elementary Math）。
- **為什麼適合**：
    - 數字逆序儲存的特性與我們手寫加法「從右往左」的順序完全一致，只需單向遍歷即可。
- **為什麼其他解法不理想**：
    - 轉換成整數再相加會面臨整數溢位（Integer Overflow）的問題，且轉換過程會產生額外的時間與空間開銷。直接在鏈結串列上逐位計算是最符合題意的最優解。

### P – Plan（制定計畫）
- **解題流程**：
    1. 使用 `dummyHead` 簡化首個節點的處理，並用 `curr` 指標追蹤結果串列的尾部。
    2. 初始化進位變數 `carry = 0`。
    3. 遍歷條件：只要 `l1` 沒走完 OR `l2` 沒走完 OR 還有進位 `carry != 0`：
        - 從 `l1` 取值（若為 nil 則補 0）。
        - 從 `l2` 取值（若為 nil 則補 0）。
        - `sum = val1 + val2 + carry`。
        - `carry = sum / 10`。
        - 建立值為 `sum % 10` 的新節點，接在 `curr.Next`。
        - 移動 `curr` 以及 `l1` / `l2` 指標。
    4. 回傳 `dummyHead.Next`。
- **資料流與狀態變化**：
    - `carry` 狀態跨越每一次疊代傳遞。
    - 透過 `l1` 與 `l2` 的同步移動確保位數對齊。
- **預防 Bug**：
    - 確保 `while` 迴圈結束後，如果 `carry` 仍大於 0，必須補上最後一個節點。
    - 在存取 `l1.Val` 之前務必檢查 `l1 != nil`。

### I – Implement（實際實作，Golang）
```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{Val: 0}
	curr := dummyHead
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		val1, val2 := 0, 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		sum := val1 + val2 + carry
		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
	}

	return dummyHead.Next
}
```

### R – Review（檢查與回顧）
- **Dry Run**:
    - 輸入：$L1: [2, 4], L2: [5, 6, 4]$
    - 1st: $2+5+0 = 7, carry=0, Result: [7]$
    - 2nd: $4+6+0 = 10, carry=1, Result: [7, 0]$
    - 3rd: $0+4+1 = 5, carry=0, Result: [7, 0, 5]$
    - 回傳 $[7, 0, 5]$，正確（$42 + 465 = 507$）。
- **關鍵狀態轉移**：進位正確地在疊代間傳遞。
- **邊界確認**：單節點 $5+5$ 會產生兩個節點的結果 $[0, 1]$，邏輯正確。

### E – Evaluate（總結與評估）
- **時間複雜度**: $O(\max(M, N))$，其中 $M, N$ 為兩串列長度。我們只遍歷了最長串列一次。
- **空間複雜度**: $O(\max(M, N))$，用於儲存結果串列（若不計入回傳值則為 $O(1)$）。
- **權衡與優化**：本解法已達最優。若要極致優化空間，可考慮原位修改（In-place），但會破壞輸入資料且程式碼複雜度會大幅提升，通常不建議這樣做。

---

## Output 2: 面試官口語回答腳本（精簡可直接說）

### 1️⃣ 開場：題目理解
這題要求我們處理兩個逆序排列的鏈結串列，將它們代表的整數相加。主要的挑戰在於處理進位以及當兩個串列長度不一致時的對齊問題。

### 2️⃣ 解法選擇說明
我選擇採用**線性遍歷（Simultaneous Traversal）**配合**基礎進位邏輯**。因為串列已經是逆序排列（低位在首），這正好符合我們手算加法的順序，直接一邊走一邊加是最有效率的，時間複雜度是 $O(\max(M, N))$。

### 3️⃣ 解題策略概覽
我會建立一個 **Dummy Head** 來方便管理結果串列的成長。在迴圈中，我會同時檢查兩個指標，只要其中一個還有數字或者還有進位，我就把值加總起來，存入新節點，並把產生的進位傳給下一輪。

### 4️⃣ 寫程式時會補充的關鍵說明
*   **Dummy Head 技巧**：使用虛擬頭節點可以讓程式碼更簡潔，避免在迴圈內一直判斷 `head` 是否為空的邏輯。
*   **靈活處理不等長**：當其中一個串列結束時，我會將其值視為 $0$，這樣可以用同一套邏輯處理完所有的位數。
*   **強健的結束條件**：我的迴圈條件會包含 `carry != 0`，這能自動處理像是 $5+5$ 這種最後多出一位進位的情況。

### 5️⃣ 快速 Dry Run 說明
假設 $L1=[2, 4]$，$L2=[5, 6, 4]$。
第一位相加 $2+5=7$，無進位。
第二位相加 $4+6=10$，結果寫 $0$，進位 $1$。
最後 $L1$ 走完了，剩 $L2$ 的 $4$ 加上剛才的進位 $1$ 等於 $5$。
最終結果就是 $[7, 0, 5]$。

### 6️⃣ 收尾總結
這套演算法在**時間與空間複雜度上都是線性 $O(\max(M, N))$**。它能有效處理任意長度的數字，且不受程式語言整數型態大小的限制。
