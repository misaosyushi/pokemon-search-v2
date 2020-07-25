package dto

type LineMessageDto struct {
	Events []Event `json:"events"`
}

type Event struct {
	ReplyToken string `json:"replyToken"`
	Message    struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"message"`
}
