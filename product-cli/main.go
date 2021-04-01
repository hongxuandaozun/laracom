package main

import (
	"fmt"
	pb "github.com/hongxuandaozun/laracom/product-service/proto/product"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
	"log"
)

func main() {

	service := micro.NewService()
	client := pb.NewProductServiceClient("laracom.service.product", service.Client())

	service.Init(micro.Action(func(c *cli.Context) {
		//res, err := client.Create(context.TODO(), &pb.Product{
		//	BrandId:     2,
		//	Sku:         "test.sku",
		//	Name:        "test product",
		//	Description: "test description",
		//	Price:       99,
		//	SalePrice:   99,
		//	Status:      1,
		//})

		res, err := client.GetDetail(context.TODO(), &pb.Product{Id: 1})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res.Product)
	}))
	service.Run()
}
