package dao

import (
	"context"
	"log"
	"sync"
	"time"
)

type UserDao struct {
}

var (
	userDao *UserDao
	once    sync.Once
)

// type User struct {
// 	Id        string    `bson:"_id,omitempty"`
// 	Username  string    `bson:"username"`
// 	Password  string    `bson:"password"`
// 	Create_at time.Time `bson:"create_at"`
// 	Delete_at time.Time `bson:"delete_at"`
// }
type User struct {
	Id        int64     `db:"id"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	Create_at time.Time `db:"create_time"`
	Delete_at time.Time `db:"delete_time"`
}

func GetUserDao() *UserDao {
	once.Do(func() {
		userDao = &UserDao{}
	})
	return userDao
}

//自动化生成预编译sql语句

func (u *UserDao) GetUserInfo(userName string) (*User, error) {
	var user *User
	precompile := "select * from  users where username=? and delete_time=''"
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	row := MysqlClient.QueryRowContext(ctx, precompile, user)
	err := row.Scan(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserDao) InsertUserInfo(user *User) error {
	prepare, err := MysqlClient.Prepare("insert into users(username,password,create_time) values(?,?,?)")
	if err != nil {
		return err
	}
	result, err := prepare.Exec(user.Username, user.Password, user.Create_at)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil || affected == 0 {
		return err
	}
	return nil
}

func (u *UserDao) UpdateUserInfo(user *User) error {
	prepare, err := MysqlClient.Prepare("update users set password=?  where username=? and delete_time=''")
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	execContext, err := prepare.ExecContext(ctx, user.Password, user.Username)
	if err != nil {
		return err
	}
	result, err := execContext.RowsAffected()
	if err != nil || result == 0 {
		return err
	}
	return nil
}

func (u *UserDao) DeleteUserDao(user *User) error {
	prepare, err := MysqlClient.Prepare("update users set delete_time=? where username=? and password=? and delete_time=''")
	if err != nil {
		return err
	}
	exec, err := prepare.Exec(time.Now(), user.Username, user.Password)
	if err != nil {
		return err
	}
	affected, err := exec.RowsAffected()
	if err != nil || affected == 0 {
		return err
	}
	return nil
}

func (u *UserDao) DeleteHardUser() {
	queryContext, err := MysqlClient.Query("select * from users where delete_time!=''")
	if err != nil {
		return
	}
	for queryContext.Next() {
		var user *User
		err := queryContext.Scan(user)
		if err != nil {
			log.Panicf("deleteHard error:%s\n", err.Error())
			continue
		}
		if user.Delete_at.Add(time.Hour * 24).Before(time.Now()) {
			prepare, err := MysqlClient.Prepare("delete from users where username=? and password=?")
			if err != nil {
				continue
			}
			execContext, err := prepare.Exec(user.Username, user.Password)
			if err != nil {
				continue
			}
			affected, err := execContext.RowsAffected()
			if err != nil || affected == 0 {
				log.Panicf("hard delete user error:%s \n{username:%s,password:%s,id:%d}\n", err.Error(), user.Username, user.Password, user.Id)
				continue
			}
		}
	}
}

// func (u *UserDao) GetUserInfo(username string) *User {
// 	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	var collection = MongoClient.Database("users").Collection("users")
// 	var user *User
// 	collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
// 	return user
// }

// func (u *UserDao) InsertUserInfo(user *User) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	result, err := MongoClient.Database("users").Collection("user").InsertOne(ctx, &user)
// 	if err != nil || result == nil {
// 		return err
// 	}
// 	return nil
// }

// func (u *UserDao) UpdateUserInfo(user *User) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	MongoClient.Database("users").Collection("users").UpdateOne(ctx, bson.M{{"username", user.Username}, {}}, &user)
// }

// func (u *UserDao) deleteUserInfo(user *User) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	MongoClient.Database("users").Collection("users").UpdateOne(ctx,bson.)
// }
