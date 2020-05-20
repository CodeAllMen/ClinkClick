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
	Index   int
	Total   int
	Mux     *sync.Mutex
}

var TrackRemarkStruct *TrackRemark

func InitTrackRemark() {

	var (
		err error
	)

	// 获取 当天 最新的id
	trackModel := new(models.AffTrack)
	trackRemarkModel := new(models.TrackRemarkModel)

	if err = trackRemarkModel.GetId(); err != nil {
		fmt.Println(err)
	}

	if err = trackModel.GetLast(); err != nil {
		fmt.Println(err)
	}

	TrackRemarkStruct = &TrackRemark{
		Index: int(trackModel.TrackID),
		Total: trackRemarkModel.Count,
		Mux:   &sync.Mutex{},
	}

	fmt.Println(TrackRemarkStruct)
}
