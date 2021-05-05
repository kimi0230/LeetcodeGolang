# [721. Accounts Merge](https://leetcode.com/problems/accounts-merge/)


## 题目

Given a list `accounts`, each element `accounts[i]` is a list of strings, where the first element `accounts[i][0]` is a name, and the rest of the elements are emailsrepresenting emails of the account.

Now, we would like to merge these accounts. Two accounts definitely belong to the same person if there is some email that is common to both accounts. Note that even if two accounts have the same name, they may belong to different people as people could have the same name. A person can have any number of accounts initially, but all of their accounts definitely have the same name.

After merging the accounts, return the accounts in the following format: the first element of each account is the name, and the rest of the elements are emails **in sorted order**. The accounts themselves can be returned in any order.

**Example 1:**

    Input: 
    accounts = [["John", "johnsmith@mail.com", "john00@mail.com"], ["John", "johnnybravo@mail.com"], ["John", "johnsmith@mail.com", "john_newyork@mail.com"], ["Mary", "mary@mail.com"]]
    Output: [["John", 'john00@mail.com', 'john_newyork@mail.com', 'johnsmith@mail.com'],  ["John", "johnnybravo@mail.com"], ["Mary", "mary@mail.com"]]
    Explanation: 
    The first and third John's are the same person as they have the common email "johnsmith@mail.com".
    The second John and Mary are different people as none of their email addresses are used by other accounts.
    We could return these lists in any order, for example the answer [['Mary', 'mary@mail.com'], ['John', 'johnnybravo@mail.com'], 
    ['John', 'john00@mail.com', 'john_newyork@mail.com', 'johnsmith@mail.com']] would still be accepted.

**Note:**

- The length of `accounts` will be in the range `[1, 1000]`.
- The length of `accounts[i]` will be in the range `[1, 10]`.
- The length of `accounts[i][j]` will be in the range `[1, 30]`.

## 題目大意

給定一個列表 `accounts`，每個元素 `accounts[i]` 是一個字符串列表，其中第一個元素 `accounts[i][0]` 是 名稱 (name)，其餘元素是 emails 表示該賬戶的郵箱地址。 現在，我們想合併這些賬戶。如果兩個賬戶都有一些共同的郵箱地址，則兩個賬戶必定屬於同一個人。請注意，即使兩個賬戶具有相同的名稱，它們也可能屬於不同的人，因為人們可能具有相同的名稱。一個人最初可以擁有任意數量的賬戶，但其所有賬戶都具有相同的名稱。 合併賬戶後，按以下格式返回賬戶：每個賬戶的第一個元素是名稱，其餘元素是按字符 ASCII 順序排列的郵箱地址。賬戶本身可以以任意順序返回。

## 解題思路
- 給出一堆賬戶和對應的郵箱。要求合併同一個人的多個郵箱賬戶。
- 如果這個人名和所屬的其中之一的郵箱是相同的，就判定這是同一個人的郵箱，那麼就合併這些郵箱。
- 這題的解題思路是並查集。不過如果用暴力合併的方法，時間複雜度非常差。優化方法是先把每組數據都進行編號，人編號，每個郵箱都進行編號。這個映射關係用 `map` 記錄起來。如果利用並查集的 `union()` 操作，把這些編號都進行合併。最後把人的編號和對應郵箱的編號拼接起來。
- 這一題有 2 處比較“坑”的是，不需要合併的用戶的郵箱列表也是需要排序和去重的，同一個人的所有郵箱集合都要合併到一起。具體見測試用例。不過題目中也提到了這些點，也不能算題目坑，只能歸自己沒注意這些邊界情況。

## 來源
https://books.halfrost.com/leetcode/ChapterFour/0700~0799/0721.Accounts-Merge/


## 時間複雜度：`O(n log n)`
其中 n 是不同郵箱地址的數量。
需要遍歷所有郵箱地址，在並查集內進行查找和合併操作，對於兩個不同的郵箱地址，
如果它們的祖先不同則需要進行合併，需要進行 2 次查找和最多 1 次合併。一共需要進行 2n 次查找和最多 n 次合併，
因此時間複雜度是 O(2n log n) = O(n log n)。
這裡的並查集使用了路徑壓縮，但是沒有使用按秩合併，
最壞情況下的時間複雜度是O(n log n)，平均情況下的時間複雜度依然是 O(nα(n))，其中α 為阿克曼函數的反函數，α(n) 可以認為是一個很小的常數。
整理出題目要求的返回賬戶的格式時需要對郵箱地址排序，時間複雜度是 O(n log n)。
其餘操作包括遍歷所有郵箱地址，在哈希表中記錄相應的信息，時間複雜度是O(n)，
在漸進意義下O(n) 小於O(nlog n)。
因此總時間複雜度是 O(n log n)。

作者：LeetCode-Solution
鏈接：https://leetcode-cn.com/problems/accounts-merge/solution/zhang-hu-he-bing-by-leetcode-solution-3dyq/
來源：力扣（LeetCode）
著作權歸作者所有。商業轉載請聯繫作者獲得授權，非商業轉載請註明出處。

## 來源
* https://books.halfrost.com/leetcode/ChapterFour/0700~0799/0721.Accounts-Merge/
* https://leetcode-cn.com/problems/accounts-merge/