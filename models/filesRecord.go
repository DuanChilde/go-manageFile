package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func init() {}

type filesRecord struct {
	Id          int64  `json:"id",orm:"pk;auto"`
	Type      int8  `json:"type"`
	Name      string  `json:"name"`
	Extension string `json:"extension"`
	Ctime       int64  `json:"ctime"`
	Mtime       int64  `json:"mtime"`
}

func AddFile(f filesRecord) int64 {
	f.Ctime = time.Now().Unix()
	f.Mtime = time.Now().Unix()
	//beego.Debug(s)
	fid, err := db.Insert(&f)
	if err != nil {
		beego.Debug(err)
		return 0
	}
	return fid
}

func GetFileById(id int64) (f *filesRecord, err error) {
	fileRecord := new(filesRecord)
	fileRecord.Id = id
	if err := db.Read(fileRecord); err != nil {
		return nil, errors.New("record not exists")
	}
	return fileRecord, nil
}

func UpdateFile(id int64, f *filesRecord) (a *filesRecord, err error) {
	f.Id = id
	fileRecord, err := GetFileById(id)
	if err != nil {
		return nil, errors.New("file not exists")
	}
	f.Ctime = fileRecord.Ctime
	f.Mtime = time.Now().Unix()
	_, err = db.Update(f)
	if err != nil {
		return nil, errors.New("Update failure")
	}
	return f, nil
}

func GetAllFiles() []*filesRecord {
	var list []*filesRecord
	db.Raw("select * from files_Record").QueryRows(&list)
	return list
}

func DeleteFile(id int64) int64 {
	record := new(filesRecord)
	record.Id = id
	num, err := db.Delete(record)
	if err != nil {
		return 0
	}
	return num
}
