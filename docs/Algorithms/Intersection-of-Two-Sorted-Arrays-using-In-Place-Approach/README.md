---
title: "Intersection of Two Sorted Arrays using In Place Approach"
subtitle: "Intersection of Two Sorted Arrays using In Place Approach"
date: 2023-06-19T14:32:25+08:00
lastmod: 2023-06-19T14:32:25+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: ""
license: ""
images: []

tags:
  - Golang
  - Algorithms
  - Intersection
categories: [Algorithms]

featuredImage: ""
featuredImagePreview: ""

hiddenFromHomePage: false
hiddenFromSearch: false
twemoji: false
lightgallery: true
ruby: true
fraction: true
fontawesome: true
linkToMarkdown: true
rssFullText: false

toc:
  enable: true
  auto: true
code:
  copy: true
  maxShownLines: 50
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
# Intersection of Two Sorted Arrays using In Place Approach

Intersection of Two Sorted Arrays using In Place Approach

要在原地（in-place）解決這個問題，可以使用雙指針的方法。假設給定的兩個數組分別為A和B，它們已經按升序排序。

首先，我們可以初始化兩個指針i和j分別指向A和B的起始位置，然後開始進行比較。如果A[i]小於B[j]，則移動指針i向後移動一位；如果A[i]大於B[j]，則移動指針j向後移動一位；如果A[i]等於B[j]，則將該值添加到結果中，並將兩個指針都向後移動一位。

重複上述步驟，直到其中一個數組的指針達到數組末尾為止。最終，得到的結果就是兩個數組的交集。

```go
package intersection

func FindIntersection(A, B []int) []int {
	var i, j int = 0, 0
	result := []int{}

	for i < len(A) && j < len(B) {
		if A[i] < B[j] {
			i++
		} else if A[i] > B[j] {
			j++
		} else {
			result = append(result, A[i])
			i++
			j++
		}
	}
	return result
}
```

```go
func TestFindIntersection(t *testing.T) {
	var tests = []struct {
		arg1 []int
		arg2 []int
		want []int
	}{
		{
			arg1: []int{1, 3, 4, 6, 7},
			arg2: []int{2, 4, 6, 8, 9},
			want: []int{4, 6},
		},
	}

	for _, tt := range tests {
		if got := FindIntersection(tt.arg1, tt.arg2); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}
```

TODO: 
延伸 
1. [349. Intersection of Two Arrays (easy)](https://leetcode.com/problems/intersection-of-two-arrays/)
2. [350. Intersection of Two Arrays II](https://leetcode.com/problems/intersection-of-two-arrays-ii/)