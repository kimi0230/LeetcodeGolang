package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"
)

type LeetcodeResponse struct {
	Data struct {
		Question struct {
			QuestionID string `json:"questionId"`
			Title      string `json:"title"`
			Content    string `json:"content"`
			Difficulty string `json:"difficulty"`
			Tags       []struct {
				Name string `json:"name"`
			} `json:"topicTags"`
		} `json:"question"`
	} `json:"data"`
}

type ProblemListResponse struct {
	StatStatusPairs []struct {
		Stat struct {
			QuestionID   int    `json:"question_id"`
			QuestionSlug string `json:"question__title_slug"`
		} `json:"stat"`
	} `json:"stat_status_pairs"`
}

const readmeTemplate = `---
title: {{.ID}}.{{.Title}}
subtitle: "https://leetcode.com/problems/{{.Slug}}/description/"
date: {{.Now}}
lastmod: {{.Now}}
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "{{.Title}}"
license: ""
images: []

tags: [LeetCode, Go, {{.Difficulty}}{{range .Tags}}, {{.}}{{end}}]
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
mapbox:
share:
  enable: true
comment:
  enable: true
library:
  css:
  js:
seo:
  images: []
---

# [{{.ID}}.{{.Title}}](https://leetcode.com/problems/{{.Slug}}/description/)

## 題目
{{.Problem}}

## 題目大意


## 解題思路


## 來源
* https://leetcode.com/problems/{{.Slug}}/description/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/{{.ID}}.{{.Title}}/main.go



## Benchmark


`

type ReadmeData struct {
	ID         string
	Title      string
	Slug       string
	Difficulty string
	Tags       []string
	Problem    string
	Now        string
}

func fetchSlugFromID(id int) (string, error) {
	cacheFile := ".leetcode_problems_cache.json"
	var problems ProblemListResponse

	if content, err := os.ReadFile(cacheFile); err == nil {
		_ = json.Unmarshal(content, &problems)
	} else {
		resp, err := http.Get("https://leetcode.com/api/problems/all/")
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		if err := json.Unmarshal(body, &problems); err != nil {
			return "", err
		}
		_ = os.WriteFile(cacheFile, body, 0644)
	}

	for _, p := range problems.StatStatusPairs {
		if p.Stat.QuestionID == id {
			return p.Stat.QuestionSlug, nil
		}
	}
	return "", errors.New("slug not found for id")
}

func fetchLeetcode(slug string) (*ReadmeData, error) {
	query := `query questionTitleSlug($titleSlug: String!) {
  question(titleSlug: $titleSlug) {
    questionId
    title
    content
    difficulty
    topicTags {
      name
    }
  }
}`

	payload := map[string]interface{}{
		"query":     query,
		"variables": map[string]string{"titleSlug": slug},
	}
	jsonPayload, _ := json.Marshal(payload)

	resp, err := http.Post("https://leetcode.com/graphql", "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var result LeetcodeResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	q := result.Data.Question
	tags := []string{}
	for _, tag := range q.Tags {
		tags = append(tags, tag.Name)
	}

	return &ReadmeData{
		ID:         q.QuestionID,
		Title:      strings.ReplaceAll(q.Title, " ", "-"),
		Slug:       slug,
		Difficulty: q.Difficulty,
		Tags:       tags,
		Problem:    stripHTML(q.Content),
		Now:        time.Now().Format(time.RFC3339),
	}, nil
}

func stripHTML(input string) string {
	var output strings.Builder
	inTag := false
	for _, r := range input {
		switch r {
		case '<':
			inTag = true
		case '>':
			inTag = false
		default:
			if !inTag {
				output.WriteRune(r)
			}
		}
	}
	return strings.TrimSpace(output.String())
}

func writeREADME(data *ReadmeData) error {
	dir := fmt.Sprintf("Leetcode/%s.%s", data.ID, data.Title)
	os.MkdirAll(dir, 0755)
	f, err := os.Create(filepath.Join(dir, "README.md"))
	if err != nil {
		return err
	}
	defer f.Close()
	tmpl, err := template.New("readme").Parse(readmeTemplate)
	if err != nil {
		return err
	}
	return tmpl.Execute(f, data)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: leetcode-readme-gen <id>")
		return
	}

	idNum, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Please provide a valid Leetcode problem number (e.g. 217)")
		return
	}

	slug, err := fetchSlugFromID(idNum)
	if err != nil {
		panic(err)
	}

	data, err := fetchLeetcode(slug)
	if err != nil {
		panic(err)
	}
	if err := writeREADME(data); err != nil {
		panic(err)
	}
	fmt.Printf("README generated for %s (%s)\n", data.Title, data.ID)
}
