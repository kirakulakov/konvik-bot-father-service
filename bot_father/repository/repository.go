package repository

import (
	models "github.com/kirakulakov/konvik-bot-father-service/bot_father"
)

type BotFatherRepository interface {
	GetById(int64) (*models.Bot, error)
}
