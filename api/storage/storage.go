package storage

import (
	"github.com/godofprodev/tally/api/models"
)

type Storage interface {
	Connect() error
	Disconnect() error

	GetGuilds() ([]*models.Guild, error)
	GetGuildById(id int) (*models.Guild, error)
	GetUsers() ([]*models.User, error)
	GetUserById(id int) (*models.User, error)

	CreateGuild(guild *models.Guild) error
	CreateUser(user *models.User) error

	UpdateGuild(guild *models.Guild) error
	UpdateUser(user *models.User) error
}
