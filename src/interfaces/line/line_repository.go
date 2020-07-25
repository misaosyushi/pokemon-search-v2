package line

import "github.com/pokemon-search-v2/src/presentation/dto"

type ILineRepository interface {
	ReplayMessage(dto *dto.LineMessageDto, message string)
}
