package main

import (
	"errors"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus"
	db2 "github.com/nonfu/laracom/product-service/db"
	"github.com/nonfu/laracom/product-service/handler"
	"github.com/nonfu/laracom/product-service/model"
	pb "github.com/nonfu/laracom/product-service/proto/product"
	"github.com/nonfu/laracom/product-service/repo"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func prometheusBoot() {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		err := http.ListenAndServe(":9092", nil)
		if err != nil {
			log.Fatal("listenAndServer:", err)
		}
	}()
}
func main() {
	db, err := db2.CreateConnection()
	if err != nil {
		log.Fatalf("Could not connect to DB:%v", err)
		errors.New("链接数据失败")
	}
	defer db.Close()
	// 设置编码格式
	db.Set("gorm:table_options", "charset=utf8")
	// 数据库迁移（商品、图片、品牌、类目、属性相关数据表）

	db.AutoMigrate(
		&model.Product{},
		&model.Attribute{},
		&model.Brand{},
		&model.ProductImage{},
		&model.Category{},
		&model.ProductAttribute{},
		&model.AttributeValue{},
	)

	// 初始化 Repo 实例用于后续数据库操作

	productRepo := &repo.ProductRepository{db}
	imageRepo := &repo.ImageRepository{db}
	brandRepo := &repo.BrandRepository{db}
	categoryRepo := &repo.CategoryRepository{db}
	attributeRepo := &repo.AttributeRepository{db}

	srv := micro.NewService(
		micro.Name("laracom.service.product"),
		micro.Version("latest"),
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
	)
	prometheusBoot()
	srv.Init()
	// 注册处理器
	pb.RegisterProductServiceHandler(srv.Server(), &handler.ProductService{productRepo})
	pb.RegisterImageServiceHandler(srv.Server(), &handler.ImageService{imageRepo})
	pb.RegisterBrandServiceHandler(srv.Server(), &handler.BrandService{brandRepo})
	pb.RegisterCategoryServiceHandler(srv.Server(), &handler.CategoryService{categoryRepo})
	pb.RegisterAttributeServiceHandler(srv.Server(), &handler.AttributeService{attributeRepo})

	err = srv.Run()
	if err != nil {
		fmt.Printf("出错误了哦:%v", err)
	}
}
