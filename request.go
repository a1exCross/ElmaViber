package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

const api_host string = "https://chatapi.viber.com/pa/"

func (v *Viber) requset_api(method string, body []byte) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, api_host+method, bytes.NewBuffer(body))

	req.Header.Add("X-Viber-Auth-Token", v.Token)

	req.Close = true

	if err != nil {
		return nil, err
	}

	res, err := v.Client.Do(req)

	check := v.getError(res)
	if check != "ok" {
		return nil, errors.New(check)
	}

	return ioutil.ReadAll(res.Body)
}
