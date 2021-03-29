package main

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	database "github.com/nonfu/laracom/user-service/db"
	"github.com/nonfu/laracom/user-service/handler"
	"github.com/nonfu/laracom/user-service/model"
	pb "github.com/nonfu/laracom/user-service/proto/user"
	repository "github.com/nonfu/laracom/user-service/repo"
	"github.com/nonfu/laracom/user-service/service"
)

func main() {
	db, err := database.CreateConnection()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	defer db.Close()
	// 每次启动服务时都会检查，如果数据表不存在则创建，已存在检查是否有修改

	db.AutoMigrate(
		&model.User{},
		&model.PasswordReset{},
	)
	// 初始化 Repo 实例用于后续数据库操作

	repo := &repository.UserRepository{db}
	resetRepo := &repository.PasswordResetRepository{db}
	// 以下是 Micro 创建微服务流程
	srv := micro.NewService(
		micro.Name("laracom.service.user"),
		micro.Version("latest"),
	)

	srv.Init()
	token := &service.TokenService{repo}

	// 获取broker实例
	broker := srv.Server().Options().Broker

	// 注册处理器
	pb.RegisterUserServiceHandler(srv.Server(), &handler.UserService{repo, token, resetRepo, broker})

	// 启动用户服务
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
