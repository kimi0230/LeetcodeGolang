# [59. Spiral Matrix II](https://leetcode.com/problems/spiral-matrix-ii/)


## 題目

Given a positive integer *n*, generate a square matrix filled with elements from 1 to *n*2 in spiral order.

**Example:**


    Input: 3
    Output:
    [
     [ 1, 2, 3 ],
     [ 8, 9, 4 ],
     [ 7, 6, 5 ]
    ]


## 題目大意

給定一個正整數 n，生成一個包含 1 到 n^2 所有元素，且元素按順時針順序螺旋排列的正方形矩陣。


## 解題思路

- 給出一個數組 n，要求輸出一個 n * n 的二維數組，裡面元素是 1 - n*n，且數組排列順序是螺旋排列的
- 這一題是第 54 題的加強版，沒有需要注意的特殊情況，直接模擬即可。

## 來源
* https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0059.Spiral-Matrix-II/
* https://leetcode-cn.com/problems/spiral-matrix-ii/
* [數組：這個循環可以轉懵很多人！](https://mp.weixin.qq.com/s/KTPhaeqxbMK9CxHUUgFDmg)