package main

import (
	"context"
	"github.com/hongxuan/laracom/demo-service"
	pb "github.com/hongxuan/laracom/demo-service/proto/demo"
	"github.com/micro/go-micro"
	"log"
)

func main() {

	service := micro.NewService(micro.Name("laracom.demo.cli"))
	service.Init()

	client := pb.NewDemoServiceClient("laracom.demo.service", service.Client())
	rsp, err := client.SayHello(context.TODO(), &pb.DemoRequest{Name: "学院君"})
	if err != nil {
		log.Fatalf("服务调用失败：%v", err)
		return
	}
	log.Println(rsp.Text)
}
