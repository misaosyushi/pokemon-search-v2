package line

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
)

func ReplayMessage() {
	bot, err := linebot.New("", "")
	if err != nil {
		log.Fatal("")
	}
	if _, err := bot.ReplyMessage("", linebot.NewTextMessage("hello")).Do(); err != nil {
		log.Fatal("")
	}
}

func createMessage() string{
	return ""
}
