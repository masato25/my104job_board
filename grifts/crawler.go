package grifts

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/markbates/grift/grift"
	m "github.com/masato25/my104job_board/models"
	"github.com/masato25/my104job_board/models/dataobj"
	log "github.com/sirupsen/logrus"
)

func getDataThanInsertToDB(c *grift.Context, cat string) (err error) {
	timeparseLayout := "20060102150405"
	totalpage := 1
	for currentpage := 1; currentpage <= totalpage; currentpage++ {
		url := fmt.Sprintf("http://www.104.com.tw/i/apis/jobsearch.cfm?cat=%s&sltp=S&slmin=50000&slmax=150000&fmt=8&page=%d&pgsz=25", cat, currentpage)
		resp, err := http.Get(url)
		if err != nil {
			log.Error(err.Error())
		}
		// read all from response body
		bytebody, err := ioutil.ReadAll(resp.Body)
		var b dataobj.Api104Response
		err = json.Unmarshal(bytebody, &b)
		totalpage, _ = strconv.Atoi(b.TOTALPAGE)
		jobs := m.Jobs{}
		for _, d := range b.DATA {
			salHight, _ := strconv.Atoi(d.SalMonthHIGH)
			salLow, _ := strconv.Atoi(d.SalMonthLOW)
			parsedTime, _ := time.Parse(timeparseLayout, d.AppearTIME)
			jb := m.Job{
				Area:        d.AddrNoDESCRIPT,
				C:           d.C,
				CompanyName: d.NAME,
				CreatedAt:   parsedTime,
				J:           d.J,
				Cat:         cat,
				JobCat:      d.JobcatDESCRIPT,
				JobName:     d.JOB,
				SalHigh:     salHight,
				SalLow:      salLow,
				Welfare:     d.WELFARE,
				DESCRIPTION: d.DESCRIPTION,
				OtherDes:    d.DESCRIPTION,
				Profile:     d.PROFILE,
				Link:        d.LINK,
				Manager:     d.S2 == "1",
				NeedOnBt:    d.S3 == "1",
			}
			jobs = append(jobs, jb)
		}
		for _, j := range jobs {
			err = m.DB.Create(&j)
			if err != nil {
				log.Error(err.Error())
			}
		}
		jobs = m.Jobs{}
		log.Infof("totallpage: %d currentpage: %d", totalpage, currentpage)
	}
	return
}

var _ = grift.Namespace("crawler", func() {
	grift.Desc("get", "get data from 104 api")
	grift.Add("get", func(c *grift.Context) (err error) {
		file, err := os.Open("./catlist")
		if err != nil {
			return err
		}
		rd := bufio.NewReader(file)
		var readfileerr error
		for readfileerr == nil {
			var dbody []byte
			dbody, _, readfileerr = rd.ReadLine()
			getsets := strings.Split(string(dbody), ",")
			err = getDataThanInsertToDB(c, getsets[0])
		}
		return
	})
})
