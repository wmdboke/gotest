package weather

import (
	"context"
	"crypto/tls"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type SliceMock struct {
	addr uintptr
	len  int
	cap  int
}

// NewRequest new a http request.
func NewRequest(method, uri string, params url.Values) (req *http.Request, err error) {
	if method == "GET" {
		req, err = http.NewRequest(method, uri+"?"+params.Encode(), nil)
	} else {
		req, err = http.NewRequest(method, uri, strings.NewReader(params.Encode()))
	}
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return
}

func Do(c context.Context, req *http.Request, timeout time.Duration) (body []byte, err error) {
	var res *http.Response
	dialer := &net.Dialer{
		Timeout:   timeout, //time.Duration(timeout * int64(time.Second)),
		KeepAlive: time.Duration(1 * int64(time.Second)),
	}
	transport := &http.Transport{
		DialContext:     dialer.DialContext,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: transport,
	}
	req = req.WithContext(c)
	if res, err = client.Do(req); err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode >= http.StatusInternalServerError {
		err = errors.New("http status code 5xx")
		return
	}
	if body, err = ioutil.ReadAll(res.Body); err != nil {
		return
	}
	return
}
