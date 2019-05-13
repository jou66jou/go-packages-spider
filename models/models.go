// init and connection db

package models

import (
	"fmt"
	"gopkgspider/pkg/setting"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Gopkg struct {
	Pkgid    int `gorm:"primary_key"`
	Tier     int
	Refer    string
	PkgName  string
	Synopsis string
	Link     string
}

type Gopkg_ov struct {
	PkgOVid  int `gorm:"primary_key"`
	Pkgid    int
	PkgName  string
	Overview string `gorm:"type:text"`
}

func init() {
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
	if !db.HasTable(&Gopkg{}) {
		fmt.Println("Models: not found Gopkgs table in db and create a new table")
	} else {
		db.DropTable(&Gopkg{})
		fmt.Println("Models: delete and create a new Gopkgs table")
	}
	if !db.HasTable(&Gopkg_ov{}) {
		fmt.Println("Models: not found Gopkg_ovs table in db and create a new table")
	} else {
		db.DropTable(&Gopkg_ov{})
		fmt.Println("Models: delete and create a new Gopkg_ovs table")
	}
	db.CreateTable(&Gopkg{})
	db.CreateTable(&Gopkg_ov{})

}

func CloseDB() {
	defer db.Close()
}
