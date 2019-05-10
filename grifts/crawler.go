package grifts

import (
	"encoding/json"
	"github.com/markbates/grift/grift"
	"github.com/masato25/my104job_board/models/dataobj"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

var _ = grift.Namespace("crawler", func() {
	grift.Desc("get", "get data from 104 api")
	grift.Add("get", func(c *grift.Context) (err error) {
		url := "http://www.104.com.tw/i/apis/jobsearch.cfm?cat=2007001000&sltp=S&slmin=50000&slmax=150000&fmt=8&page=1&pgsz=25"
		resp, err := http.Get(url)
		// read all from response body
		bytebody, err := ioutil.ReadAll(resp.Body)
		var b dataobj.Api104Response
		err = json.Unmarshal(bytebody, &b)
		log.Info(b.PAGECOUNT)
		return
	})

})
