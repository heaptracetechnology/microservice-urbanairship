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

func TransfromRequestParamToMessage(r *http.Request) (urbanairship.UAMessage, error) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	var requestparam RequestParam

	var message urbanairship.UAMessage
	err = json.Unmarshal(body, &requestparam)
	var audiance urbanairship.Audiance
	var notification urbanairship.Notification

	if requestparam.Tag != "" {
		audiance.Tag = requestparam.Tag
	}

	if requestparam.ChannelId != "" {
		audiance.Channel_id = requestparam.ChannelId
	}

	if requestparam.NamedUser != "" {
		audiance.Named_user = requestparam.NamedUser
	}
	if requestparam.Notification != "" {
		notification.Alert = requestparam.Notification
	}

	message.Audience = audiance
	message.Notification = notification
	message.Device_types = requestparam.DeviceTypes

	return message, err
}
func Send(w http.ResponseWriter, r *http.Request) {
	var appkey = os.Getenv("APP_KEY")
	var mastersec = os.Getenv("MASTER_SECRET")
	message, err := TransfromRequestParamToMessage(r)

	if err != nil {
		result.WriteErrorResponse(w, err)
		return
	}

	client := *urbanairship.NewUAClient(appkey, mastersec)
	client.Message = message

	response, err := client.Send()

	if err != nil {
		result.WriteErrorResponse(w, err)
		return
	}

	bytes, _ := json.Marshal(response)
	result.WriteJsonResponse(w, bytes, http.StatusOK)
}
