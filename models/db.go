package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db orm.Ormer
)
var (
	dbhost     string = "localhost:3306" //主机，米有用到( ⊙ o ⊙ )！
	dbuser     string = "root"           //数据库用户名
	dbpassword string = ""               //数据库密码
	dbname     string = "store"          //数据库名字
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	conn := dbuser + ":" + dbpassword + "@/" + dbname + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", conn)
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Group))
	db = orm.NewOrm()   //注册新的orm
	db.Using("dafault") //使用数据库，默认default
}
