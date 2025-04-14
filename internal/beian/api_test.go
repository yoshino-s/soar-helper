package beian

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

func TestApi(t *testing.T) {
	zap.ReplaceGlobals(zaptest.NewLogger(t))
	Convey("TestApi", t, func() {
		chinaz := New()
		chinaz.config.WerplusKey = "prKer1HQyPWAa5Vc4ZdtbeuAhb"

		Convey("TestWerplusQuery", func() {
			res, err := chinaz.werplusQuery(context.Background(), "ahapp.net")
			So(err, ShouldBeNil)
			zap.S().Infof("%+v", res)
		})
	})
}

// "domain": "baidu.com",
// "domainId": 10000245113,
// "leaderName": "",
// "limitAccess": "否",
// "mainId": 282751,
// "mainLicence": "京ICP证030173号",
// "natureName": "企业",
// "serviceId": 282911,
// "serviceLicence": "京ICP证030173号-1",
// "unitName": "北京百度网讯科技有限公司",
// "updateRecordTime": "2019-05-16 16:06:21"
