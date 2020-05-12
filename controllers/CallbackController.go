/**
  create by yy on 2020/5/11
*/

package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/MobileCPX/PreBaseLib/splib"
	"github.com/MobileCPX/PreBaseLib/splib/common"
	"github.com/MobileCPX/PreBaseLib/splib/mo"
	"github.com/MobileCPX/PreBaseLib/splib/tracking"
	"github.com/angui001/ClinkClick/data_struct"
	"github.com/angui001/ClinkClick/initial"
	"github.com/angui001/ClinkClick/models"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
)

type CallbackController struct {
	BaseController
}

func (c *CallbackController) Notification() {

	var isExist bool

	serviceConf := models.ServiceInfo{} // 服务配置

	moT := new(mo.Mo)
	notificationType := ""

	var moBase = common.MoBase{} // mo 基础数据

	reqData := new(models.ClinkLink)

	data, _ := ioutil.ReadAll(c.Ctx.Request.Body)

	// 打印一下数据以便后期查找和调试
	fmt.Println(string(data))

	err := json.Unmarshal(data, reqData)
	if err != nil {
		logs.Error("BodyToTrack 转为JSON数据失败")
		fmt.Println(err)
		c.Data["json"] = data_struct.FailedReply("Failure")
		c.ServeJSON()
		c.StopRun()
	}

	// 保存数据
	reqData.Insert()

	// 如果是订阅，则进行回传
	if reqData.TransType == "SUB" {

		initial.TrackRemarkStruct.Mux.Lock()

	START:
		track := new(models.AffTrack) // 点击表
		trackId := initial.TrackRemarkStruct.Index
		initial.TrackRemarkStruct.Index++
		trackRemarkModel := new(models.TrackRemarkModel)
		trackRemarkModel.TrackId = initial.TrackRemarkStruct.Index
		trackRemarkModel.Insert()
		track.TrackID = int64(trackId)
		_ = track.GetOne(tracking.ByTrackID)

		if track.ClickID == "" {
			goto START
		}
		initial.TrackRemarkStruct.Mux.Unlock()

		// 首先获取 配置，配置里有对应的 信息
		serviceId := reqData.ShortCode + "-" + reqData.Keyword
		serviceConf, isExist = c.serviceConfig(serviceId)

		if track.TrackID != 0 {
			moBase.Track = track.Track
			moBase.ServiceID = track.ServiceID
			moBase.TrackID = track.TrackID
		} else {
			moBase.ServiceID = serviceId
		}

		if isExist {
			moBase.Operator = serviceConf.Operator
			moBase.Price = serviceConf.Price
			moBase.Country = serviceConf.Country
		}
		moBase.Msisdn = reqData.Msisdn

		// 存入MO数据
		moT, notificationType = splib.InsertMO(moBase, false, true, "ClinkClick")

		fmt.Println("moT: ", moT, "\nnotification type: ", notificationType)
	}

	c.Data["json"] = data_struct.SuccessReply("Success")
	c.ServeJSON()
}

// 获取网盟的click_id，2020-05-12 解决回传问题专用
func (c *CallbackController) GetClickId() {

	// 加锁防止脏数据和错误数据
	initial.TrackRemarkStruct.Mux.Lock()
START:
	track := new(models.AffTrack)
	trackId := initial.TrackRemarkStruct.Index
	initial.TrackRemarkStruct.Index++
	trackRemarkModel := new(models.TrackRemarkModel)
	trackRemarkModel.TrackId = initial.TrackRemarkStruct.Index
	trackRemarkModel.Insert()

	track.TrackID = int64(trackId)
	_ = track.GetOne(tracking.ByTrackID)

	if track.ClickID == "" {
		// 如果click_id不存在，则进行递归
		goto START
	}

	// 解锁 释放操作权
	initial.TrackRemarkStruct.Mux.Unlock()

	c.Data["json"] = data_struct.SuccessReply(track.ClickID)
	c.ServeJSON()
}

func (c *CallbackController) VisitTrackId() {

	// 加锁防止脏数据和错误数据
	initial.TrackRemarkStruct.Mux.Lock()

	c.Data["json"] = data_struct.SuccessReply(fmt.Sprintf("%v", initial.TrackRemarkStruct.Index))

	// 解锁 释放操作权
	initial.TrackRemarkStruct.Mux.Unlock()

	c.ServeJSON()
}
