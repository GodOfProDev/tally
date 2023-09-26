package storage

import (
	"github.com/godofprodev/tally/api/models"
	"os/user"
)

type Storage interface {
	GetGuilds() ([]*models.Guild, error)
	GetGuildById() (*models.Guild, error)
	GetUsers() ([]*models.User, error)
	GetUserById() (*models.User, error)

	CreateGuild(guild *models.Guild) error
	CreateUser(user user.User) error
}
