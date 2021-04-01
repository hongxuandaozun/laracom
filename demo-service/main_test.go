package main

import (
	"fmt"
	pb "github.com/hongxuandaozun/laracom/demo-service/proto/demo"
	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/net/context"
	"testing"
)

func TestDemoServiceHandler_SayHello(t *testing.T) {
	Convey("Given a test for DemoService.SayHello", t, func() {
		context := context.Background()
		service := &DemoServiceHandler{}

		Convey("When send request with specified name", func() {
			var resp pb.DemoResponse
			service.SayHello(context, &pb.DemoRequest{Name: "laodage"}, &resp)

			Convey("Then the response text should be a '你好, laodage'", func() {
				So(resp.Text, ShouldEqual, "你好, laodage")
			})
		})

		Convey("When send request without name", func() {
			var resp pb.DemoResponse
			service.SayHelloByUserId(context, &pb.HelloRequest{Id: "1"}, &resp)

			Convey("Then the response text should be '你好, xxx'", func() {
				fmt.Println(resp.Text)
				So(resp.Text, ShouldEqual, "")
			})
		})
	})
}
