package crud

import (
	"fmt"
	"testing"
)

func TestPost(t *testing.T) {
	tests := []struct {
		name   string
		url    string
		header map[string]string
		params map[string]string
		body   map[string]interface{}
	}{
		{
			name: "Post",
			url:  "https://10.36.235.47:7071/api/packages",
			header: map[string]string{
				"Version":                 "1.1.0",
				"Windows-Package-Manager": "winget-cli-restsource-dev",
				"Content-Type":            "application/json",
				"Accept":                  "application/Json",
			},
			params: nil,
			body: map[string]interface{}{
				"PackageIdentifier": "Demo.Postman",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			crud := NewCrud(tt.url)
			crud.Header = tt.header
			crud.Params = tt.params
			crud.Body = tt.body
			if got := crud.Post(); got.Error != nil {
				t.Errorf("got = %v", got)
			} else {
				fmt.Println(got.HttpCode)
			}
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		name   string
		url    string
		header map[string]string
		params map[string]string
	}{
		{
			name:   "Get",
			url:    "https://kimi0230.github.io",
			header: nil,
			params: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			crud := NewCrud(tt.url)
			crud.Header = tt.header
			crud.Params = tt.params
			if got := crud.Get(); got.Error != nil {
				t.Errorf("got = %v", got)
			} else {
				fmt.Println(got.HttpCode)
			}
		})
	}
}
