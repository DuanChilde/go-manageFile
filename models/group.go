package models

import (
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func init() {

}

type Group struct {
	Id      int64 `orm:"pk;auto"`
	UserId  int64
	GroupId int64
	Ctime   int64
	Mtime   int64
}

func AddGroup(g Group) int64 {
	g.GroupId = 0
	g.GroupId = 0
	g.Ctime = time.Now().Unix()
	g.Mtime = time.Now().Unix()

	id, err := db.Insert(&g)
	if err != nil {
		return 0
	}
	return id
}

func GetGroupById(id int64) (g *Group, err error) {
	group := new(Group)
	group.Id = id
	if err := db.Read(group); err != nil {
		return nil, errors.New("Group not exists")
	}
	return group, nil
}

func GetAllGroups() map[string]*Group {
	return nil
}

func UpdateGroup(id int64, g *Group) (a *Group, err error) {
	group, err := GetGroupById(id)
	if err != nil {
		return nil, errors.New("Group not exists")
	}
	g.Id = id
	g.Ctime = group.Ctime
	g.Mtime = time.Now().Unix()
	//beego.Debug(Group)
	_, err = db.Update(g)
	if err != nil {
		return nil, errors.New("Update Group failure")
	}
	return g, nil
}

func DeleteGroup(uid int64) int64 {
	Group := new(Group)
	Group.Id = uid
	num, err := db.Delete(Group)
	if err != nil {
		return 0
	}
	return num
}
