---
title: "A1B2C3: Two Go Routine Print A1B2C3....Z26"
subtitle: "LEETCODELINK"
date: 2022-07-25T13:11:00+08:00
lastmod: 2022-07-25T13:11:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "Two Go Routine Print A1B2C3....Z26"
license: ""
images: []

tags: [Algorithm, A1B2C3]
categories: [Algorithm]

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

# A1B2C3

用兩個 go routine 印出 A1B2C3....Z26
Two Go Routine Print A1B2C3....Z26

## Channle With Buffer
```go
func ChannelWBuffer() {
	abc := make(chan struct{}, 1)
	num := make(chan struct{}, 1)
	done := make(chan struct{})
	abc <- struct{}{}

	go func() {
		for i := 65; i <= 90; i++ {
			<-abc
			// fmt.Printf("%v", string(rune(i)))
			num <- struct{}{}
		}
	}()
	go func() {
		for i := 1; i <= 26; i++ {
			<-num
			// fmt.Printf("%v", i)
			abc <- struct{}{}
		}
		done <- struct{}{}
	}()
	// time.Sleep(1 * time.Second)
	<-done
}
```

## Channel Without Buffer
```go
func ChannelWOBuffer() {
	abc := make(chan struct{})
	num := make(chan struct{})
	done := make(chan struct{})

	go func() {
		for i := 65; i <= 90; i++ {
			// fmt.Printf("%v", string(rune(i)))
			abc <- struct{}{}
			<-num
		}
	}()
	go func() {
		for i := 1; i <= 26; i++ {
			<-abc
			// fmt.Printf("%v", i)
			num <- struct{}{}
		}
		done <- struct{}{}
	}()
	// time.Sleep(1 * time.Second)
	<-done
}
```

## Wait Group
```go

func WGLock() {
	abcMux := sync.Mutex{}
	numMux := sync.Mutex{}

	numMux.Lock()
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 65; i <= 90; i++ {
			abcMux.Lock()
			// fmt.Printf("%v", string(rune(i)))
			numMux.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		for i := 1; i <= 26; i++ {
			numMux.Lock()
			// fmt.Printf("%v", i)
			abcMux.Unlock()
		}
	}()
	wg.Wait()
}
```

## Benchmark
```sh
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Algorithms/A1B2C3
cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
BenchmarkChannelWBuffer-4          79159             14612 ns/op             344 B/op          5 allocs/op
BenchmarkChannelWOBuffer-4         83068             14451 ns/op             344 B/op          5 allocs/op
BenchmarkWGLock-4                  51303             23072 ns/op              96 B/op          5 allocs/op
PASS
ok      LeetcodeGolang/Algorithms/A1B2C3        4.092s
```