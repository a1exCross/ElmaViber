package main

import (
	"bytes"
	"net/http"
)

const api_host string = "https://chatapi.viber.com/pa/"

func (v *Viber) requset_api(method string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, api_host+method, bytes.NewBuffer(body))

	req.Header.Add("X-Viber-Auth-Token", v.Token)
	//req.Header.Set("")

	req.Close = true

	if err != nil {
		return nil, err
	}

	//log.Println(req.Body)

	return v.Client.Do(req)
}
