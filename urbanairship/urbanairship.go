package urbanairship

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// ua_server_url fcm server url
	ua_server_url = "https://go.urbanairship.com/api/push/"
)

var (
	// uaServerURL for testing purposes
	uaServerURL = ua_server_url
)

type Audiance struct {
	Tag       string `json:"tag,omitempty"`
	ChannelId string `json:"android_channel,omitempty"`
	NamedUser string `json:"named_user,omitempty"`
}

type Notification struct {
	Alert string `json:"alert,omitempty"`
}

type UAMessage struct {
	Audience     Audiance     `json:"audience,omitempty"`
	DeviceTypes  []string     `json:"device_types,omitempty"`
	Notification Notification `json:"notification,omitempty"`
}

// UrbanAirshipResponseStatus represents urban airship response message
type UAResponseStatus struct {
	Ok           bool        `json:"ok"`
	OperationId  string      `json:"operation_id"`
	PushIds      []string    `json:"push_ids"`
	MessageIds   []string    `json:"message_ids,omitempty"`
	ContentURLs  []string    `json:"content_urls,omitempty"`
	LocalizedIds []string    `json:"localized_ids,omitempty"`
	Error        string      `json:"error,omitempty"`
	ErrorCode    int         `json:"error_code,omitempty"`
	Details      interface{} `json:"details,omitempty"`
}

// UrbanAirshipClient struct
type UAClient struct {
	ApiKey        string
	MasterKey     string
	Authorization string
	Message       UAMessage
}

// NewUAClient generates the value of the Authorization key
func NewUAClient(apiKey string, masterKey string) *UAClient {
	ua := new(UAClient)
	ua.ApiKey = apiKey
	ua.MasterKey = masterKey
	generateAuth := apiKey + ":" + masterKey

	ua.Authorization = base64.StdEncoding.EncodeToString([]byte(generateAuth))

	return ua
}

// authorizationHeader generates the value of the Authorization key
func (this *UAClient) authorizationHeader() string {
	return fmt.Sprintf("key=%v", this.Authorization)
}

// NewUATagsMsg sets the targeted tagged devices
func (this *UAClient) NewUATagsMsg(authorizationKey string, tag string, deviceTypes []string, notification Notification) *UAClient {

	this.NewSendTagMsg(authorizationKey, tag, deviceTypes, notification)

	return this
}

// NewUANamedUserMsg sets the targeted nameuser
func (this *UAClient) NewUANamedUserMsg(authorizationKey string, namedUser string, deviceTypes []string, notification Notification) *UAClient {

	this.NewSendnamedUserMsg(authorizationKey, namedUser, deviceTypes, notification)

	return this
}

// NewUAChannelIdMsg sets the targeted to channelid
func (this *UAClient) NewUAChannelIdMsg(authorizationKey string, channelId string, deviceTypes []string, notification Notification) *UAClient {

	this.NewSendChannelIdMsg(authorizationKey, channelId, deviceTypes, notification)

	return this
}

// NewSendTagMsg sets the targeted tag and the data payload
func (this *UAClient) NewSendTagMsg(authorizationKey string, tag string, deviceTypes []string, notification Notification) *UAClient {

	this.Authorization = authorizationKey
	this.Message.Audience.Tag = tag
	this.Message.DeviceTypes = deviceTypes
	this.Message.Notification = notification

	return this
}

// NewSendnamedUserMsg sets the targeted nameduser and the data payload
func (this *UAClient) NewSendnamedUserMsg(authorizationKey string, namedUser string, deviceTypes []string, notification Notification) *UAClient {

	this.Authorization = authorizationKey
	this.Message.Audience.NamedUser = namedUser
	this.Message.DeviceTypes = deviceTypes
	this.Message.Notification = notification

	return this
}

// NewSendChannelIdMsg sets the targeted channelid and the data payload
func (this *UAClient) NewSendChannelIdMsg(authorizationKey string, channelId string, deviceTypes []string, notification Notification) *UAClient {

	this.Authorization = authorizationKey
	this.Message.Audience.ChannelId = channelId
	this.Message.DeviceTypes = deviceTypes
	this.Message.Notification = notification

	return this
}

// toJsonByte converts uaMsg to a json byte
func (this *UAMessage) toJsonByte() ([]byte, error) {

	return json.Marshal(this)

}

// parseStatusBody parse UA response body
func (this *UAResponseStatus) parseStatusBody(body []byte) error {

	if err := json.Unmarshal([]byte(body), &this); err != nil {
		return err
	}
	return nil

}

// sendOnce send a single request to ua
func (this *UAClient) sendOnce() (*UAResponseStatus, error) {

	uaResponseStatus := new(UAResponseStatus)

	jsonByte, err := this.Message.toJsonByte()
	if err != nil {
		return uaResponseStatus, err
	}

	request, err := http.NewRequest("POST", uaServerURL, bytes.NewBuffer(jsonByte))
	request.Header.Set("Authorization", "Basic "+this.Authorization)
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return uaResponseStatus, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return uaResponseStatus, err
	}

	err = uaResponseStatus.parseStatusBody(body)
	if err != nil {
		return uaResponseStatus, err
	}
	uaResponseStatus.Ok = true

	return uaResponseStatus, nil
}

// Send to ua
func (this *UAClient) Send() (*UAResponseStatus, error) {
	return this.sendOnce()

}
