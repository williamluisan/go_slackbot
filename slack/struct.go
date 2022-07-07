package slack

type AppMentionPost struct {
	Channel string `json:"channel"`
	Text string `json:"text"`
}