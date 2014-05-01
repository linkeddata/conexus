package conexus

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lunny/xorm"
)

type User struct {
	Id          int64
	Uri         string    `xorm:"unique varbinary(500) not null"`
	CreatedTime time.Time `xorm:"index created"`
	UpdatedTime time.Time `xorm:"index updated"`
}

type Subscription struct {
	User        int64     `xorm:"index"`
	Source      string    `xorm:"varchar(1000) not null"`
	Destination string    `xorm:"varchar(1000) not null"`
	CreatedTime time.Time `xorm:"index created"`
	UpdatedTime time.Time `xorm:"index updated"`
}

var (
	db *xorm.Engine
)

func init() {
	var err error

	db, err = xorm.NewEngine("mysql", *dsn+"?charset=utf8&parseTime=True&autocommit=true")
	if err != nil {
		log.Panicln(err)
	}
	db.SetMaxIdleConns(4)
	if *debug {
		db.ShowSQL = true
	}
}
