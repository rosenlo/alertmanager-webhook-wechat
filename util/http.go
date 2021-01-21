package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

const (
	DefaultDialTimeout         = 1 * time.Second
	DefaultKeepAliveTimeout    = 30 * time.Second
	DefaultMaxIdleConnsPerHost = 10
)

var DefaultTransport = &http.Transport{
	MaxIdleConnsPerHost: DefaultMaxIdleConnsPerHost,
	Dial: (&net.Dialer{
		Timeout:   DefaultDialTimeout,
		KeepAlive: DefaultKeepAliveTimeout,
	}).Dial,
}

func RestRequest(method, url string, body interface{}, headers map[string]string) ([]byte, error) {
	data, err := json.Marshal(body)
	if err != nil {
		log.Printf("[error] jsonMarshal failed due to %s", err)
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{Transport: DefaultTransport}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[error] http post failed due to %s", err)
		return nil, err
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[error] could't read response body due to %s", err)
		return nil, err
	}

	log.Printf("[debug] response: %s", string(respBody))
	return respBody, nil
}
