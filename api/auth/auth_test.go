package auth

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestVerify(t *testing.T) {
	secret = "test123"

	token, _ := Sign(int64(8))

	userID, err := Verify(token)
	Convey("Test Verify", t, func() {
		So(err, ShouldBeNil)
		So(userID, ShouldEqual, 8)
	})
}
