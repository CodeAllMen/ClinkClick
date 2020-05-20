package main

import (
	"fmt"
	"github.com/MobileCPX/PreBaseLib/databaseutil/redis"
	"github.com/MobileCPX/PreBaseLib/splib/click"
	"github.com/angui001/ClinkClick/initial"
	"github.com/angui001/ClinkClick/models"
	_ "github.com/angui001/ClinkClick/routers"
	"github.com/astaxie/beego"
	"github.com/robfig/cron"
)

func init() {
	redis.Open("127.0.0.1", 6379, "")

	// 初始化服务配置
	models.InitServiceConfig()
	// SendClickDataToAdmin()
	// task.SendMtDaily()

	// 每日发送MT数据
	// sendMtTask()
}

func updateTrackRemark() {
	var (
		err error
	)

	initial.TrackRemarkStruct.Mux.Lock()

	trackModel := new(models.AffTrack)

	if err = trackModel.GetLast(); err != nil {
		fmt.Println(err)
	}

	initial.TrackRemarkStruct.Index = int(trackModel.TrackID)

	initial.TrackRemarkStruct.Mux.Unlock()
}

// 返回一个支持至 秒 级别的 cron
func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}

func startUpdateTrackRemark() {
	var (
		err error
	)

	cr := newWithSeconds()
	if _, err = cr.AddFunc("0 0 0,12 * * ?", updateTrackRemark); err != nil {
		fmt.Println(err)
	}

	cr.Start()
}

func main() {

	initial.InitTrackRemark()

	startUpdateTrackRemark()

	beego.Run()
}

// func sendMtTask() {
//	cr := cron.New()
//	cr.AddFunc("0 7 */1 * * ?", SendClickDataToAdmin) // 一个小时存一次点击数据并且发送到Admin
//	cr.Start()
// }

func SendClickDataToAdmin() {
	models.InsertHourClick()

	for _, service := range models.ServiceData {
		click.SendHourData(service.CampID, click.PROD) // 发送有效点击数据
	}

}
