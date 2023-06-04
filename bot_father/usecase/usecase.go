package usecase

import (
	models "github.com/kirakulakov/konvik-bot-father-service/bot"
	"github.com/kirakulakov/konvik-bot-father-service/bot/repository"
)

type BotFatherUsecase interface {
	GetById(int64) (*models.Bot, error)
}

type botFatherUsecase struct {
	botFatherRepos repository.BotFatherRepository
}

func (b *botFatherUsecase) GetById(id int64) (*models.Bot, error) {
	return b.botFatherRepos.GetById(id)
}
