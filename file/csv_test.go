package file

import (
	"testing"
)

type OrderInfo struct {
	ID          int64   `csv:"id"`
	AppID       int64   `csv:"app_id"`
	Mid         int64   `csv:"mid"`
	PayMid      int64   `csv:"pay_mid"`
	OrderNo     string  `csv:"order_no"`
	GoodsID     int64   `csv:"goods_id"`
	GoodsNum    float64 `csv:"goods_num"`
	ElecNum     float64 `csv:"elec_num"`
	CalElecNum  float64 `csv:"cal_elec_num"`
	ChannelType string  `csv:"channel_type"`
	Rate        float64 `csv:"rate"`
	Status      int64   `csv:"status"`
	PayWay      int64   `csv:"pay_way"`
	Ctime       string  `csv:"ctime"`
	Mtime       string  `csv:"mtime"`
	LogDate     string  `csv:"log_date"`
}

func TestParseCsv(t *testing.T) {
	fName := "../../common/testdata/elec_order.csv"
	dist := make([]*OrderInfo, 0, 0)

	err := ParseCsv(fName, &dist)
	t.Log("size=", len(dist), err)
	for _, item := range dist {
		t.Logf("%10d, %10d, %10.2f, \t%v, %v\n", item.Mid, item.PayMid, item.CalElecNum, item.Ctime, item.Mtime)
	}
}
