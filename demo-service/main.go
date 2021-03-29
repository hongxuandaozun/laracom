package main

import (
	"context"
	"fmt"
	pb "github.com/hongxuandaozun/laracom/demo-service/proto/demo"
	"github.com/hongxuandaozun/laracom/demo-service/trace"
	"github.com/micro/go-micro"
	traceplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"log"
	"os"
)

type DemoServiceHandler struct {
}

func (s *DemoServiceHandler) SayHello(ctx context.Context, req *pb.DemoRequest, rsp *pb.DemoResponse) error {
	fmt.Println(req.Name)
	rsp.Text = "你好, " + req.Name
	return nil
}

func main() {

	// c初始化全局追踪
	tracer, closer, err := trace.NewTracer("laracom.service.demo", os.Getenv("MICRO_TRACE_SERVER"))
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	// 注册服务名必须和 demo.proto 中的 package 声明一致
	service := micro.NewService(
		micro.Name("laracom.demo.service"),
		micro.WrapHandler(traceplugin.NewHandlerWrapper(opentracing.GlobalTracer())),
	)
	service.Init()

	pb.RegisterDemoServiceHandler(service.Server(), &DemoServiceHandler{})
	if err := service.Run(); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
