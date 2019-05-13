package setting

import (
	"log"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	RunMode string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

}
