package ElmaViber

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (v *Viber) GetError(r *http.Response) string {
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

	//log.Println(b)

	if b.StatusMessage != "ok" && b.Status != 0 {
		//return "Error " + strconv.Itoa(b.Status) + ": " + b.StatusMessage + ", chat_hostname: " + b.ChatHostname
		return fmt.Sprintf("Error '%d': '%s', chat_hostname: '%s'", b.Status, b.StatusMessage, b.ChatHostname)
	}

	return "ok"
}
