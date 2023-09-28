package storage

import (
	"context"
	"github.com/godofprodev/tally/api/config"
	"github.com/godofprodev/tally/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os/user"
)

type MongoStore struct {
	Cfg              *config.DBConfig
	Client           *mongo.Client
	GuildsCollection *mongo.Collection
	UsersCollection  *mongo.Collection
}

func NewMongoStore(cfg *config.DBConfig) *MongoStore {
	return &MongoStore{
		Cfg: cfg,
	}
}

func (m *MongoStore) Connect() error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.Cfg.Uri))
	if err != nil {
		return err
	}

	m.UsersCollection = client.Database("tally").Collection("users")
	m.GuildsCollection = client.Database("tally").Collection("guilds")

	m.Client = client

	return nil
}

func (m *MongoStore) Disconnect() error {
	return m.Client.Disconnect(context.TODO())
}

func (m *MongoStore) GetGuilds() ([]*models.Guild, error) {
	var query bson.M

	cursor, err := m.GuildsCollection.Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}

	var results []*models.Guild
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, err
}

func (m *MongoStore) GetGuildById(id int) (*models.Guild, error) {
	filter := bson.D{{"serverId", id}}

	var result *models.Guild
	err := m.GuildsCollection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return nil, err
	}

	return result, err
}

func (m *MongoStore) GetUsers() ([]*models.User, error) {
	var query bson.M

	cursor, err := m.UsersCollection.Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}

	var results []*models.User
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, err
}

func (m *MongoStore) GetUserById(id int) (*models.User, error) {
	filter := bson.D{{"userId", id}}

	var result *models.User
	err := m.GuildsCollection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return nil, err
	}

	return result, err
}

func (m *MongoStore) CreateGuild(guild *models.Guild) error {
	_, err := m.GuildsCollection.InsertOne(context.TODO(), guild)
	if err != nil {
		return err
	}

	return nil
}

func (m *MongoStore) CreateUser(user user.User) error {
	_, err := m.UsersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	return nil
}
