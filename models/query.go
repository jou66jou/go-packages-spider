// mysql db operating
package models

import (
	"container/list"
	"fmt"
)

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

func FindPkgNameAndOV(name string) (gopkg Gopkg, overview string) {
	list := list.New()
	db.Where("pkg_name = ?", name).First(&gopkg)
	list.PushBack(gopkg.PkgName)
	temp := Gopkg{PkgName: gopkg.Refer}
	for i := 1; i < gopkg.Tier; i++ {
		db.Where("pkg_name = ?", temp.PkgName).First(&temp)
		list.PushBack(temp.PkgName)
		temp = Gopkg{PkgName: temp.Refer}
	}
	gopkg.PkgName = ""
	for e := list.Back(); e != nil; e = e.Prev() {
		gopkg.PkgName += e.Value.(string) + "/"
	}
	gopkg.PkgName = gopkg.PkgName[:len(gopkg.PkgName)-1]
	ov := Gopkg_ov{Pkgid: gopkg.Pkgid}
	db.Where("pkgid = ?", ov.Pkgid).First(&ov)
	overview = ov.Overview
	fmt.Printf("%+v\n", gopkg)

	return
}
