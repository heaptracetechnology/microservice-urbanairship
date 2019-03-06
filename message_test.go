package messaging

import (
	"bytes"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"os"
)

var _ = Describe("Urban Airship messaging, send by tag", func() {

	apikey := "_i3ZHwoUSxKJzD_oA1QuCQ"
	masterkey := "rPOZp9WsQ1i-bQV6nYJpSA"

	os.Setenv("APP_KEY", apikey)
	os.Setenv("MASTER_SECRET", masterkey)

	var requestparam RequestParam
	requestparam.Tag = "rohit-tag"
	requestparam.Notification = "Test to push on android using tag"
	requestparam.DeviceTypes = []string{"android"}

	reqbody := new(bytes.Buffer)
	json.NewEncoder(reqbody).Encode(requestparam)

	req, err := http.NewRequest("POST", "/send", reqbody)
	if err != nil {
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(Send)
	handler.ServeHTTP(recorder, req)

	Describe("Send message by tag", func() {
		Context("Send", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("Urban Airship messaging, send by chanelid", func() {

	apikey := "_i3ZHwoUSxKJzD_oA1QuCQ"
	masterkey := "rPOZp9WsQ1i-bQV6nYJpSA"

	os.Setenv("APP_KEY", apikey)
	os.Setenv("MASTER_SECRET", masterkey)

	var requestparam RequestParam
	requestparam.ChannelId = "544e8079-8c10-448b-b4b6-3a7b00cf2a40"
	requestparam.Notification = "Test to push on android using chanelid"
	requestparam.DeviceTypes = []string{"android"}

	reqbody := new(bytes.Buffer)
	json.NewEncoder(reqbody).Encode(requestparam)

	req, err := http.NewRequest("POST", "/send", reqbody)
	if err != nil {
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(Send)
	handler.ServeHTTP(recorder, req)

	Describe("Send message by chanelid", func() {
		Context("Send", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("Urban Airship messaging, send by named user", func() {

	apikey := "_i3ZHwoUSxKJzD_oA1QuCQ"
	masterkey := "rPOZp9WsQ1i-bQV6nYJpSA"

	os.Setenv("APP_KEY", apikey)
	os.Setenv("MASTER_SECRET", masterkey)

	var requestparam RequestParam
	requestparam.NamedUser = "rohit"
	requestparam.Notification = "Test to push on android using named user"
	requestparam.DeviceTypes = []string{"android"}

	reqbody := new(bytes.Buffer)
	json.NewEncoder(reqbody).Encode(requestparam)

	req, err := http.NewRequest("POST", "/send", reqbody)
	if err != nil {
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(Send)
	handler.ServeHTTP(recorder, req)

	Describe("Send message by named user", func() {
		Context("Send", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})
