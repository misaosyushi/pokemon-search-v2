package line

import (
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"

	repository "github.com/pokemon-search-v2/src/interfaces/line"
	"github.com/pokemon-search-v2/src/presentation/dto"
)

type LineRepository struct{}

func NewLineRepository() repository.ILineRepository {
	return &LineRepository{}
}

func (repository *LineRepository) ReplayMessage(dto *dto.LineMessageDto, message string) {
	bot, err := linebot.New(os.Getenv("CHANNEL_SECRET"), os.Getenv("CHANNEL_ACCESS_TOKEN"))
	if err != nil {
		log.Fatalf("Failed to create line client: %s", err)
	}
	if _, err := bot.ReplyMessage(dto.Events[0].ReplyToken, linebot.NewTextMessage(message)).Do(); err != nil {
		log.Printf("Failed to reply message: %s", err)
	}
}
