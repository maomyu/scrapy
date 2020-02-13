package gofish

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Request struct {
	Url     string
	Methord string
	Headers *http.Header
	Body    io.Reader
	Handle  Handle
	Client  http.Client
}

func (r *Request) Do() error {
	req, err := http.NewRequest(r.Methord, r.Url, r.Body)
	if err != nil {
		return err
	}
	req.Header = *r.Headers

	resp, err := r.Client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error status code： %d", resp.StatusCode)
	}
	r.Handle.Worker(resp.Body, r.Url)
	defer resp.Body.Close()
	return nil
}
func NewRequest(method, Url, userAgent string, handle Handle, body io.Reader) (*Request, error) {
	// 检验网址的合法性
	_, err := url.Parse(Url)
	if err != nil {
		return nil, err
	}
	hdr := http.Header{}
	if userAgent != "" {
		hdr.Add("User-Agent", userAgent)
	} else {
		hdr.Add("User-Agent", UserAgent)
	}
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return nil
		},
	}

	return &Request{
		Url:     Url,
		Methord: method,
		Headers: &hdr,
		Handle:  handle,
		Body:    body,
		Client:  client,
	}, nil
}
