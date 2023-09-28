package models

type User struct {
	UserId      int `bson:"userId" json:"userId"`
	TotalCounts int `bson:"totalCounts" json:"totalCounts"`
}
