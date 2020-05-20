/**
  create by yy on 2020/5/11
*/

package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
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

	if _, err = db.Raw("select * from clink_link c " +
		"where c.trans_time>=? and " +
		"c.trans_time<=? and c.trans_type='RENEW' and c.status=1 and  " +
		"c.msisdn in(select msisdn from clink_link d " +
		"where d.trans_type='SUB' and d.status=1 and d.trans_time>=? and d.trans_time<=?);", startTime, endTime, startTime2, endTime2).QueryRows(&list); err != nil {
		// err = libs.NewReportError(err)
	}

	return
}