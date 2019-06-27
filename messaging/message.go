package messaging

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	result "github.com/heaptracetechnology/microservice-urbanairship/result"
	urbanairship "github.com/heaptracetechnology/microservice-urbanairship/urbanairship"
)

type Message struct {
	Success    string `json:"success"`
	Message    string `json:"message"`
	StatusCode int    `json:"statuscode"`
}

type RequestParam struct {
	NamedUser    string   `json:"namedUser,omitempty"`
	Tag          string   `json:"tag,omitempty"`
	ChannelID    string   `json:"channelID,omitempty"`
	ChannelType  string   `json:"channelType,omitempty"`
	DeviceTypes  []string `json:"deviceList,omitempty"`
	Notification string   `json:"message,omitempty"`
}

func TransfromRequestParamToMessage(request *http.Request) (urbanairship.UAMessage, string, error) {

	var requestparam RequestParam
	var message urbanairship.UAMessage

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return message, "", err
	}

	defer request.Body.Close()

	er := json.Unmarshal(body, &requestparam)
	if er != nil {
		return message, requestparam.ChannelType, er
	}

	var audience urbanairship.Audience
	var notification urbanairship.Notification

	if requestparam.Tag != "" {
		audience.Tag = requestparam.Tag
	}

	if requestparam.ChannelID != "" {
		if requestparam.ChannelType != "" {
			if requestparam.ChannelType == "android" {
				audience.AndroidChannelId = requestparam.ChannelID
			} else if requestparam.ChannelType == "ios" {
				audience.IOSChannelId = requestparam.ChannelID
			}
		}
	}

	if requestparam.NamedUser != "" {
		audience.NamedUser = requestparam.NamedUser
	}
	if requestparam.Notification != "" {
		notification.Alert = requestparam.Notification
	}

	message.Audience = audience
	message.Notification = notification
	message.DeviceTypes = requestparam.DeviceTypes

	return message, requestparam.ChannelType, err
}
func Send(responseWriter http.ResponseWriter, request *http.Request) {

	var appKey = os.Getenv("APP_KEY")
	var masterSecret = os.Getenv("MASTER_SECRET")

	message, channelType, err := TransfromRequestParamToMessage(request)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}
	if message.Audience.Tag == "" && message.Audience.NamedUser == "" && message.Audience.IOSChannelId == "" && message.Audience.AndroidChannelId == "" {
		message := Message{"false", "Please provide any of this Tag/NamedUser/ChannelID", http.StatusBadRequest}
		bytes, _ := json.Marshal(message)
		result.WriteJsonResponse(responseWriter, bytes, http.StatusBadRequest)
		return
	}
	fmt.Println("channelType----------", message.Audience)
	client := *urbanairship.NewUAClient(appKey, masterSecret, channelType)
	client.Message = message

	response, er := client.Send()
	if er != nil {
		result.WriteErrorResponse(responseWriter, er)
		return
	}

	bytes, _ := json.Marshal(response)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}
