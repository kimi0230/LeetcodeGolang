---
title: Bank Requests
subtitle: ""
date: 2024-03-10T14:31:00+08:00
lastmod: 2024-03-10T14:31:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "Bank Requests"
license: ""
images: []

tags: [CodeSignal, Go, Bank Requests]
categories: [CodeSignal]

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
# Bank Requests

_A solution to one of my coding interview questions. Complete solution - written in **GoLang**_

---
### Task:

You've been asked to program a bot for a popular bank that will automate the management of incoming requests. There are three types of requests the bank can receive:
 
* transfer i j sum: request to transfer sum amount of money from the `i-th` account to the `j-th` one
* deposit `i` sum: request to deposit sum amount of money in the `i-th` account
* withdraw `i` sum: request to withdraw sum amount of money from the `i-th` account.

Your bot should also be able to process invalid requests. There are two types of invalid requests: invalid account number in the requests; deposit / withdrawal of a larger amount of money than is currently available.

For the given list of accounts and requests, return the state of accounts after all requests have been processed, or an array of a single element `[- <request_id>]` (please note the minus sign), where ***<request_id>*** is the **1-based** index of the first invalid request.

Example for `accounts = [10, 100, 20, 50, 30]` and `requests = ["withdraw 2 10", "transfer 5 1 20", "deposit 5 20", "transfer 3 4 15"]`, the output should be `bankRequests(accounts, requests) = [30, 90, 5, 65, 30]`.

`accounts = [10, 100, 20, 50, 30]` : 代表 ID:1 有10元, ID:2 有100元 ...
"withdraw 2 10" : 從ID:2 提領 10元, 所以100-10 最後得到accounts= [10, 90, 20, 50, 30]

Here are the states of accounts after each request:
* `"withdraw 2 10": [10, 90, 20, 50, 30]`
* `"transfer 5 1 20": [30, 90, 20, 50, 10]`
* `"deposit 5 20": [30, 90, 20, 50, 30]`
* `"transfer 3 4 15": [30, 90, 5, 65, 30],` which is the answer

For `accounts = [20, 1000, 500, 40, 90]` and `requests = ["deposit 3 400", "transfer 1 2 30", "withdraw 4 50"]`, the output should be `bankRequests(accounts, requests) = [-2]`.

After the first request, accounts becomes equal to `[20, 1000, 900, 40, 90]`, but the second one turns it into `[-10, 1030, 900, 40, 90]`, which is invalid. Thus, the second request is invalid, and the answer is `[-2]`. Note that the last request is also invalid, but it shouldn't be included in the answer.

中文意思:
你被要求為一家知名銀行編寫一個機器人，該機器人將自動處理來自客戶的請求。銀行可能會收到三種類型的請求：

* `transfer i j sum`：從第`i`個帳戶轉移`sum`金額到第`j`個帳戶的請求
* `deposit i sum`：向第`i`個帳戶存入`sum`金額的請求
* `withdraw i sum`：從第`i`個帳戶提取`sum`金額的請求。

你的機器人還應該能夠處理無效的請求。無效的請求有兩種類型：請求中的帳戶號碼無效；存款/提款金額大於當前可用金額。

對於給定的帳戶列表和請求，返回處理完所有請求後的帳戶狀態，或單一元素為`[- <request_id>]`的數組（請注意負號），其中 **<request_id>** 是第一個無效請求的**從1開始**的索引。

以`accounts = [10, 100, 20, 50, 30]` 和 `requests = ["withdraw 2 10", "transfer 5 1 20", "deposit 5 20", "transfer 3 4 15"]` 為例，輸出應為 `bankRequests(accounts, requests) = [30, 90, 5, 65, 30]`。

`accounts = [10, 100, 20, 50, 30]` 代表：ID:1 有10元, ID:2 有100元 ...
"withdraw 2 10" 從ID:2 提領 10元，所以100-10 最後得到accounts= [10, 90, 20, 50, 30]

這是每個請求後帳戶的狀態：
* `"withdraw 2 10": [10, 90, 20, 50, 30]`
* `"transfer 5 1 20": [30, 90, 20, 50, 10]`
* `"deposit 5 20": [30, 90, 20, 50, 30]`
* `"transfer 3 4 15": [30, 90, 5, 65, 30]`，這是答案

對於`accounts = [20, 1000, 500, 40, 90]` 和 `requests = ["deposit 3 400", "transfer 1 2 30", "withdraw 4 50"]`，輸出應為 `bankRequests(accounts, requests) = [-2]`。

第一個請求後，帳戶變成 `[20, 1000, 900, 40, 90]`，但第二個請求將其變為 `[-10, 1030, 900, 40, 90]`，這是無效的。因此，第二個請求無效，答案為`[-2]`。請注意最後一個請求也是無效的，但不應包含在答案中。

## 解答

