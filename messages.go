package main

import (
	"encoding/json"
	"errors"
)

func (v Viber) GetTypeMessage(m MessageEvent) string {
	var t MessageType

	json.Unmarshal(m.Message, &t)

	return t.Type
}

type MessageType struct {
	Type string `json:"type"`
}

type SendMessageTextParams struct {
	MessageParams
	Text string `json:"text"`
}

type SendMessagePictureParams struct {
	MessageParams
	Text      string `json:"text"`
	Media     string `json:"media"`
	Thumbnail string `json:"thumbnail,omitempty"`
}

type SendMessageVideoParams struct {
	MessageParams
	Media     string `json:"media"`
	Thumbnail string `json:"thumbnail"`
	Size      int    `json:"size"`
	Duration  int    `json:"duration"`
}

type SendMessageFileParams struct {
	MessageParams
	Media    string `json:"media"`
	Size     int    `json:"size"`
	Filename string `json:"file_name"`
}

type Contacts struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type SendMessageContactParams struct {
	MessageParams
	Contacts `json:"contact"`
}

type Locate struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

type SendMessageLocationParams struct {
	MessageParams
	Location Locate `json:"location"`
}

type SendMessageURLParams struct {
	URL string `json:"media"`
}

type SendMessageStickerParams struct {
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

func (v *Viber) SendMessageText(p SendMessageTextParams) (WebhookResponse, error) {
	if p.BroadcastList != nil {
		v.Broadcast = true
	}

	if p.Receiver == "" {
		return WebhookResponse{}, errors.New("Required field 'Reciever' is empty. Method: SendMessageText")
	}

	if p.Sender.Name == "" {
		return WebhookResponse{}, errors.New("Required field 'Sender.Name' is empty. Method: SendMessageText")
	}

	return v.sendMessage(SendMessageTextParams{
		MessageParams: MessageParams{
			MinAPIVersion: v.APIVersion,
			TrackingData:  p.TrackingData,
			Receiver:      p.Receiver,
			Sender: Sender{
				Name:   p.Sender.Name,
				Avatar: p.Sender.Avatar,
			},
			Type:          Text,
			BroadcastList: p.BroadcastList,
			Keyboard:      p.Keyboard,
		},
		Text: p.Text,
	})
}

func (v *Viber) SendMessagePicture(p SendMessagePictureParams) (WebhookResponse, error) {
	if p.BroadcastList != nil {
		v.Broadcast = true
	}

	if p.Receiver == "" {
		return WebhookResponse{}, errors.New("Required field 'Reciever' is empty. Method: SendMessagePicture")
	}

	if p.Sender.Name == "" {
		return WebhookResponse{}, errors.New("Required field 'Sender.Name' is empty. Method: SendMessagePicture")
	}

	return v.sendMessage(SendMessagePictureParams{
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

func (v *Viber) SendMessageVideo(p SendMessageVideoParams) (WebhookResponse, error) {
	if p.BroadcastList != nil {
		v.Broadcast = true
	}

	if p.Receiver == "" {
		return WebhookResponse{}, errors.New("Required field 'Reciever' is empty. Method: SendMessageVideo")
	}

	if p.Sender.Name == "" {
		return WebhookResponse{}, errors.New("Required field 'Sender.Name' is empty. Method: SendMessageVideo")
	}

	return v.sendMessage(SendMessageVideoParams{
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

func (v *Viber) SendMessageFile(p SendMessageFileParams) (WebhookResponse, error) {
	if p.BroadcastList != nil {
		v.Broadcast = true
	}

	if p.Receiver == "" {
		return WebhookResponse{}, errors.New("Required field 'Reciever' is empty. Method: SendMessageFile")
	}

	if p.Sender.Name == "" {
		return WebhookResponse{}, errors.New("Required field 'Sender.Name' is empty. Method: SendMessageFile")
	}

	return v.sendMessage(SendMessageFileParams{
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

func (v *Viber) SendMessageContact(p SendMessageContactParams) (WebhookResponse, error) {
	if p.BroadcastList != nil {
		v.Broadcast = true
	}

	if p.Receiver == "" {
		return WebhookResponse{}, errors.New("Required field 'Reciever' is empty. Method: SendMessageContact")
	}

	if p.Sender.Name == "" {
		return WebhookResponse{}, errors.New("Required field 'Sender.Name' is empty. Method: SendMessageContact")
	}

	return v.sendMessage(SendMessageContactParams{
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

func (v *Viber) SendMessageLocation(p SendMessageLocationParams) (WebhookResponse, error) {
	if p.BroadcastList != nil {
		v.Broadcast = true
	}

	if p.Receiver == "" {
		return WebhookResponse{}, errors.New("Required field 'Reciever' is empty. Method: SendMessageLocation")
	}

	if p.Sender.Name == "" {
		return WebhookResponse{}, errors.New("Required field 'Sender.Name' is empty. Method: SendMessageLocation")
	}

	return v.sendMessage(SendMessageLocationParams{
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

func (v *Viber) SendMessageSticker(p SendMessageStickerParams) (WebhookResponse, error) {
	if p.BroadcastList != nil {
		v.Broadcast = true
	}

	if p.Receiver == "" {
		return WebhookResponse{}, errors.New("Required field 'Reciever' is empty. Method: SendMessageSticker")
	}

	if p.Sender.Name == "" {
		return WebhookResponse{}, errors.New("Required field 'Sender.Name' is empty. Method: SendMessageSticker")
	}

	return v.sendMessage(SendMessageStickerParams{
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
func (v *Viber) sendMessage(p interface{}) (WebhookResponse, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return WebhookResponse{}, err
	}

	method := "send_message"

	if v.Broadcast {
		method = "broadcast_message"
		v.Broadcast = false
	}

	body, err := v.requset_api(method, data)
	if err != nil {
		return WebhookResponse{}, err
	}

	var r WebhookResponse

	err = json.Unmarshal(body, &r)
	if err != nil {
		return WebhookResponse{}, err
	}

	return r, nil
}
