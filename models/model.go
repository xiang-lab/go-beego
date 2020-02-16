package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id 		int
	Name 	string
	Pwd		string
}

type Article struct {
	Id			int
	Aname		string
	Atime		time.Time
	Acount		int
	Acontent	string
	Aimage		string
}

func init() {
	/* 涉资数据库基本信息 */
	orm.RegisterDataBase("default", "mysql",
		"root:zhangxiang133@tcp(127.0.0.1:3306)/test?charset=utf8")
	/* 映射model数据 */
	orm.RegisterModel(new(User), new(Article))
	/* 生成流表 */
	orm.RunSyncdb("default", false, true)
}
