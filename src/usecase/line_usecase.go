package usecase

import (
	"github.com/pokemon-search-v2/src/infrastructure/line"
	"github.com/pokemon-search-v2/src/presentation/dto"
)

func ReplayMessage(dto *dto.LineMessageDto, message string) {
	repository := line.NewLineRepository()
	repository.ReplayMessage(dto, message)
}
