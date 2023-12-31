package models

type User struct {
	UserId      int `bson:"userId" json:"userId"`
	TotalCounts int `bson:"totalCounts" json:"totalCounts"`
}

func (u *User) Increment() {
	u.TotalCounts++
}

func NewUser(p *CreateUserParams) *User {
	return &User{
		UserId:      p.UserId,
		TotalCounts: 0,
	}
}
