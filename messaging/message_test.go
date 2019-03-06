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

	apiKey := "_i3ZHwoUSxKJzD_oA1QuCQ"
	masterKey := "rPOZp9WsQ1i-bQV6nYJpSA"

	os.Setenv("APP_KEY", apiKey)
	os.Setenv("MASTER_SECRET", masterKey)

	var requestParam RequestParam
	requestParam.Tag = "rohit-tag"
	requestParam.Notification = "Test to push on android using tag"
	requestParam.DeviceTypes = []string{"android"}

	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(requestParam)

	request, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(Send)
	handler.ServeHTTP(recorder, request)

	Describe("Send message by tag", func() {
		Context("Send", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("Urban Airship messaging, send by chanelid", func() {

	apiKey := "_i3ZHwoUSxKJzD_oA1QuCQ"
	masterKey := "rPOZp9WsQ1i-bQV6nYJpSA"

	os.Setenv("APP_KEY", apiKey)
	os.Setenv("MASTER_SECRET", masterKey)

	var requestParam RequestParam
	requestParam.ChannelId = "544e8079-8c10-448b-b4b6-3a7b00cf2a40"
	requestParam.Notification = "Test to push on android using chanelid"
	requestParam.DeviceTypes = []string{"android"}

	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(requestParam)

	request, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(Send)
	handler.ServeHTTP(recorder, request)

	Describe("Send message by chanelid", func() {
		Context("Send", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("Urban Airship messaging, send by named user", func() {

	apiKey := "_i3ZHwoUSxKJzD_oA1QuCQ"
	masterKey := "rPOZp9WsQ1i-bQV6nYJpSA"

	os.Setenv("APP_KEY", apiKey)
	os.Setenv("MASTER_SECRET", masterKey)

	var requestParam RequestParam
	requestParam.NamedUser = "rohit"
	requestParam.Notification = "Test to push on android using named user"
	requestParam.DeviceTypes = []string{"android"}

	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(requestParam)

	request, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(Send)
	handler.ServeHTTP(recorder, request)

	Describe("Send message by named user", func() {
		Context("Send", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})