```go
package bankrequests

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var (
	errInvaldTransfer = errors.New("invalid transfer")
	errInvaldWithdraw = errors.New("invalid withdraw")
	errInvaldDeposit  = errors.New("invalid deposit")
	errUnknownAction  = errors.New("unknown action")
)

var requestScheme = [][]string{
	3: {"action", "from", "amount"},
	4: {"action", "from", "to", "amount"},
}

type BankService interface {
	getActionName() string
	transfer() error
	withdraw() error
	deposit() error
}

type Bank struct {
	BankService
	action        string
	requestId     int
	from          int
	to            int
	amount        int
	balances      []int
	requestFailed bool
	failedRequest []int

	mutex sync.RWMutex
}

func (bank *Bank) getActionName() string {
	return bank.action
}

func (bank *Bank) transfer() error {
	bank.mutex.Lock()
	defer bank.mutex.Unlock()

	invalidAmout := bank.balances[bank.from-1] < bank.amount
	invalidReceiver := len(bank.balances) < bank.to
	if invalidAmout || invalidReceiver {
		return errInvaldTransfer
	}
	bank.balances[bank.from-1] -= bank.amount
	bank.balances[bank.to-1] += bank.amount
	return nil
}

func (bank *Bank) withdraw() error {
	bank.mutex.Lock()
	defer bank.mutex.Unlock()

	if bank.balances[bank.from-1] < bank.amount {
		return errInvaldWithdraw
	}
	bank.balances[bank.from-1] -= bank.amount
	return nil
}

func (bank *Bank) deposit() error {
	bank.mutex.Lock()
	defer bank.mutex.Unlock()

	invalidAccount := len(bank.balances) < bank.from
	if invalidAccount {
		return errInvaldDeposit
	}
	bank.balances[bank.from-1] += bank.amount
	return nil
}

func NewBank() *Bank {
	return &Bank{}
}

// 錯誤處理: 回傳錯誤為
// 單一元素為`[- <request_id>]`的數組（請注意負號）
func (bank *Bank) failedAction() {
	bank.mutex.Lock()
	defer bank.mutex.Unlock()
	failed := []int{-bank.requestId}

	bank.balances = []int{}
	bank.requestFailed = true
	bank.failedRequest = failed
}

func (bank *Bank) Action() {
	var requestErr error
	action := bank.getActionName()
	// fmt.Println("action:", action)

	switch action {
	case "transfer":
		requestErr = bank.transfer()
	case "withdraw":
		requestErr = bank.withdraw()
	case "deposit":
		requestErr = bank.deposit()
	default:
		requestErr = errUnknownAction
	}
	if requestErr != nil {
		fmt.Println("requestErr:", requestErr)
		bank.failedAction()
	}
}

func extractRequestParams(request string) map[string]interface{} {
	res := map[string]interface{}{}
	reqSlice := strToSlice(request)
	fmt.Println("reqSlice:", reqSlice)

	// len is 3: {"action", "from", "amount"},
	// len is 4: {"action", "from", "to", "amount"},
	scheme := requestScheme[len(reqSlice)]

	for i, v := range scheme {
		res[v] = reqSlice[i]
	}
	return res
}

func strToSlice(str string) []interface{} {
	var res []interface{}
	erp := strings.Fields(str)

	re, err := regexp.Compile("[0-9]+")
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return nil
	}

	for _, v := range erp {
		if isNum := re.MatchString(v); isNum {
			if n, _ := strconv.Atoi(v); n != 0 {
				res = append(res, n)
			}
		} else {
			res = append(res, v)
		}
	}

	return res
}

func bankRequests(requests []string, balances []int) []int {
	var res []int
	for index, request := range requests {
		reqParams := extractRequestParams(request)
		bank := NewBank()
		bank.requestId = index + 1
		bank.action = reqParams["action"].(string)
		bank.amount = reqParams["amount"].(int)
		bank.from = reqParams["from"].(int)
		if _, ok := reqParams["to"]; ok {
			bank.to = reqParams["to"].(int)
		}

		bank.balances = balances
		bank.Action()

		if bank.requestFailed {
			return bank.failedRequest
		}
		res = bank.balances
	}
	return res
}
```

unit test
```go
package bankrequests

import (
	"reflect"
	"testing"
)

var tests = []struct {
	arg1 []string
	arg2 []int
	want []int
}{
	{
		[]string{
			"transfer 1 4 10",
			"deposit 3 10",
			"withdraw 5 15",
		},
		[]int{20, 30, 10, 90, 60},
		[]int{10, 30, 20, 100, 45},
	},
	{
		[]string{
			"transfer 1 4 40",
			"deposit 3 10",
			"withdraw 5 65",
		},
		[]int{20, 30, 10, 90, 60},
		[]int{-1},
	},
	{
		[]string{
			"withdraw 2 10",
			"transfer 5 1 20",
			"deposit 5 20",
			"transfer 3 4 15",
		},
		[]int{10, 100, 20, 50, 30},
		[]int{30, 90, 5, 65, 30},
	},
	{
		[]string{
			"deposit 3 400",
			"transfer 1 2 30",
			"withdraw 4 50",
		},
		[]int{20, 1000, 500, 40, 90},
		[]int{-2},
	},
}

func TestBankRequests(t *testing.T) {
	for _, tt := range tests {
		if got := bankRequests(tt.arg1, tt.arg2); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got = %v, want = %v", got, tt.want)
		}
	}
}
```

## Reference
* https://github.com/crusty0gphr/bank-request-processor/tree/master
