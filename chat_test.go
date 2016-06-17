package gochatbot

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestChat(t *testing.T) {

	client, _ := NewClient()

	Convey("When chat", t, func() {

		Convey("With empty message", func() {
			res, err := client.Chat("")
			Convey("It should return a error", func() {
				So(err, ShouldNotBeNil)
				So(res, ShouldBeNil)
			})
		})

		Convey("With valid message", func() {
			res, err := client.Chat("test")
			Convey("It should return a valid response", func() {
				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
				So(res.IsSuccess(), ShouldBeTrue)
				So(res.IsEmptyResult(), ShouldBeFalse)
			})
		})
	})

	Convey("When character chat", t, func() {

		Convey("With empty message", func() {
			res, err := client.CharacterChat("", "")
			Convey("It should return a error", func() {
				So(err, ShouldNotBeNil)
				So(res, ShouldBeNil)
			})
		})

		Convey("With empty character type", func() {
			res, err := client.CharacterChat("test", "")
			Convey("It should return a error", func() {
				So(err, ShouldNotBeNil)
				So(res, ShouldBeNil)
			})
		})

		Convey("With valid message and valid character type", func() {
			res, err := client.CharacterChat("test", CharacterTypeCat)
			Convey("It should return a valid response", func() {
				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
				So(res.IsSuccess(), ShouldBeTrue)
				So(res.IsEmptyResult(), ShouldBeFalse)
			})
		})
	})
}
