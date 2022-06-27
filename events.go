package main

import (
	"encoding/json"
)

type FuncList struct {
	NewMessage       func(e MessageEvent)
	SeenMessage      func(e Events)
	DeliveredMessage func(e Events)
	Subscribed       func(e Events)
	Unsibcribed      func(e Events)
	Failed           func(e Events)
}

type Event string

const (
	Delivered           Event = "delivered"
	Message             Event = "message"
	Subscribed          Event = "subscribed"
	Unsubscribed        Event = "unsubscribed"
	Seen                Event = "seen"
	Failed              Event = "falied"
	Webhook             Event = "webhook"
	ConversationStarted Event = "conversation_started"
)

type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
	Language   string `json:"language"`
	Country    string `json:"country"`
	APIVersion int    `json:"api_version"`
}

type Events struct {
	Event         Event  `json:"event"`
	Timestamp     int64  `json:"timestamp"`
	ChatHostname  string `json:"chat_hostname"`
	MessageToken  int64  `json:"message_token"`
	User          *User  `json:"user,omitempty"`
	UserID        string `json:"user_id,omitempty"`
	Desc          string `json:"desc,omitempty"`
	StatusMessage string `json:"status_message,omitempty"`
}

const (
	Text     string = "text"
	Picture  string = "picture"
	Video    string = "video"
	File     string = "file"
	Contact  string = "contact"
	Location string = "location"
	URL      string = "url"
	Sticker  string = "sticker"
)

type MessageText struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

type MessagePicture struct {
	Type      string `json:"type"`
	Media     string `json:"media"`
	Thumbnail string `json:"thumbnail"`
	FileName  string `json:"file_name"`
}

type MessageVideo struct {
	Type      string `json:"type"`
	Media     string `json:"media"`
	Thumbnail string `json:"thumbnail"`
	FileName  string `json:"file_name"`
	Size      int    `json:"size"`
	Duration  int    `json:"duration,omitempty"`
}

type MessageFile struct {
	Type     string `json:"type"`
	Media    string `json:"media"`
	FileName string `json:"file_name"`
	Size     int    `json:"size"`
}

type MessageContact struct {
	Type    string `json:"type"`
	Contact struct {
		Name        string `json:"name"`
		PhoneNumber string `json:"phone_number"`
		Avatar      string `json:"avatar"`
	} `json:"contact"`
}

type MessageLocation struct {
	Type     string `json:"type"`
	Location struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"location"`
}

type MessageSticker struct {
	Type      string `json:"type"`
	Media     string `json:"media"`
	StickerID int    `json:"sticker_id"`
}

type MessageURL struct {
}

//Надо объединить как то типы
type MessageEvent struct {
	Event        Event           `json:"event"`
	Timestamp    int64           `json:"timestamp"`
	ChatHostname string          `json:"chat_hostname"`
	MessageToken int64           `json:"message_token"`
	Sender       User            `json:"sender"`
	Message      json.RawMessage `json:"message"`
	Text         MessageText
	Picture      MessagePicture
	Video        MessageVideo
	File         MessageFile
	Contact      MessageContact
	Location     MessageLocation
	Sticker      MessageSticker
	Silent       bool `json:"silent"`
}

func (v Viber) сallbackMessageFunc(data []byte, e MessageEvent) error {
	t := v.GetTypeMessage(e)

	switch t {
	case Text:
		{
			var t MessageText

			err := json.Unmarshal(e.Message, &t)
			if err != nil {
				return err
			}

			e.Text = t

			go v.Funcs.NewMessage(e)
		}
	case Picture:
		{
			var p MessagePicture

			err := json.Unmarshal(e.Message, &p)
			if err != nil {
				return err
			}

			e.Picture = p

			go v.Funcs.NewMessage(e)
		}
	case Video:
		{
			var vv MessageVideo

			err := json.Unmarshal(e.Message, &vv)
			if err != nil {
				return err
			}

			e.Video = vv

			go v.Funcs.NewMessage(e)
		}

	case File:
		{
			var f MessageFile

			err := json.Unmarshal(e.Message, &f)
			if err != nil {
				return err
			}

			e.File = f

			go v.Funcs.NewMessage(e)
		}
	case Contact:
		{
			var c MessageContact

			err := json.Unmarshal(e.Message, &c)
			if err != nil {
				return err
			}

			e.Contact = c

			go v.Funcs.NewMessage(e)
		}
	case Location:
		{
			var l MessageLocation

			err := json.Unmarshal(e.Message, &l)
			if err != nil {
				return err
			}

			e.Location = l

			go v.Funcs.NewMessage(e)
		}

	case Sticker:
		{
			var s MessageSticker

			err := json.Unmarshal(e.Message, &s)
			if err != nil {
				return err
			}

			e.Sticker = s

			go v.Funcs.NewMessage(e)
		}
	}

	return nil
}

func (v *Viber) сallFuncList(data []byte, ev Events) error {
	switch ev.Event {
	case Message:
		{
			var e MessageEvent

			err := json.Unmarshal(data, &e)
			if err != nil {
				return err
			}

			go v.сallbackMessageFunc(data, e)
		}
	case Seen:
		{
			go v.Funcs.SeenMessage(ev)
		}

	case Delivered:
		{
			go v.Funcs.DeliveredMessage(ev)
		}
	case Subscribed:
		{
			go v.Funcs.Subscribed(ev)
		}
	case Unsubscribed:
		{
			go v.Funcs.Unsibcribed(ev)
		}
	case Failed:
		{
			go v.Funcs.Failed(ev)
		}
	}

	return nil
}
