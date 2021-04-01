package main

import (
	"context"
	tracer "github.com/hongxuandaozun/laracom/common/tracer"
	"github.com/hongxuandaozun/laracom/common/wrapper/breaker/hystrix"
	pb "github.com/hongxuandaozun/laracom/demo-service/proto/demo"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/v2"
	traceplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"log"
	"os"
)

func main() {

	// 初始化追踪器
	t, io, err := tracer.NewTracer("laracom.demo.cli", os.Getenv("MICRO_TRACE_SERVER"))
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()

	hystrix.Configure([]string{"laracom.service.demo.DemoService.SayHello"})

	service := micro.NewService(
		micro.Name("laracom.demo.cli"),
		micro.WrapClient(traceplugin.NewClientWrapper(t), hystrix.NewClientWrapper()),
	)
	service.Init()

	client := pb.NewDemoServiceClient("laracom.service.demo", service.Client())

	// 创建空的上下文, 生成追踪 span
	span, ctx := opentracing.StartSpanFromContext(context.Background(), "call")
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(map[string]string)
	}
	defer span.Finish()

	// 注入 opentracing textmap 到空的上下文用于追踪
	opentracing.GlobalTracer().Inject(span.Context(), opentracing.TextMap, opentracing.TextMapCarrier(md))
	ctx = opentracing.ContextWithSpan(ctx, span)
	ctx = metadata.NewContext(ctx, md)
	// 记录请求 && 响应 && 错误
	req := &pb.DemoRequest{Name: "学院君22222"}
	span.SetTag("req", req)
	resp, err := client.SayHello(ctx, req)
	if err != nil {
		span.SetTag("err", err)
		log.Fatalf("服务调用失败：%v", err)
		return
	}
	span.SetTag("resp", resp)
	log.Println(resp.Text)
}
