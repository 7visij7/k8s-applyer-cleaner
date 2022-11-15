package http

import (
	"fmt"
	"bytes"
	"net/http"
	"crypto/tls"
	"k8s-applyer-cleaner/pkg/errors"
)

func SetPlayload(payload string) *bytes.Buffer {
	if payload == "" {
		return bytes.NewBuffer([]byte(nil))
	}
	return bytes.NewBuffer([]byte(payload))
}

func HttpRequest(method string, namespace string, token string, url string, payload string) (*http.Response){
	req, err := http.NewRequest(method, url, SetPlayload(payload))
	errors.CheckError(err)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true},},}
	resp, err := client.Do(req)
	errors.CheckError(err)
	
	return resp
}