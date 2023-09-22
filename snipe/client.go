package snipe

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/johnmikee/sync2snipe/pkg/helpers"
)

type SnipeClient struct {
	Token string

	BaseUrl string
	Client  *http.Client
}

func NewClient(sc *SnipeClient) *SnipeClient {
	return &SnipeClient{
		BaseUrl: helpers.URLShaper(sc.BaseUrl, "/api/v1/"),
		Client:  sc.Client,
		Token:   helpers.TokenValidator(sc.Token, "Bearer "),
	}
}

func (s *SnipeClient) NewRequest(method, url string, body interface{}) *SnipeRespBody {
	var buf bytes.Buffer
	r := &SnipeRespBody{
		response: nil,
		body:     nil,
		errMsg:   nil,
	}

	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			r.errMsg = err
			return r
		}
	}

	req, err := http.NewRequest(method, s.BaseUrl+url, &buf)
	if err != nil {
		r.errMsg = err
		return r
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("Authorization", s.Token)

	resp, err := s.Client.Do(req)
	if err != nil {
		r.errMsg = err
		return r
	}
	r.response = resp

	snipeBody, err := io.ReadAll(resp.Body)
	if err != nil {
		r.errMsg = err
		return r
	}
	r.body = snipeBody

	return r
}
