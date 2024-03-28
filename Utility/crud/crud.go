package crud

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

var (
	errStatusNo200 = errors.New("http code not 200")
)

type Crud struct {
	URL string
	Request
}

type Request struct {
	Header map[string]string
	Params map[string]string
	Body   map[string]interface{}
}

type Response struct {
	HttpCode int
	Response string
	Header   http.Header
	Error    error
}

func NewCrud(url string) *Crud {
	return &Crud{URL: url}
}

func handleRequest(url string, method string, header map[string]string, params map[string]string, body map[string]interface{}) (int, string, http.Header, error) {
	// body
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	jsonStr, _ := json.Marshal(body)

	// ignore SSL
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return 0, "", nil, err
	}

	// headers
	for key, value := range header {
		req.Header.Add(key, value)
	}

	// params
	q := req.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}

	client := &http.Client{}

	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, "", nil, err
	}
	respBody := string(content)
	headers := resp.Header

	return resp.StatusCode, respBody, headers, nil
}

func (crud *Crud) Post() Response {
	httpCode, response, headers, error := handleRequest(crud.URL, "POST", crud.Header, crud.Params, crud.Body)

	resp := Response{
		HttpCode: httpCode,
		Response: response,
		Header:   headers,
		Error:    error,
	}

	return resp
}

func (crud *Crud) Get() Response {
	httpCode, response, headers, error := handleRequest(crud.URL, "GET", crud.Header, crud.Params, crud.Body)

	resp := Response{
		HttpCode: httpCode,
		Response: response,
		Header:   headers,
		Error:    error,
	}

	return resp
}

func (crud *Crud) Put() Response {
	httpCode, response, headers, error := handleRequest(crud.URL, "PUT", crud.Header, crud.Params, crud.Body)

	resp := Response{
		HttpCode: httpCode,
		Response: response,
		Header:   headers,
		Error:    error,
	}

	return resp
}

func (crud *Crud) Delete() Response {
	httpCode, response, headers, error := handleRequest(crud.URL, "DELETE", crud.Header, crud.Params, crud.Body)

	resp := Response{
		HttpCode: httpCode,
		Response: response,
		Header:   headers,
		Error:    error,
	}

	return resp
}
