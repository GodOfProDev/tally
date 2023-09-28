package models

type User struct {
	UserId      int `bson:"userId"`
	TotalCounts int `bson:"totalCounts"`
}
