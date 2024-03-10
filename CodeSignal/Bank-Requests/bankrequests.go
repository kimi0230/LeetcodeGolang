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

// 錯誤處理: 回傳錯誤的 "負號" + request id
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
		bank.requestId = index
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
