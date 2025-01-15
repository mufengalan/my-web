package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Http struct {
	Token string
}

func NewHttp(token string) *Http {
	return &Http{
		Token: token,
	}
}

func (h *Http) GetRequest(url string) ([]byte, error) {
	// 发送 GET 请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %w", err)
	}
	//设置header
	req.Header.Set("utoken", h.Token)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %w", err)
	}
	defer response.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return body, fmt.Errorf("error reading response: %w", err)
	}

	return body, nil
}
