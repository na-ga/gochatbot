package gochatbot

import (
	"os"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestChatbotClient(t *testing.T) {

	Convey("When creating a client", t, func() {

		Convey("Without no environment variables", func() {
			os.Setenv("CHATBOT_ENDPOINT_URL", "")
			os.Setenv("CHATBOT_API_KEY", "")
			client, err := NewClient()
			Convey("It should return a valid client", func() {
				So(client, ShouldNotBeNil)
				So(err, ShouldBeNil)
				So(client.URL, ShouldEqual, defaultEndPoint)
				So(client.APIKey, ShouldEqual, defaultAPIKey)
				So(client.TimeOut.Seconds(), ShouldEqual, 1)
			})
		})

		Convey("With environment variables", func() {
			os.Setenv("CHATBOT_ENDPOINT_URL", "http://localhost:8080/api")
			os.Setenv("CHATBOT_API_KEY", "develop")
			client, err := NewClient()
			Convey("It should return a valid client", func() {
				So(client, ShouldNotBeNil)
				So(err, ShouldBeNil)
				So(client.URL, ShouldEqual, "http://localhost:8080/api")
				So(client.APIKey, ShouldEqual, "develop")
				So(client.TimeOut.Seconds(), ShouldEqual, 1)
			})
		})
	})

	Convey("When creating a client with configuration", t, func() {

		Convey("With empty configuration", func() {
			client, err := NewClientWithConfig(ChatbotConfig{})
			Convey("It should return a valid client", func() {
				So(client, ShouldNotBeNil)
				So(err, ShouldBeNil)
				So(client.URL, ShouldEqual, defaultEndPoint)
				So(client.APIKey, ShouldEqual, defaultAPIKey)
				So(client.TimeOut.Seconds(), ShouldEqual, 1)
			})
		})

		Convey("With valid configuration", func() {
			client, err := NewClientWithConfig(ChatbotConfig{
				URL:     "http://localhost:8080/api",
				APIKey:  "develop",
				TimeOut: 2 * time.Second,
			})
			Convey("It should return a valid client", func() {
				So(client, ShouldNotBeNil)
				So(err, ShouldBeNil)
				So(client.URL, ShouldEqual, "http://localhost:8080/api")
				So(client.APIKey, ShouldEqual, "develop")
				So(client.TimeOut.Seconds(), ShouldEqual, 2)
			})
		})
	})
}
