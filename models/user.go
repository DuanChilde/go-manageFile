package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

}

type User struct {
	Id        int64  `json:"id",orm:"pk;auto"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Nickname  string `json:"nickname"`
	Lastlogin int    `json:"lastlogin"`
	Ctime     int64  `json:"ctime"`
	Mtime     int64  `json:"mtime"`
}

func AddUser(u User) int64 {
	u.Lastlogin = 0
	u.Ctime = time.Now().Unix()
	u.Mtime = time.Now().Unix()

	uid, err := db.Insert(&u)
	if err != nil {
		return 0
	}
	return uid
}

func GetUserById(uid int64) (u *User, err error) {
	user := new(User)
	user.Id = uid
	if err := db.Read(user); err != nil {
		return nil, errors.New("User not exists")
	}
	return user, nil
}

func GetAllUsers() map[string]*User {
	return nil
}

func UpdateUser(uid int64, u *User) (a *User, err error) {
	u.Id = uid
	user, err := GetUserById(uid)
	if err != nil {
		return nil, errors.New("User not exists")
	}
	u.Ctime = user.Ctime
	u.Mtime = time.Now().Unix()
	//beego.Debug(user)
	_, err = db.Update(u)
	if err != nil {
		return nil, errors.New("Update User failure")
	}
	return u, nil
}

func Login(username, password string) (uid int64, err error) {
	user := new(User)
	user.Username = username
	//user.Password = password
	beego.Debug(db.Read(user))
	if err := db.Read(user); err != nil {
		return 0, errors.New("User not exists")
	}

	//生成token
	return user.Id, nil
}

func DeleteUser(uid int64) int64 {
	user := new(User)
	user.Id = uid
	num, err := db.Delete(user)
	if err != nil {
		return 0
	}
	return num
}
