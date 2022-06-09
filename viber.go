package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Viber struct {
	Client       http.Client
	Token        string
	Funcs        FuncList
	Sender       Sender
	APIVersion   int
	Broadcast    bool
	StartMessage string
}

type Sender struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type SessionParams struct {
	Token        string
	Funcs        FuncList
	Sender       Sender
	StartMessage string
}

func Session(p SessionParams) *Viber {
	if p.StartMessage == "" {
		p.StartMessage = "Привет!"
	}
	return &Viber{
		Client:       *http.DefaultClient,
		Token:        p.Token,
		Funcs:        p.Funcs,
		Sender:       p.Sender,
		StartMessage: p.StartMessage,
		APIVersion:   1,
	}
}

type GetAccountInfoResponse struct {
	Status        int    `json:"status"`
	StatusMessage string `json:"status_message"`
	ID            string `json:"id"`
	ChatHostname  string `json:"chat_hostname"`
	Name          string `json:"name"`
	URI           string `json:"uri"`
	Icon          string `json:"icon"`
	Category      string `json:"category"`
	Subcategory   string `json:"subcategory"`
	Location      struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"location"`
	Country    string   `json:"country"`
	Webhook    string   `json:"webhook"`
	EventTypes []string `json:"event_types"`
	Members    []struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Avatar string `json:"avatar"`
		Role   string `json:"role"`
	} `json:"members"`
	SubscribersCount int `json:"subscribers_count"`
}

//https://developers.viber.com/docs/api/rest-bot-api/#get-account-info
func (v Viber) GetAccountInfo() (GetAccountInfoResponse, error) {
	res, err := v.requset_api("get_account_info", nil)
	if err != nil {
		return GetAccountInfoResponse{}, err
	}

	check := v.GetError(res)
	if check != "ok" {
		return GetAccountInfoResponse{}, errors.New(check)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return GetAccountInfoResponse{}, err
	}

	var r GetAccountInfoResponse

	err = json.Unmarshal(body, &r)
	if err != nil {
		return GetAccountInfoResponse{}, err
	}

	return r, nil
}

type GetUserDetailsResponse struct {
	Status        int    `json:"status"`
	StatusMessage string `json:"status_message"`
	ChatHostname  string `json:"chat_hostname"`
	User          struct {
		ID              string `json:"id"`
		Name            string `json:"name"`
		Avatar          string `json:"avatar"`
		Language        string `json:"language"`
		Country         string `json:"country"`
		PrimaryDeviceOs string `json:"primary_device_os"`
		APIVersion      int    `json:"api_version"`
		ViberVersion    string `json:"viber_version"`
		Mcc             int    `json:"mcc"`
		Mnc             int    `json:"mnc"`
		DeviceType      string `json:"device_type"`
	} `json:"user"`
}

//https://developers.viber.com/docs/api/rest-bot-api/#get-user-details
func (v Viber) GetUserDetails(id string) (GetUserDetailsResponse, error) {
	param := struct {
		ID string `json:"id"`
	}{
		ID: id,
	}

	data, err := json.Marshal(param)
	if err != nil {
		return GetUserDetailsResponse{}, err
	}

	res, err := v.requset_api("get_user_details", data)
	if err != nil {
		return GetUserDetailsResponse{}, err
	}

	check := v.GetError(res)
	if check != "ok" {
		return GetUserDetailsResponse{}, errors.New(check)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return GetUserDetailsResponse{}, err
	}

	var r GetUserDetailsResponse

	err = json.Unmarshal(body, &r)
	if err != nil {
		return GetUserDetailsResponse{}, err
	}

	return r, nil
}

type GetOnlineResponse struct {
	Status        int    `json:"status"`
	StatusMessage string `json:"status_message"`
	Users         []struct {
		ID                  string `json:"id"`
		OnlineStatus        int    `json:"online_status"`
		OnlineStatusMessage string `json:"online_status_message"`
		LastOnline          int    `json:"last_online"`
	} `json:"users"`
}

//https://developers.viber.com/docs/api/rest-bot-api/#get-online
func (v Viber) GetOnline(ids []string) (GetOnlineResponse, error) {

	param := struct {
		IDs []string `json:"ids"`
	}{
		IDs: ids,
	}

	data, err := json.Marshal(param)
	if err != nil {
		return GetOnlineResponse{}, err
	}

	res, err := v.requset_api("get_online", data)
	if err != nil {
		return GetOnlineResponse{}, err
	}

	check := v.GetError(res)
	if check != "ok" {
		return GetOnlineResponse{}, errors.New(check)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return GetOnlineResponse{}, err
	}

	var r GetOnlineResponse

	err = json.Unmarshal(body, &r)
	if err != nil {
		return GetOnlineResponse{}, err
	}

	return r, nil
}
