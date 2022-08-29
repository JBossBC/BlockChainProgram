package dao

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type UserDao struct {
}

var (
	userDao *UserDao
	once    sync.Once
)

type User struct {
	Id        string    `bson:"_id,omitempty"`
	Username  string    `bson:"username"`
	Password  string    `bson:"password"`
	Create_at time.Time `bson:"create_at"`
}

func GetUserDao() *UserDao {
	once.Do(func() {
		userDao = &UserDao{}
	})
	return userDao
}
func (u *UserDao) GetUserInfo(username string) *User {
	var ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	var collection = MongoClient.Database("users").Collection("users")
	var user *User
	collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	return user
}
