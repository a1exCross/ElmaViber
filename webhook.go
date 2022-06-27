package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type SetWebhookParams struct {
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
func (v *Viber) SetWebhook(p SetWebhookParams) (WebhookResponse, error) {
	if p.Events == nil {
		return WebhookResponse{}, errors.New("Required field 'Events' is empty. Method: SetWebhook")
	}

	if p.Webhook == "" {
		return WebhookResponse{}, errors.New("Required field 'Webhook' is empty. Method: SetWebhook")
	}

	data, err := json.Marshal(p)
	if err != nil {
		return WebhookResponse{}, err
	}

	body, err := v.requset_api("set_webhook", data)
	if err != nil {
		return WebhookResponse{}, err
	}

	var r WebhookResponse

	err = json.Unmarshal(body, &r)
	if err != nil {
		return WebhookResponse{}, err
	}

	if r.StatusMessage == "ok" {
		log.Println("Webhook set sucessfull", r.ChatHostname)
	}

	return r, nil
}

func (v Viber) checkSignature(res *http.Request) bool {
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false
	}

	res.Body.Close()
	res.Body = ioutil.NopCloser(bytes.NewBuffer(data))

	h := hmac.New(sha256.New, []byte(v.Token))
	h.Write([]byte(data))

	return res.Header.Get("X-Viber-Content-Signature") == hex.EncodeToString(h.Sum(nil))
}

func (v *Viber) HandleFunc(w http.ResponseWriter, r *http.Request) {
	if v.checkSignature(r) {
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

			st := SendMessageTextParams{
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

		go func() {
			err = v.сallFuncList(data, e)
			if err != nil {
				log.Println(err)
			}
		}()

		w.Write([]byte(http.StatusText(http.StatusOK)))
	}
}
