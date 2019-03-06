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
	// uaServerUrl for testing purposes
	uaServerUrl = ua_server_url
)

type Audiance struct {
	Tag        string `json:"tag,omitempty"`
	Channel_id string `json:"android_channel,omitempty"`
	Named_user string `json:"named_user,omitempty"`
}

type Notification struct {
	Alert string `json:"alert,omitempty"`
}

type UAMessage struct {
	Audience     Audiance     `json:"audience,omitempty"`
	Device_types []string     `json:"device_types,omitempty"`
	Notification Notification `json:"notification,omitempty"`
}

// UrbanAirshipResponseStatus represents urban airship response message
type UAResponseStatus struct {
	Ok            bool
	StatusCode    int
	Operation_id  string      `json:"operation_id"`
	Push_ids      []string    `json:"push_ids"`
	Message_ids   []string    `json:"message_ids,omitempty"`
	Content_urls  []string    `json:"content_urls,omitempty"`
	Localized_ids []string    `json:"localized_ids,omitempty"`
	Error         string      `json:"error,omitempty"`
	Error_code    int         `json:"error_code,omitempty"`
	Details       interface{} `json:"details,omitempty"`
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
func (this *UAClient) NewUATagsMsg(authorizationKey string, tag string, devicetypes []string, notification Notification) *UAClient {

	this.NewSendTagMsg(authorizationKey, tag, devicetypes, notification)

	return this
}

// NewUANamedUserMsg sets the targeted nameuser
func (this *UAClient) NewUANamedUserMsg(authorizationKey string, nameduser string, devicetypes []string, notification Notification) *UAClient {

	this.NewSendnamedUserMsg(authorizationKey, nameduser, devicetypes, notification)

	return this
}

// NewUAChannelIdMsg sets the targeted to channelid
func (this *UAClient) NewUAChannelIdMsg(authorizationKey string, channelid string, devicetypes []string, notification Notification) *UAClient {

	this.NewSendChannelIdMsg(authorizationKey, channelid, devicetypes, notification)

	return this
}

// NewSendTagMsg sets the targeted tag and the data payload
func (this *UAClient) NewSendTagMsg(authorizationKey string, tag string, devicetypes []string, notification Notification) *UAClient {

	this.Authorization = authorizationKey
	this.Message.Audience.Tag = tag
	this.Message.Device_types = devicetypes
	this.Message.Notification = notification

	return this
}

// NewSendnamedUserMsg sets the targeted nameduser and the data payload
func (this *UAClient) NewSendnamedUserMsg(authorizationKey string, nameduser string, devicetypes []string, notification Notification) *UAClient {

	this.Authorization = authorizationKey
	this.Message.Audience.Named_user = nameduser
	this.Message.Device_types = devicetypes
	this.Message.Notification = notification

	return this
}

// NewSendChannelIdMsg sets the targeted channelid and the data payload
func (this *UAClient) NewSendChannelIdMsg(authorizationKey string, channelid string, devicetypes []string, notification Notification) *UAClient {

	this.Authorization = authorizationKey
	this.Message.Audience.Channel_id = channelid
	this.Message.Device_types = devicetypes
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

	uaRespStatus := new(UAResponseStatus)

	//jsonByte, err := this.Message.toJsonByte()
	jsonByte, err := this.Message.toJsonByte()
	fmt.Println("=============" + string(jsonByte))
	if err != nil {
		return uaRespStatus, err
	}
	//UAClient.data = this.Message

	request, err := http.NewRequest("POST", uaServerUrl, bytes.NewBuffer(jsonByte))
	request.Header.Set("Authorization", "Basic "+this.Authorization)
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return uaRespStatus, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return uaRespStatus, err
	}

	uaRespStatus.StatusCode = response.StatusCode

	//uaRespStatus.RetryAfter = response.Header.Get(retry_after_header)

	if response.StatusCode != 200 {
		return uaRespStatus, nil
	}

	err = uaRespStatus.parseStatusBody(body)
	if err != nil {
		return uaRespStatus, err
	}
	uaRespStatus.Ok = true

	return uaRespStatus, nil
}

// Send to ua
func (this *UAClient) Send() (*UAResponseStatus, error) {
	return this.sendOnce()

}
