package chinaz

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

func TestApi(t *testing.T) {
	zap.ReplaceGlobals(zaptest.NewLogger(t))
	Convey("TestApi", t, func() {
		chinaz := New()
		//chinaz.config.Cookie = `toolUserGrade=DA558BECA59696EB6D6F7073658259097496A34F9B3E8B35F3075E72A88B4A26B9173A6415D4CA331C61DBB7B0154FF267C69A5E3189BA1373D01C9EB138F2DDC90982EE95764B6E26452A187D4C6FBB03FE33FBDE50D3500A6703ED03B4A73340E2F01F9304A7FBB03D156FBABC554AA8E0653046362855; bbsmax_user=1`
		chinaz.config.Token = `apiuser_quantity_1b3ce90389f3e35a9277a743bf7b56de_3d0732b434854b5cbaafae5a6eb999c5`

		// Convey("TestNewChinaz", func() {
		// 	sign := chinaz.sign("123")
		// 	So(sign, ShouldNotBeNil)
		// 	So(sign["rd"], ShouldEqual, "266")
		// 	So(sign["key"], ShouldEqual, "266,416")
		// })

		Convey("TestQUery", func() {
			res, err := chinaz.query("baidu.com")
			So(err, ShouldBeNil)
			zap.S().Infof("%+v", res)
		})
	})
}
