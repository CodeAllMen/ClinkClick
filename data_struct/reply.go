/**
  create by yy on 2020/5/11
*/

package data_struct

type Reply struct {
	Status int    `json:"status"`
	Desc   string `json:"desc"`
}

func SuccessReply(desc string) *Reply {
	return &Reply{
		Status: 1,
		Desc:   desc,
	}
}

func FailedReply(desc string) *Reply {
	return &Reply{
		Status: 0,
		Desc:   desc,
	}
}
