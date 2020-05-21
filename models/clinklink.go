/**
  create by yy on 2020/5/11
*/

package models

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

type ClinkLink struct {
	Id        int     `orm:"pk;auto;column(id)" json:"id"`
	RequestId string  `json:"request_id"`
	Carrier   string  `json:"carrier"`
	ShortCode string  `json:"short_code"`
	Keyword   string  `json:"keyword"`
	Msisdn    string  `json:"msisdn"`
	TransId   string  `json:"trans_id"`
	TransType string  `json:"trans_type"`
	TransTime string  `json:"trans_time"`
	Fee       float64 `json:"fee"`
	Status    int     `json:"status"`
}

func (c *ClinkLink) TableName() string {
	return "clink_link"
}

func (c *ClinkLink) Insert() {
	o := orm.NewOrm()
	_, err := o.Insert(c)
	if err != nil {
		logs.Error("ReqData  Insert  ERROR: ", err.Error())
	}
}

func (c *ClinkLink) GetListSub(startTime, endTime, startTime2, endTime2 string) (list []*ClinkLink, err error) {
	db := orm.NewOrm()

	startTime = changeTime(startTime)
	endTime = changeTime(endTime)

	if _, err = db.Raw("select * from clink_link c "+
		"where c.trans_time>=? and "+
		"c.trans_time<=? and c.trans_type='RENEW' and c.status=1 and  "+
		"c.msisdn in(select msisdn from clink_link d "+
		"where d.trans_type='SUB' and d.status=1 and d.trans_time>=? and d.trans_time<=?);", startTime, endTime, startTime2, endTime2).QueryRows(&list); err != nil {
		// err = libs.NewReportError(err)
	}

	return
}

func changeTime(timeStr string) string {
	var (
		time2  time.Time
		err    error
		result string
	)

	// time1 := "2020-05-12 11:17:19"

	if time2, err = time.Parse("2006-01-02 15:04:05", timeStr); err != nil {
		fmt.Println(err)
	}

	year, mon, day := time2.UTC().Date()
	hour, min, sec := time2.UTC().Clock()
	result = fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d+08:00", year, mon, day, hour, min, sec)

	return result
}
