/**
  create by yy on 2020/5/11
*/

package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// 此表用于存储最后一次的 track_id
type TrackRemarkModel struct {
	Id      int `orm:"pk;auto;column(id)" json:"id"`
	Count   int `orm:"count" json:"count"`
	TrackId int `orm:"track_id" json:"track_id"`
}

func (t *TrackRemarkModel) TableName() string {
	return "track_remark"
}

func (t *TrackRemarkModel) Insert() {
	o := orm.NewOrm()
	_, err := o.Insert(t)
	if err != nil {
		logs.Error("ReqData  Insert  ERROR: ", err.Error())
	}
}

func (t *TrackRemarkModel) GetId() (err error) {
	// 获取最后一条数据即可
	db := orm.NewOrm()

	err = db.QueryTable(t.TableName()).OrderBy("-id").One(t)

	return
}
