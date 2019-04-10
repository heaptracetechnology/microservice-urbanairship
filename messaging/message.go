package messaging

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	result "github.com/heaptracetechnology/microservice-urbanairship/result"
	urbanairship "github.com/heaptracetechnology/microservice-urbanairship/urbanairship"
)

type RequestParam struct {
	NamedUser    string   `json:"named_user,omitempty"`
	Tag          string   `json:"tag,omitempty"`
	ChannelId    string   `json:"channel_id,omitempty"`
	ChannelType  string   `json:"channel_type,omitempty"`
	DeviceTypes  []string `json:"device_list,omitempty"`
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

	if requestparam.ChannelId != "" {
		if requestparam.ChannelType != "" {
			if requestparam.ChannelType == "android" {
				audience.AndroidChannelId = requestparam.ChannelId
			} else if requestparam.ChannelType == "ios" {
				audience.IOSChannelId = requestparam.ChannelId
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

	responseWriter.Header().Set("Content-Type", "application/json")
	var appKey = os.Getenv("APP_KEY")
	var masterSecret = os.Getenv("MASTER_SECRET")

	message, channelType, err := TransfromRequestParamToMessage(request)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

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
