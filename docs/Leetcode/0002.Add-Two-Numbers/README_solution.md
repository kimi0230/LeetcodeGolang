---
title: UMPIRE 0002.Add-Two-Numbers
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
  # ...
mapbox:
  # ...
share:
  enable: true
  # ...
comment:
  enable: true
  # ...
library:
  css:
    # someCSS = "some.css"
    # located in "assets/"
    # Or
    # someCSS = "https://cdn.example.com/some.css"
  js:
    # someJS = "some.js"
    # located in "assets/"
    # Or
    # someJS = "https://cdn.example.com/some.js"
seo:
  images: []
  # ...
---

# UMPIRE 0002.Add-Two-Numbers

## Output 1: UMPIRE 解題（完整思考版）

### U – Understand（理解題目）
- **題目描述**：給定兩個非空的鏈結串列，表示兩個非負整數。數字以逆序儲存，每個節點包含一個位數。將這兩個數字相加，並以鏈結串列形式返回結果。
- **關鍵限制**：
    - 兩個數字都不會以 0 開頭，除非是數字 0 本身。
    - 數字是逆序儲存的（這實際上簡化了從個位數開始加的過程）。
- **Happy Path**：
    - `l1 = [2, 4, 3], l2 = [5, 6, 4]` -> 返回 `[7, 0, 8]` ($342 + 465 = 807$)。
- **Edge Cases**：
    - 串列長度不同：`l1 = [0], l2 = [7, 3]` -> 返回 `[7, 3]` ($0 + 37 = 37$)。
    - 最後一位產生進位：`l1 = [9, 9], l2 = [1]` -> 返回 `[0, 0, 1]` ($99 + 1 = 100$)。

### M – Match（匹配知識）
- **主要模式**：**鏈結串列遍歷 (Linked List Traversal) + 模擬數學加法**。
- **為什麼適合**：
    - 由於數字已經是逆序（個位在頭），我們可以直接從頭節點開始逐位相加，並處理進位 (Carry)。這正是手寫加法的標準過程。
- **其他方案**：
    - **轉換為整數再相加**：將鏈結串列轉為 BigInt，相加後轉回。雖然直觀，但若數字極長會超過語言原生整數類型的範圍（LeetCode 鏈結串列可能非常長），且這種方法違背了考察鏈結串列操作的初衷。

### P – Plan（制定計畫）
1. 初始化一個虛擬頭節點 `dummy` 以方便操作，並建立一個指針 `current` 指向當前節點。
2. 初始化 `carry` (進位) 為 0。
3. 開始遍歷 `l1` 或 `l2`，只要其中一個不為 `nil` 或 `carry` 不為 0 就繼續：
    - 取得 `l1` 的值（若為 `nil` 則為 0），取得 `l2` 的值（若為 `nil` 則為 0）。
    - 計算 `sum = val1 + val2 + carry`。
    - 更新 `carry = sum / 10`。
    - 建立新節點，值為 `sum % 10`，接到 `current.Next`，並移動 `current`。
    - 若 `l1`, `l2` 還有節點，則往後移動。
4. 返回 `dummy.Next`。
- **避免的 Bug**：特別注意當遍歷完 `l1` 和 `l2` 後，如果 `carry` 還大於 0，一定要補上最後一個節點（例如 $9+1$ 的情況）。

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
    dummy := &ListNode{}
    curr := dummy
    carry := 0
    
    // 只要還有數字或還有進位，就繼續計算
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
    
    return dummy.Next
}
```

### R – Review（檢查與回顧）
- **Dry Run**：`l1 = [9, 9], l2 = [1]`
    1. `val1=9, val2=1, carry=0` -> `sum=10, carry=1, node=0`。
    2. `val1=9, val2=0, carry=1` -> `sum=10, carry=1, node=0`。
    3. `val1=0, val2=0, carry=1` (Loop 條件成立，因為 carry=1) -> `sum=1, carry=0, node=1`。
    4. 結果為 `[0, 0, 1]`。
- **關鍵狀態**：`carry` 是連結不同位數計算的唯一橋樑，必須在每一輪精確更新。

### E – Evaluate（總結與評估）
- **時間複雜度**：$O(\max(m, n))$。其中 $m, n$ 分別是兩個鏈結串列的長度，我們只需掃描一次。
- **空間複雜度**：$O(1)$。若不計算存儲結果所需的節點空間。若算入，則為 $O(\max(m, n) + 1)$。
- **權衡**：使用虛擬頭節點 `dummy` 大大簡化了串列初始化的判斷邏輯。

---

## Output 2: 面試官口語回答腳本（精簡可直接說）

### 1️⃣ 開場：題目理解
這題要求我們模擬兩個大數相加。題目給的是兩個鏈結串列，每個節點代表一位數，並且是逆序排列的。我們需要回傳一個同樣格式的相加結果串列。

### 2️⃣ 解法選擇說明
我會使用**直接模擬加法運算**的方式，從鏈結串列的頭部（也就是個位數）開始遍歷。這是一個非常典型的雙指針（或多指針）遍歷問題，配合進位處理。

### 3️⃣ 解題策略概覽
我會同步移動兩個串列的指針。在每一步，我把兩個節點的值加上前一步留下的進位 `carry`。相加後的個位數就是新節點的值，而十位數則成為下一輪的進位。有一個細節要注意，如果兩個串列都走完了，但最後還剩下進位，一定要補上最後一個節點。

### 4️⃣ 寫程式時會補充的關鍵說明
在寫程式時，我會使用一個 **dummy node (虛擬頭節點)**，這樣我就不用在第一輪特別去判斷回傳串列是否為空，程式碼會簡潔很多。另外，我會用一個 `while` 迴圈同時處理串列剩餘與進位剩餘，這樣逻辑最完整。

### 5️⃣ 快速 Dry Run 說明
以 $99 + 1$ 為例：
第一位 $9+1=10$，寫下 $0$ 進位 $1$；
第二位 $9+0$ 加上進位 $1$ 又得到 $10$，寫下 $0$ 進位 $1$；
最後兩邊都沒數字了，但進位還有 $1$，所以我再補一個節點寫下 $1$。
最後回傳的鏈結串列就是 `[0, 0, 1]`，代表 $100$。

### 6️⃣ 收尾總結
這個解法的**時間複雜度是 $O(\max(m, n))$**，我們只遍歷了最長的那個串列一次。這在效率上已經是最優解了。
