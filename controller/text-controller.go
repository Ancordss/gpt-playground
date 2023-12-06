package controller

import (
	"github.com/Ancordss/gpt-playground/entity"
	"github.com/Ancordss/gpt-playground/service"
	"github.com/gin-gonic/gin"
)

type TextController interface {
	Ask(ctx *gin.Context) entity.Text
	Gpt() []entity.Text
}

type tcontroller struct {
	service service.TextService
}

func (c *tcontroller) Gpt() []entity.Text {
	return c.service.Gpt()
}

func New1(service service.TextService) TextController {
	return &tcontroller{
		service: service,
	}
}

func (c *tcontroller) Ask(ctx *gin.Context) entity.Text {

	var data entity.Text
	ctx.BindJSON(&data)
	c.service.Ask(data)
	return data
}
