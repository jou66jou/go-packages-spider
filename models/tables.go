package models

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
