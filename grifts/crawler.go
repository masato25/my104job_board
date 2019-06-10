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

	"github.com/gofrs/uuid"
	"github.com/markbates/grift/grift"
	m "github.com/masato25/my104job_board/models"
	"github.com/masato25/my104job_board/models/dataobj"
	log "github.com/sirupsen/logrus"
)

var timeparseLayout string = "20060102150405"

func setJobInfo(jb m.Job, d dataobj.Api104Record, cat string) m.Job {
	salHight, _ := strconv.Atoi(d.SalMonthHIGH)
	salLow, _ := strconv.Atoi(d.SalMonthLOW)
	parsedTime, _ := time.Parse(timeparseLayout, d.AppearTIME)
	jb.Area = d.AddrNoDESCRIPT
	jb.C = d.C
	jb.CompanyName = d.NAME
	jb.CreatedAt = parsedTime
	jb.J = d.J
	jb.Cat = cat
	jb.JobCat = d.JobcatDESCRIPT
	jb.JobName = d.JOB
	jb.SalHigh = salHight
	jb.SalLow = salLow
	jb.Welfare = d.WELFARE
	jb.DESCRIPTION = d.DESCRIPTION
	jb.OtherDes = d.DESCRIPTION
	jb.Profile = d.PROFILE
	jb.Link = d.LINK
	jb.Manager = d.S2 == "1"
	jb.NeedOnBt = d.S3 == "1"
	return jb
}

func getDataThanInsertToDB(c *grift.Context, cat string) (err error) {
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
		if totalpage == 1 {
			totalpage, _ = strconv.Atoi(b.TOTALPAGE)
		}
		for _, d := range b.DATA {
			jb := m.Job{}
			m.DB.Where("j = ? and c = ?", d.J, d.C).First(&jb)
			if jb.ID == uuid.Nil {
				jb = setJobInfo(jb, d, cat)
			}
			err = m.DB.Save(&jb)
			if err != nil {
				log.Error(err.Error())
			}
		}
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
