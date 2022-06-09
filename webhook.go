package ElmaViber

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type SetWebhookParam struct {
	Webhook   string  `json:"url"`
	Events    []Event `json:"event_types"`
	SendName  bool    `json:"send_name"`
	SendPhoto bool    `json:"send_photo"`
}

type WebhookResponse struct {
	Status        int    `json:"status"`
	StatusMessage string `json:"status_message"`
	ChatHostname  string `json:"chat_hostname"`
	MessageToken  int    `json:"message_token,omitempty"`
}

//https://developers.viber.com/docs/api/rest-bot-api/#setting-a-webhook
func (v *Viber) SetWebhook(p SetWebhookParam) error {
	body, err := json.Marshal(p)
	if err != nil {
		return err
	}

	res, err := v.requset_api("set_webhook", body)
	if err != nil {
		return err
	}

	check := v.GetError(res)
	if check != "ok" {
		return errors.New(check)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var r WebhookResponse

	err = json.Unmarshal(data, &r)
	if err != nil {
		return err
	}

	if r.StatusMessage == "ok" {
		log.Println("Webhook set sucessfull", r.ChatHostname)
	}

	return nil
}

func (v *Viber) HandleFunc(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	var e Events

	err = json.Unmarshal(data, &e)
	if err != nil {
		log.Println(err)
	}

	if e.Event == Webhook {
		w.Write([]byte(http.StatusText(http.StatusOK)))
		return
	}

	if e.Event == ConversationStarted {
		var e Events

		err := json.Unmarshal(data, &e)
		if err != nil {
			log.Println(err)
		}

		log.Println("Отправлено приветственное сообщение пользователю", e.User.ID)

		st := SendMessageTextParam{
			MessageParams: MessageParams{
				Receiver:      e.User.ID,
				MinAPIVersion: v.APIVersion,
				Sender:        v.Sender,
				Type:          Text,
			},
			Text: v.StartMessage,
		}

		data, err := json.Marshal(st)
		if err != nil {
			log.Println(err)
		}

		w.Write(data)
	}

	err = v.CallFuncList(data, e)
	if err != nil {
		log.Println(err)
	}

	w.Write([]byte(http.StatusText(http.StatusOK)))
}
