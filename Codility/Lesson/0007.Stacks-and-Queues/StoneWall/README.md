# [StoneWall](https://app.codility.com/programmers/lessons/7-stacks_and_queues/stone_wall/)
Cover "Manhattan skyline" using the minimum number of rectangles.

You are going to build a stone wall. The wall should be straight and N meters long, and its thickness should be constant; however, it should have different heights in different places. The height of the wall is specified by an array H of N positive integers. H[I] is the height of the wall from I to I+1 meters to the right of its left end. In particular, H[0] is the height of the wall's left end and H[N−1] is the height of the wall's right end.

The wall should be built of cuboid (長方體) stone blocks (that is, all sides of such blocks are rectangular). Your task is to compute the minimum number of blocks needed to build the wall.

Write a function:

func Solution(H []int) int

that, given an array H of N positive integers specifying the height of the wall, returns the minimum number of blocks needed to build it.

For example, given array H containing N = 9 integers:

  H[0] = 8    H[1] = 8    H[2] = 5
  H[3] = 7    H[4] = 9    H[5] = 8
  H[6] = 7    H[7] = 4    H[8] = 8
the function should return 7. The figure shows one possible arrangement of seven blocks.

![](https://codility-frontend-prod.s3.amazonaws.com/media/task_static/stone_wall/static/images/auto/4f1cef49cc46d451e88109d449ab7975.png)


Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array H is an integer within the range [1..1,000,000,000].
Copyright 2009–2021 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.


## 題目大意
如何用最少的數量來貼出符合 H 的牆呢
要建立一面長N米的強. 厚度固定, 每個地方的高度不同.
H[I]代表牆從 I 到 I+1 米處的高度.
H[0]大表牆最左到1米處的高度
H[N-1]大表牆N-1米處到最右的高度

## 解題思路
尋遍整個array, 當當前高度大於先前高度時,加入stack裡, 並視為一個矩形. 將結果+1
若當前高度小於先前,將先前高度pop出去.直到stack 為空或當前高度大於等於先前高度

## 來源
* https://app.codility.com/programmers/lessons/7-stacks_and_queues/stone_wall/
* https://github.com/Anfany/Codility-Lessons-By-Python3/blob/master/L7_Stacks%20and%20Queues/7.4%20StoneWall.md