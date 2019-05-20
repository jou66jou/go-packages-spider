package debug

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func PrintReuest(r *http.Request) {

	requestDump, e := httputil.DumpRequest(r, true)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(string(requestDump))
}
