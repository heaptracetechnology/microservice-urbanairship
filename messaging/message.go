package messaging

import (
	"encoding/json"
	result "github.com/heaptracetechnology/microservice-urbanairship/result"
	urbanairship "github.com/heaptracetechnology/microservice-urbanairship/urbanairship"
	"io/ioutil"
	"net/http"
	"os"
)

type RequestParam struct {
	NamedUser    string   `json:"named_user,omitempty"`
	Tag          string   `json:"tag,omitempty"`
	ChannelId    string   `json:"channel_id,omitempty"`
	DeviceTypes  []string `json:"device_list,omitempty"`
	Notification string   `json:"message,omitempty"`
}

func TransfromRequestParamToMessage(request *http.Request) (urbanairship.UAMessage, error) {
	body, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()

	var requestparam RequestParam

	var message urbanairship.UAMessage
	err = json.Unmarshal(body, &requestparam)
	var audiance urbanairship.Audiance
	var notification urbanairship.Notification

	if requestparam.Tag != "" {
		audiance.Tag = requestparam.Tag
	}

	if requestparam.ChannelId != "" {
		audiance.ChannelId = requestparam.ChannelId
	}

	if requestparam.NamedUser != "" {
		audiance.NamedUser = requestparam.NamedUser
	}
	if requestparam.Notification != "" {
		notification.Alert = requestparam.Notification
	}

	message.Audience = audiance
	message.Notification = notification
	message.DeviceTypes = requestparam.DeviceTypes

	return message, err
}
func Send(responseWriter http.ResponseWriter, request *http.Request) {
	var appKey = os.Getenv("APP_KEY")
	var masterSecret = os.Getenv("MASTER_SECRET")
	message, err := TransfromRequestParamToMessage(request)

	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	client := *urbanairship.NewUAClient(appKey, masterSecret)
	client.Message = message

	response, err := client.Send()

	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	bytes, _ := json.Marshal(response)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}
