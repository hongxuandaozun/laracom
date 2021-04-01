package main

import (
	"context"
	"fmt"
	pb "github.com/hongxuandaozun/laracom/user-service/proto/user"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"log"
	"os"
)

func main() {
	// 初始化客户端服务，定义命令行参数标识
	service := micro.NewService(
		micro.Flags(
			cli.StringFlag{
				Name:  "name",
				Usage: "Your Name",
			},
			cli.StringFlag{
				Name:  "email",
				Usage: "Your Email",
			},
			cli.StringFlag{
				Name:  "password",
				Usage: "Your Password",
			},
			cli.StringFlag{
				Name:  "id",
				Usage: "Your Id",
			},
		),
	)
	// 远程服务客户端调用句柄
	client := pb.NewUserServiceClient("laracom.service.user", service.Client())
	action := func(c *cli.Context) {

		name := c.String("name")
		email := c.String("email")
		password := c.String("password")
		id := c.String("id")

		log.Println("参数:", name, email, password, id)
		var err error
		res, err := client.Get(context.TODO(), &pb.User{
			Id: id,
		})
		fmt.Println(res)
		os.Exit(0)
		var token *pb.Token
		token, err = client.Auth(context.TODO(), &pb.User{
			Email:    email,
			Name:     name,
			Password: password,
		})
		if err != nil {
			log.Printf("用户登录失败%v", err)
			// 调用用户服务
			r, err := client.Create(context.TODO(), &pb.User{
				Name:     name,
				Email:    email,
				Password: password,
			})
			if err != nil {
				log.Fatalf("user create error :%v", err)
			}
			log.Fatalf("user create success ", r.User.Id)
		}
		log.Printf("用户登录成功,token:%v", token)
		// 调用用户验证服务
		token, err = client.ValidateToken(context.TODO(), token)
		if err != nil {
			log.Fatalf("用户认证失败%v", err)
		}
		log.Printf("用户认证成功")
		// 调用获取所有用户服务
		getAll, err := client.GetAll(context.Background(), &pb.Request{})
		if err != nil {
			log.Fatalf("获取所有用户失败: %v", err)
		}
		for _, v := range getAll.Users {
			log.Println(v)
		}
		fmt.Println("执行完毕")
		os.Exit(0)
	}
	action = func(c *cli.Context) {

		name := c.String("name")
		email := c.String("email")
		password := c.String("password")
		//id := c.String("id")
		token, err := client.Auth(context.TODO(), &pb.User{
			Email:    email,
			Name:     name,
			Password: password,
		})
		if err != nil {
			log.Fatalf(err.Error())
		}
		fmt.Println(token)
	}
	service.Init(

		micro.Action(action),
	)
	service.Run()
}
