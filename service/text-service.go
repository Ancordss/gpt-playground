package service

import (
	"fmt"

	"github.com/Ancordss/gpt-playground/entity"
)

type TextService interface {
	Ask(entity.Text) entity.Text
	Gpt() []entity.Text
}

type textService struct {
	texts []entity.Text
}

func New1() TextService {
	return &textService{}
}

func (service *textService) Ask(texts entity.Text) entity.Text {
	fmt.Println(texts)
	//service.texts = append(service.texts, texts)
	return texts
}

func (service *textService) Gpt() []entity.Text {
	return service.texts
}
