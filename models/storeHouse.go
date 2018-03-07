package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func init() {}

type StoreHouse struct {
	Id          int64  `json:"id",orm:"pk;auto"`
	UserId      int64  `json:"userId"`
	FileId      int64  `json:"fieldId"`
	Description string `json:"description"`
	Ctime       int64  `json:"ctime"`
	Mtime       int64  `json:"mtime"`
}

func AddStore(s StoreHouse) int64 {

	s.Ctime = time.Now().Unix()
	s.Mtime = time.Now().Unix()
	beego.Debug(s)
	sid, err := db.Insert(&s)
	if err != nil {
		beego.Debug(err)
		return 0
	}
	return sid
}

func GetStoreById(id int64) (s *StoreHouse, err error) {
	store := new(StoreHouse)
	store.Id = id
	if err := db.Read(store); err != nil {
		return nil, errors.New("record not exists")
	}
	return store, nil
}

func UpdateStore(id int64, s *StoreHouse) (a *StoreHouse, err error) {
	s.Id = id
	store, err := GetStoreById(id)
	if err != nil {
		return nil, errors.New("store not exists")
	}
	s.Ctime = store.Ctime
	s.Mtime = time.Now().Unix()
	_, err = db.Update(s)
	if err != nil {
		return nil, errors.New("Update failure")
	}
	return s, nil
}

func GetAllStore() []*StoreHouse {
	var storeList []*StoreHouse
	db.Raw("select * from StoreHouse").QueryRows(&storeList)
	return storeList
}

func DeleteStore(id int64) int64 {
	record := new(StoreHouse)
	record.Id = id
	num, err := db.Delete(record)
	if err != nil {
		return 0
	}
	return num
}
