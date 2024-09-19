package chinaz

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
		chinaz.config.Token = `***`

		// Convey("TestNewChinaz", func() {
		// 	sign := chinaz.sign("123")
		// 	So(sign, ShouldNotBeNil)
		// 	So(sign["rd"], ShouldEqual, "266")
		// 	So(sign["key"], ShouldEqual, "266,416")
		// })

		Convey("TestQUery", func() {
			res, err := chinaz.query(context.Background(), "baidu.com")
			So(err, ShouldBeNil)
			zap.S().Infof("%+v", res)
		})
	})
}
