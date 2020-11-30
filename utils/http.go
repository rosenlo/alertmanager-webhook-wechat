package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/rosenlo/toolkits/log"
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

func RestRequest(method, url string, body interface{}, headers map[string]string) error {
	log := log.WithField("method", method).WithField("url", url)
	data, err := json.Marshal(body)
	if err != nil {
		log.Error("jsonMarshal Failed due to ", err)
		return err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{Transport: DefaultTransport}
	resp, err := client.Do(req)
	if err != nil {
		log.Error("http post failed due to ", err)
		return err
	}
	log.WithField("status", resp.Status)

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return err
	}
	log.WithField("response", string(respBody)).Debug()
	return nil
}
