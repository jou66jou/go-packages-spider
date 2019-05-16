//	spider for golang pkgs library

package models

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func SpiderGopkgs(url string) {
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
	data, getE := GetGopkgsList(url) // get golang pkgs standard lib description list
	if getE != nil {
		fmt.Println(getE)
		return
	}
	fmt.Println("Main: get gopkgslist success")

	insertDataE := InsertPkgData(data) // insert pkgs list to db
	if insertDataE != nil {
		fmt.Println(insertDataE)
		return
	}
	fmt.Println("Main: gopkgslist insert to db success")

	insertOverViewE := GetOVAndInsert() // routine get pkgs overview and insert to db

	if insertOverViewE != nil {
		fmt.Println(getE)
		return
	}
	fmt.Println("Main: pkgs overview insert to db success")
}

// GetGopkgsList parses a url and returns golang pkgs elements against
// a Gopkg Object (ex: tier,reffer,pkgNme ... some pkg's description and relationship).
func GetGopkgsList(url string) (GopkgTable []*Gopkg, getE error) {

	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	} else {
		doc.Find("div.pkg-dir").Find(".pkg-name,.pkg-synopsis").Each(func(i int, selection *goquery.Selection) {
			sHtml, _ := selection.Html()
			sHtml = (strings.TrimSpace(sHtml))
			r, rE := regexp.Compile("\"([^\"]*)\"")
			if rE != nil {
				getE = rE
			} else if (i - 1) > 0 { // ignore first and seceond selection
				data := new(Gopkg)
				// check pkgs-name(even) or pkgs-synopsis(odd)
				if i%2 == 0 {
					//a pkgs-name html like: <a href="archive/">archive</a>
					regHtml := strings.Replace(r.FindString(sHtml), "\"", "", -1)
					data.Tier = strings.Count(regHtml, "/")
					if data.Tier > 0 {
						var temp = strings.Split(regHtml, "/")
						if data.Tier > 1 {
							data.Refer = temp[len(temp)-3]
						}
						data.PkgName = temp[len(temp)-2]
						data.Link = url + regHtml
						GopkgTable = append(GopkgTable, data)
					}
				} else { // insert pkgs-synopsis content to last GopkgTable index
					GopkgTable[len(GopkgTable)-1].Synopsis = sHtml
				}

			}

		})
		if getE != nil {
			return nil, getE
		}
		return GopkgTable, nil
	}
}

// GetOVAndInsert routine query function gets overview content using pkgs link.
// And insert a Gopkg_ov object to db.
func GetOVAndInsert() (getE error) {
	gopkgs := WerePkgLinkAndId()
	gopkg_ovChan := make(chan *Gopkg_ov, len(gopkgs))
	defer close(gopkg_ovChan)
	for _, gopkg := range gopkgs {
		go GetOverView(gopkg, gopkg_ovChan)
	}
	for i := 0; i < len(gopkgs); i++ {
		ov := <-gopkg_ovChan
		InsertPkgOV(ov)
	}
	return
}

// GetOverView a rotuine function retrun Gopkg_ov object.
func GetOverView(gopkg *Gopkg, OV_Chan chan *Gopkg_ov) {
	var content string
	doc, _ := goquery.NewDocument(gopkg.Link)
	doc.Find(".toggleButton~p").Each(func(i int, selection *goquery.Selection) {
		sHtml, sE := selection.Html()
		if sE != nil {
			fmt.Println("spider: pkgs overview goquery not found")
		} else {
			if strings.Index(sHtml, "span") == -1 {
				content += strings.TrimSpace(string(selection.Text()))
			}
		}
	})
	OV_Chan <- &Gopkg_ov{Pkgid: gopkg.Pkgid, PkgName: gopkg.PkgName, Overview: content}
}
