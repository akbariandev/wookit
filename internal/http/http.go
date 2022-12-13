package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const woocommerceEndPoint = "wp-json/wc/v3"

type Request struct {
	httpClient  *http.Client
	httpRequest *http.Request
	url         string
	method      string
	body        any
}

func (r *Request) Call() (*http.Response, error) {
	res, err := r.httpClient.Do(r.httpRequest)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *Request) makeUrl(host, endPoint string, urlParams any) error {

	var mapUrlParams map[string]string
	var arrayUrlParams []string
	byteUrlParam, err := json.Marshal(urlParams)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(byteUrlParam, &mapUrlParams); err != nil {
		return err
	}

	r.url = fmt.Sprintf("%s/%s/%s", host, woocommerceEndPoint, endPoint)
	for key, param := range mapUrlParams {
		arrayUrlParams = append(arrayUrlParams, fmt.Sprintf("%s=%s", key, param))
	}

	joinedParams := strings.Join(arrayUrlParams, "&")
	if len(joinedParams) > 0 {
		r.url = fmt.Sprintf("%s?%s", r.url, joinedParams)
	}

	return nil
}

func NewGetRequest(host, endPoint string, urlParams, body any, headers ...map[string][]string) (*Request, error) {
	return newRequest(http.MethodPost, host, endPoint, urlParams, body, headers...)
}

func NewPostRequest(host, endPoint string, urlParams, body any, headers ...map[string][]string) (*Request, error) {
	return newRequest(http.MethodPost, host, endPoint, urlParams, body, headers...)
}

func newRequest(method, host, endPoint string, urlParams, body any, headers ...map[string][]string) (*Request, error) {
	r := &Request{
		httpClient: &http.Client{},
		method:     method,
	}

	if err := r.makeUrl(host, endPoint, urlParams); err != nil {
		return nil, err
	}

	var bodyBuff bytes.Buffer
	if err := json.NewEncoder(&bodyBuff).Encode(body); err != nil {
		return nil, err
	}

	httpRequest, err := http.NewRequest(r.method, r.url, &bodyBuff)
	if err != nil {
		return nil, err
	}

	r.httpRequest = httpRequest
	r.httpRequest.Header = http.Header{
		"Content-Type":   []string{"application/json"},
		"Content-Length": []string{strconv.FormatInt(r.httpRequest.ContentLength, 10)},
	}

	for _, h := range headers {
		for k, v := range h {
			r.httpRequest.Header.Set(k, v[0])
		}
	}

	return r, nil
}
