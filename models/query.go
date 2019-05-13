// mysql db operating
package models

import "fmt"

// InsertPkgData insert gopkgs data to dbserver and return error(err value as nil if succsse).
func InsertPkgData(GopkgTable []*Gopkg) (err error) {
	for i := range GopkgTable {
		data := GopkgTable[i]
		db.Create(data)
	}
	return
}

// InsertPkgOV insert gopkgs overview to dbserver and return error(err value as nil if succsse).
func InsertPkgOV(ov *Gopkg_ov) (err error) {
	data := ov
	if dbc := db.Create(data); dbc.Error != nil {
		fmt.Println(dbc.Error)
		fmt.Printf("%+v\n", ov)
	}
	return
}

func WerePkgLinkAndId() (gopkgs []*Gopkg) {
	db.Where("refer <> ?", "").Select("link,pkgid,pkg_name").Find(&gopkgs)
	return
}
