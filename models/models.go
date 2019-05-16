// init and connection db

package models

import (
	"fmt"
	"gopkgspider/pkg/setting"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	Mysql_init()
}

func Mysql_init() {
	var (
		err                                  error
		dbType, dbName, user, password, host string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		fmt.Println(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()

	fmt.Println("Models: db ready to connect...")
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		fmt.Println(err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	fmt.Println("Models: db connect success")

}

func CloseDB() {
	defer db.Close()
}
