package handler

import (
	"errors"
	"github.com/hongxuandaozun/laracom/product-service/model"
	pb "github.com/hongxuandaozun/laracom/product-service/proto/product"
	"github.com/hongxuandaozun/laracom/product-service/repo"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
)

type ProductService struct {
	ProductRepo repo.ProductRepositoryInterface
}

func (srv *ProductService) Get(ctx context.Context, req *pb.Product, res *pb.Response) error {
	if req.Id == 0 && req.Slug == "" {
		return errors.New("商品 ID 或 Slug 不能为空")
	}
	var err error
	var productModel *model.Product
	if req.Id != 0 {
		productModel, err = srv.ProductRepo.GetById(uint(req.Id))
	} else {
		productModel, err = srv.ProductRepo.GetBySlug(req.Slug)
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if productModel != nil {
		res.Product, err = productModel.ToProtoBuf()
	}

	return err
}

func (srv *ProductService) Create(ctx context.Context, req *pb.Product, out *pb.Response) error {
	productModel := &model.Product{}
	product, _ := productModel.ToOrm(req)
	if err := srv.ProductRepo.Create(product); err != nil {
		return err
	}
	out.Product, _ = product.ToProtoBuf()
	return nil
}

func (srv *ProductService) Update(ctx context.Context, req *pb.Product, res *pb.Response) error {
	if req.Id == 0 {
		return errors.New("商品 ID 不能为空")
	}
	product, err := srv.ProductRepo.GetById(uint(req.Id))
	if err != nil {
		return err
	}
	product, _ = product.ToOrm(req)
	if err := srv.ProductRepo.Update(product); err != nil {
		return err
	}
	res.Product, _ = product.ToProtoBuf()
	return nil
}
func (srv *ProductService) Delete(ctx context.Context, req *pb.Product, res *pb.Response) error {
	if req.Id == 0 {
		return errors.New("商品 ID 不能为空")
	}
	product, err := srv.ProductRepo.GetById(uint(req.Id))
	if err != nil {
		return err
	}
	if err := srv.ProductRepo.Delete(product); err != nil {
		return err
	}
	res.Product = nil
	return nil
}
func (srv *ProductService) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	products, err := srv.ProductRepo.GetAll()
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	productItems := make([]*pb.Product, len(products))
	for index, product := range products {
		productItem, _ := product.ToProtoBuf()
		productItems[index] = productItem
	}
	res.Products = productItems
	return nil
}
func (srv *ProductService) GetDetail(ctx context.Context, req *pb.Product, res *pb.Response) error {
	if req.Id == 0 {
		return errors.New("商品 ID 不能为空")
	}
	productModel, err := srv.ProductRepo.GetDetailById(uint(req.Id))
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if productModel != nil {
		res.Product, _ = productModel.ToProtoBuf()
	}
	return nil
}
