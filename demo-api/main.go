package main

import (
	"context"
	"github.com/gin-gonic/gin"
	demo "github.com/hongxuandaozun/laracom/demo-service/proto/demo"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/web"
	"log"
	"net/http"
)

type Say struct {
}

var (
	cli demo.DemoServiceClient
)

func (s *Say) Anything(c *gin.Context) {
	log.Println("Received Say.Anything API request")
	c.JSON(http.StatusOK, gin.H{
		"message": "鸿玄道尊  666",
	})
}

func (s *Say) Hello(c *gin.Context) {
	log.Println("Received Say.Hello API request")
	name := c.Param("name")
	res, err := cli.SayHello(context.TODO(), &demo.DemoRequest{
		Name: name,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, res)
}

func main() {
	service := web.NewService(web.Name("laracom.demo.api"))
	service.Init()
	cli = demo.NewDemoServiceClient("laracom.demo.service", client.DefaultClient)
	say := new(Say)
	router := gin.Default()
	router.GET("/hello", say.Anything)
	router.GET("hello/:name", say.Hello)
	service.Handle("/", router)

	service.Run()
}
