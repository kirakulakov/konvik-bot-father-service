package repository

import (
	"database/sql"

	models "github.com/kirakulakov/konvik-bot-father-service/bot_father"
)

func NewPsqlBotFatherRepository(Conn *sql.DB) BotFatherRepository {
	return &psqlBotFatherRepository{Conn: Conn}
}

type psqlBotFatherRepository struct {
	Conn *sql.DB
}

func (p *psqlBotFatherRepository) GetById(id int64) (*models.Bot, error) {
	// rows, err := p.Conn.Query("")
	return &models.Bot{ID: 10, CreatedAt: 10101010, UpdatedAt: 10101010, Name: "Test"}, nil

}
