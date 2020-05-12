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
