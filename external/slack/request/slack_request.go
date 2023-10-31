package request

type SlackRequest struct {
	Channel   string `json:"channel"`
	Username  string `json:"username"`
	Text      string `json:"text"`
	IconEmoji string `json:"icon_emoji"`
}

type SlackLogPayloadRequest struct {
	Endpoint string `json:"endpoint"`
	Request  string `json:"request"`
	Response string `json:"response"`
}
