package line

import (
	"fmt"
	"github.com/pokemon-search-v2/src/presentation/dto"
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
)

// TODO: POSTのボディから文字列を取得 -> Elasticsearchから検索
func ReplayMessage(dto dto.PokemonDto) {
	bot, err := linebot.New("", "")
	if err != nil {
		log.Fatal("")
	}
	if _, err := bot.ReplyMessage("", linebot.NewTextMessage(MakeMessage(dto))).Do(); err != nil {
		log.Fatal("")
	}
}

func MakeMessage(dto dto.PokemonDto) string {
	return fmt.Sprintf("タイプ: %s\n種族値: \nHP　　　: %s\nこうげき: %s\nぼうぎょ: %s\nとっこう: %s\nとくぼう: %s\nすばやさ: %s",
		dto.Type, dto.BaseStats.Hp, dto.BaseStats.Attack, dto.BaseStats.Defense, dto.BaseStats.SpecialAttack, dto.BaseStats.SpecialDefense, dto.BaseStats.Speed)
}
