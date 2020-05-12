/**
  create by yy on 2020/5/11
*/

package initial

import (
	"fmt"
	"github.com/angui001/ClinkClick/models"
	"sync"
)

type TrackRemark struct {
	Index int
	Mux   *sync.Mutex
}

var TrackRemarkStruct *TrackRemark

func InitTrackRemark() {

	var (
		err error
	)

	// 获取最新的id
	trackRemarkModel := new(models.TrackRemarkModel)

	if err = trackRemarkModel.GetId(); err != nil {
		fmt.Println(err)
	}

	if trackRemarkModel.TrackId == 0 {
		trackRemarkModel.TrackId = 10
	}

	TrackRemarkStruct = &TrackRemark{
		Index: trackRemarkModel.TrackId,
		Mux:   &sync.Mutex{},
	}

	fmt.Println(TrackRemarkStruct)
}
