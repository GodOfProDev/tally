package storage

import (
	"context"
	"github.com/godofprodev/tally/api/config"
	"github.com/godofprodev/tally/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(m.Cfg.Uri))
	if err != nil {
		return err
	}

	m.UsersCollection = client.Database("tally").Collection("users")
	m.GuildsCollection = client.Database("tally").Collection("guilds")

	m.Client = client

	return nil
}

func (m *MongoStore) Disconnect() error {
	return m.Client.Disconnect(context.Background())
}

func (m *MongoStore) GetGuilds() ([]*models.Guild, error) {
	var query bson.M

	cursor, err := m.GuildsCollection.Find(context.Background(), query)
	if err != nil {
		return nil, err
	}

	var results []*models.Guild
	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	return results, err
}

func (m *MongoStore) GetGuildById(id int) (*models.Guild, error) {
	filter := bson.D{{"serverId", id}}

	var result *models.Guild
	err := m.GuildsCollection.FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		return nil, err
	}

	return result, err
}

func (m *MongoStore) GetUsers() ([]*models.User, error) {
	var query bson.M

	cursor, err := m.UsersCollection.Find(context.Background(), query)
	if err != nil {
		return nil, err
	}

	var results []*models.User
	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	return results, err
}

func (m *MongoStore) GetUserById(id int) (*models.User, error) {
	filter := bson.D{{"userId", id}}

	var result *models.User
	err := m.UsersCollection.FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		return nil, err
	}

	return result, err
}

func (m *MongoStore) CreateGuild(guild *models.Guild) error {
	_, err := m.GuildsCollection.InsertOne(context.Background(), guild)
	if err != nil {
		return err
	}

	return nil
}

func (m *MongoStore) CreateUser(user *models.User) error {
	_, err := m.UsersCollection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	return nil
}

func (m *MongoStore) UpdateGuild(guild *models.Guild) error {
	filter := bson.D{{"serverId", guild.ServerId}}
	update := bson.D{{"$set", bson.D{
		{"currentCount", guild.CurrentCount},
		{"highestCount", guild.HighestCount},
		{"channelId", guild.ChannelId},
	}}}

	_, err := m.GuildsCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (m *MongoStore) UpdateUser(user *models.User) error {
	filter := bson.D{{"userId", user.UserId}}
	update := bson.D{{"$set", bson.D{
		{"totalCounts", user.TotalCounts},
	}}}

	_, err := m.UsersCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
