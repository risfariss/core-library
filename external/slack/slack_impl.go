package slack

import (
	"bitbucket.org/kawancicil/core-library/external/slack/request"
	"bitbucket.org/kawancicil/core-library/rest_api"
	"encoding/json"
	"fmt"
	"net/http"
)

type SlackUtils struct {
	restApi rest_api.RestApi
}

func InitSlackUtils(restApi rest_api.RestApi) Slack {
	return &SlackUtils{
		restApi: restApi,
	}
}

func (s *SlackUtils) SendSmsMessage(channel string, receiver string, message string) {
	slackRequestPayload := request.SlackRequest{}
	slackRequestPayload.Channel = channel
	slackRequestPayload.Username = "sms"
	slackRequestPayload.IconEmoji = ":incoming_envelope:"
	slackRequestPayload.Text = fmt.Sprintf("Hi, you have a new message replacement for sms\n"+
		"*Receiver:* %s\n*Message:* %s", receiver, message)
	s.sendToSlackChannel(slackRequestPayload)
}

func (s *SlackUtils) SendLogMessage(channel string, endpoint string, requestBody string, response string) {
	slackRequestPayload := request.SlackRequest{}
	logPayload := request.SlackLogPayloadRequest{
		Endpoint: endpoint,
		Request:  requestBody,
		Response: response,
	}
	payloadMarshal, _ := json.Marshal(logPayload)
	slackRequestPayload.Channel = channel
	slackRequestPayload.Username = "log"
	slackRequestPayload.Text = "```" + string(payloadMarshal) + "```"
	slackRequestPayload.IconEmoji = ":space_invader:"
	s.sendToSlackChannel(slackRequestPayload)
}

func (s *SlackUtils) SendGeneralMessage(channel string, message string) {
	slackRequestPayload := request.SlackRequest{}
	slackRequestPayload.Channel = channel
	slackRequestPayload.Username = "general"
	slackRequestPayload.Text = message
	slackRequestPayload.IconEmoji = ":memo:"
	s.sendToSlackChannel(slackRequestPayload)
}

func (s *SlackUtils) sendToSlackChannel(req request.SlackRequest) {
	slackWebhookUrl := "https://hooks.slack.com/services/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"
	requestMarshal, _ := json.Marshal(req)
	var requestBody map[string]interface{}
	_ = json.Unmarshal(requestMarshal, &requestBody)
	params := rest_api.NewParams(slackWebhookUrl, "", "",
		requestBody, nil, nil, "")
	_, _ = s.restApi.HitApiServiceGeneral(http.MethodPost, params)
}
