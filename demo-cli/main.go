package main

import (
	"context"
	pb "github.com/hongxuandaozun/laracom/demo-service/proto/demo"
	"github.com/hongxuandaozun/laracom/demo-service/trace"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	traceplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"log"
	"os"
)

func main() {
	tracer, closer, err := trace.NewTracer("laracom.demo.cli", os.Getenv("MICRO_TRACE_SERVER"))
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()

	service := micro.NewService(micro.Name("laracom.demo.cli"), micro.WrapClient(traceplugin.NewClientWrapper(tracer)))

	service.Init()

	client := pb.NewDemoServiceClient("laracom.demo.service", service.Client())

	// 创建空的上下文,生成追踪span
	span, ctx := opentracing.StartSpanFromContext(context.Background(), "call")
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(map[string]string)
	}
	defer span.Finish()

	// 注入opentracing textmap 到空的上下文用于追踪
	opentracing.GlobalTracer().Inject(span.Context(), opentracing.TextMap, opentracing.TextMapCarrier(md))
	metadata.NewContext(opentracing.ContextWithSpan(ctx, span), md)
	rsp, err := client.SayHello(ctx, &pb.DemoRequest{Name: "学院君"})
	if err != nil {
		log.Fatalf("服务调用失败：%v", err)
		return
	}
	log.Println(rsp.Text)
}
