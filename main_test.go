package main

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(t *testing.T) {

	Convey("With API key from environment variable DO_API_KEY", t, func() {

		data := getServers(os.Getenv("DO_API_KEY"))
		So(data, ShouldNotBeNil)
		servers := parseServers(data)
		So(servers, ShouldNotBeNil)
	})
}
