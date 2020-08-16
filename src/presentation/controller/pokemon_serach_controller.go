package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/pokemon-search-v2/src/presentation/dto"
	"github.com/pokemon-search-v2/src/usecase"
)

func Search() echo.HandlerFunc {
	return func(c echo.Context) error {
		var lineDto dto.LineMessageDto
		if err := c.Bind(&lineDto); err != nil {
			fmt.Errorf("failed to parse json to DTO: %v", err)
			return c.String(http.StatusInternalServerError, "Server error")
		} else if len(lineDto.Events) == 0 {
			fmt.Errorf("failed to parse json to DTO: %v", err)
			return c.String(http.StatusInternalServerError, "Server error")
		} else {
			usecase.ReplayMessage(&lineDto, usecase.SearchByPokemonName(lineDto.Events[0].Message.Text))
			return c.String(http.StatusOK, "Reply line message is success")
		}
	}
}
