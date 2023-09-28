package storage

import (
	"github.com/godofprodev/tally/api/models"
	"os/user"
)

type Storage interface {
	Connect() error
	Disconnect() error

	GetGuilds() ([]*models.Guild, error)
	GetGuildById(id int) (*models.Guild, error)
	GetUsers() ([]*models.User, error)
	GetUserById(id int) (*models.User, error)

	CreateGuild(guild *models.Guild) error
	CreateUser(user user.User) error
}
