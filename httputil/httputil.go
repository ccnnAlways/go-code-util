package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// sendHttp 发送请求
func sendHttp(m, u string, data interface{}) (body []byte, err error) {
	client := http.Client{}

	var up *url.URL
	up, err = url.Parse(u)

	if err != nil {
		return
	}

	var r io.Reader = nil
	if strings.ToLower(strings.TrimSpace(m)) == "get" {
		if data != nil {
			v, ok := data.(url.Values)
			if ok {
				up.RawQuery = v.Encode()
			}
		}
	} else if strings.ToLower(strings.TrimSpace(m)) == "post" {
		var d []byte = make([]byte, 0)
		d, err = json.Marshal(data)
		if err != nil {
			return
		}
		r = bytes.NewReader(d)
	} else {
		err = errors.New("method is not supported")
		return
	}

	var req *http.Request
	req, err = http.NewRequest(m, up.String(), r)
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		return
	}
	var resp *http.Response
	resp, err = client.Do(req)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = errors.New("request is bad")
		return
	}

	body, err = ioutil.ReadAll(resp.Body)
	return
}

// SendGet 发送get请求
func SendGet(u string, data url.Values) (body []byte, err error) {
	return sendHttp("GET", u, data)
}

// SendPost 发送post请求
func SendPost(u string, data interface{}) (body []byte, err error) {
	return sendHttp("POST", u, data)
}
