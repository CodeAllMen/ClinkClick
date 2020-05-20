/**
  create by yy on 2020/5/20
*/

package controllers

import (
	"fmt"
	"github.com/angui001/ClinkClick/data_struct"
	"github.com/angui001/ClinkClick/models"
	"github.com/astaxie/beego"
)

type CountController struct {
	beego.Controller
}

func (c *CountController) CountSub() {

	var (
		err   error
		list  []*models.ClinkLink
		total = 0.0
		// fee   float64
	)

	startTime := c.GetString("start_time")
	endTime := c.GetString("end_time")

	startTime2 := c.GetString("start_time_2")
	endTime2 := c.GetString("end_time_2")

	// 获取数据
	reqDataModel := new(models.ClinkLink)

	if list, err = reqDataModel.GetListSub(startTime, endTime, startTime2, endTime2); err != nil {
		// err = libs.NewReportError(err)
		fmt.Println(err)
	}

	// 遍历 数据，并且进行累加
	for _, data := range list {
		// 扣费成功的才进行计算
		total = total + data.Fee
	}

	result := "%v  到  %v 的总扣费为：%v, 扣费成功数为：%v"
	result = fmt.Sprintf(result, startTime, endTime, total, len(list))

	c.Data["json"] = data_struct.SuccessReply(result)
	c.ServeJSON()
}
