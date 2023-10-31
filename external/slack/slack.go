package slack

type Slack interface {
	SendSmsMessage(channel string, receiver string, message string)
	SendLogMessage(channel string, endpoint string, requestBody string, response string)
	SendGeneralMessage(channel string, message string)
}
