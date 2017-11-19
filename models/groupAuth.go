package models

import (
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func init() {

}

type GroupAuth struct {
	Id      int64 `orm:"pk;auto"`
	UserId  int64
	GroupId int64
	Ctime   int64
	Mtime   int64
}

func AddGroupAuth(g GroupAuth) int64 {
	g.UserId = 0
	g.GroupId = 0
	g.Ctime = time.Now().Unix()
	g.Mtime = time.Now().Unix()

	id, err := db.Insert(&g)
	if err != nil {
		return 0
	}
	return id
}

func GetGroupAuthById(id int64) (g *GroupAuth, err error) {
	groupAuth := new(GroupAuth)
	groupAuth.Id = id
	if err := db.Read(groupAuth); err != nil {
		return nil, errors.New("Group not exists")
	}
	return groupAuth, nil
}

func GetAllGroups() []*GroupAuth {
	var groupAuthList []*GroupAuth
	db.Raw("select * from groupAuth").QueryRows(&groupAuthList)
	return groupAuthList
}

func UpdateGroup(id int64, g *GroupAuth) (a *GroupAuth, err error) {
	groupAuth, err := GetGroupAuthById(id)
	if err != nil {
		return nil, errors.New("Group not exists")
	}
	g.Id = id
	g.Ctime = groupAuth.Ctime
	g.Mtime = time.Now().Unix()
	_, err = db.Update(g)
	if err != nil {
		return nil, errors.New("Update GroupAuth failure")
	}
	return g, nil
}

func DeleteGroupAuth(uid int64) int64 {
	groupAuth := new(GroupAuth)
	groupAuth.Id = uid
	num, err := db.Delete(groupAuth)
	if err != nil {
		return 0
	}
	return num
}
