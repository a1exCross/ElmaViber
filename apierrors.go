package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (v *Viber) getError(r *http.Response) string {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err.Error()
	}

	r.Body.Close()

	r.Body = ioutil.NopCloser(bytes.NewBuffer(data))

	var b WebhookResponse

	err = json.Unmarshal(data, &b)
	if err != nil {
		return err.Error()
	}

	if b.StatusMessage != "ok" && b.Status != 0 {
		return fmt.Sprintf("Error '%d': '%s', chat_hostname: '%s'", b.Status, b.StatusMessage, b.ChatHostname)
	}

	return "ok"
}
