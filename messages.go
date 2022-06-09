package ElmaViber

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

func (v Viber) GetTypeMessage(m MessageEvent) string {
	var t MessageType

	json.Unmarshal(m.Message, &t)

	return t.Type
}

type MessageType struct {
	Type string `json:"type"`
}

type SendMessageTextParam struct {
	MessageParams
	Text string `json:"text"`
}

type SendMessagePictureParam struct {
	MessageParams
	Text      string `json:"text"`
	Media     string `json:"media"`
	Thumbnail string `json:"thumbnail,omitempty"`
}

type SendMessageVideoParam struct {
	MessageParams
	Media     string `json:"media"`
	Thumbnail string `json:"thumbnail"`
	Size      int    `json:"size"`
	Duration  int    `json:"duration"`
}

type SendMessageFileParam struct {
	MessageParams
	Media    string `json:"media"`
	Size     int    `json:"size"`
	Filename string `json:"file_name"`
}

type Contacts struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type SendMessageContactParam struct {
	MessageParams
	Contacts `json:"contact"`
}

type Locate struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

type SendMessageLocationParam struct {
	MessageParams
	Location Locate `json:"location"`
}

type SendMessageURLParam struct {
	URL string `json:"media"`
}

type SendMessageStickerParam struct {
	MessageParams
	StickerID int `json:"sticker_id"`
}

type MessageParams struct {
	Receiver      string    `json:"receiver,omitempty"`
	MinAPIVersion int       `json:"min_api_version,omitempty"`
	Sender        Sender    `json:"sender"`
	TrackingData  string    `json:"tracking_data,omitempty"`
	Type          string    `json:"type"`
	Keyboard      *Keyboard `json:"keyboard,omitempty"`
	BroadcastList []string  `json:"broadcast_list,omitempty"`
	/* 	StickerMessage  *SendMessageStickerParam  `json:",omitempty"`
	   	URLMessage      *SendMessageURLParam      `json:",omitempty"`
	   	LocationMessage *SendMessageLocationParam `json:",omitempty"`
	   	ContactMessage  *SendMessageContactParam  `json:",omitempty"`
	   	FileMessage     *SendMessageFileParam     `json:",omitempty"`
	   	VideoMessage    *SendMessageVideoParam    `json:",omitempty"`
	   	PictureMessage  *SendMessagePictureParam  `json:",omitempty"`
	   	TextMessage     *SendMessageTextParam     `json:",omitempty"` */
}

func (v *Viber) SendMessageText(p SendMessageTextParam) error {
	if p.BroadcastList != nil {
		v.Broadcast = true
	}

	return v.sendMessage(SendMessageTextParam{
		MessageParams: MessageParams{
			MinAPIVersion: v.APIVersion,
			TrackingData:  p.TrackingData,
			Receiver:      p.Receiver,
			Sender: Sender{
				Name:   v.Sender.Name,
				Avatar: v.Sender.Avatar,
			},
			Type:          Text,
			BroadcastList: p.BroadcastList,
			Keyboard:      p.Keyboard,
		},
		Text: p.Text,
	})
}

func (v *Viber) SendMessagePicture(p SendMessagePictureParam) {
	v.sendMessage(SendMessagePictureParam{
		MessageParams: MessageParams{
			MinAPIVersion: v.APIVersion,
			Type:          Picture,
			Receiver:      p.Receiver,
			Sender:        p.Sender,
			TrackingData:  p.TrackingData,
		},
		Text:      p.Text,
		Media:     p.Media,
		Thumbnail: p.Thumbnail,
	})
}

func (v *Viber) SendMessageVideo(p SendMessageVideoParam) {
	v.sendMessage(SendMessageVideoParam{
		MessageParams: MessageParams{
			Receiver:      p.Receiver,
			Sender:        p.Sender,
			MinAPIVersion: v.APIVersion,
			TrackingData:  p.TrackingData,
			Type:          Video,
		},
		Media:     p.Media,
		Thumbnail: p.Thumbnail,
		Size:      p.Size,
		Duration:  p.Duration,
	})
}

func (v *Viber) SendMessageFile(p SendMessageFileParam) {
	v.sendMessage(SendMessageFileParam{
		MessageParams: MessageParams{
			Receiver:      p.Receiver,
			Sender:        p.Sender,
			MinAPIVersion: v.APIVersion,
			TrackingData:  p.TrackingData,
			Type:          File,
		},
		Media:    p.Media,
		Size:     p.Size,
		Filename: p.Filename,
	})
}

func (v *Viber) SendMessageContact(p SendMessageContactParam) {
	v.sendMessage(SendMessageContactParam{
		MessageParams: MessageParams{
			Receiver:      p.Receiver,
			MinAPIVersion: v.APIVersion,
			Sender:        p.Sender,
			TrackingData:  p.TrackingData,
			Type:          Contact,
		},
		Contacts: p.Contacts,
	})
}

func (v *Viber) SendMessageLocation(p SendMessageLocationParam) error {
	return v.sendMessage(SendMessageLocationParam{
		MessageParams: MessageParams{
			Receiver:      p.Receiver,
			Sender:        p.Sender,
			MinAPIVersion: v.APIVersion,
			TrackingData:  p.TrackingData,
			Type:          Location,
		},
		Location: p.Location,
	})
}

func (v *Viber) SendMessageSticker(p SendMessageStickerParam) {
	v.sendMessage(SendMessageStickerParam{
		MessageParams: MessageParams{
			Receiver:      p.Receiver,
			MinAPIVersion: v.APIVersion,
			Sender:        p.Sender,
			TrackingData:  p.TrackingData,
			Type:          Sticker,
		},
		StickerID: p.StickerID,
	})
}

//https://developers.viber.com/docs/api/rest-bot-api/#send-message
func (v *Viber) sendMessage(p interface{}) error {
	/* if p.Sender.Name == "" {
		p.Sender.Name = v.Sender.Name
	}

	p.Type = "text"

	p.MinAPIVersion = 1 */

	body, err := json.Marshal(p)
	if err != nil {
		return err
	}

	method := "send_message"

	if v.Broadcast {
		method = "broadcast_message"
		v.Broadcast = false
	}

	res, err := v.requset_api(method, body)
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

	return nil
}
