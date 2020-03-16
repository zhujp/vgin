package model

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/zhujp/vgin/app/util/setting"
)

var Db *gorm.DB

var PerPage = 10

type Model struct {
	ID        int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Page struct {
	Page    int `json:"page"`
	PerPage int `json:"per-page"`
}

func (model *Model) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("created_at", time.Now().Format("2006-01-02 15:04:05"))
	scope.SetColumn("updated_at", time.Now().Format("2006-01-02 15:04:05"))
	return nil
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	Db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(Db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	// Db.SingularTable(true) //禁用表名复数
	Db.LogMode(true)
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer Db.Close()
}
