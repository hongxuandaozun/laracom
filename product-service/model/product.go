package model

import (
	"github.com/jinzhu/gorm"
	pb "github.com/nonfu/laracom/product-service/proto/product"
)

type Product struct {
	gorm.Model
	BrandId      uint32  `gorm:"unsigned,default:0"`
	Sku          string  `gorm:"type:varchar(255)"`
	Name         string  `gorm:"type:varchar(255)"`
	Slug         string  `gorm:"type:varchar(255)"`
	Description  string  `gorm:"type:text"`
	Cover        string  `gorm:"type:varchar(255)"`
	Quantity     uint32  `gorm:"unsigned,default:0"`
	Price        float32 `gorm:"type:decimal(8,2)"`
	SalePrice    float32 `gorm:"type:decimal(8,2)"`
	Status       uint8   `gorm:"unsigned,default:0"`
	Length       float32 `gorm:"type:decimal(8,2)"`
	Width        float32 `gorm:"type:decimal(8,2)"`
	Height       float32 `gorm:"type:decimal(8,2)"`
	Weight       float32 `gorm:"type:decimal(8,2)"`
	DistanceUnit string  `gorm:"type:varchar(255)"`
	MassUnit     string  `gorm:"type:varchar(255)"`
	Brand        Brand
	Attributes   []*ProductAttribute
	Images       []*ProductImage
	Categories   []*Category `gorm:"many2many:category_product;"`
}

func (model *Product) ToOrm(product *pb.Product) (*Product, error) {
	if product.Id != 0 {
		model.ID = uint(product.Id)
	}
	if product.Name != "" {
		model.Name = product.Name
	}
	if product.BrandId != 0 {
		model.BrandId = product.BrandId
	}
	if product.Sku != "" {
		model.Sku = product.Sku
	}
	if product.Slug != "" {
		model.Slug = product.Slug
	}
	if product.Description != "" {
		model.Description = product.Description
	}
	if product.Cover != "" {
		model.Cover = product.Cover
	}
	if product.Quantity != 0 {
		model.Quantity = product.Quantity
	}
	if product.Price != 0 {
		model.Price = product.Price
	}
	if product.SalePrice != 0 {
		model.SalePrice = product.SalePrice
	}
	if product.Status != 0 {
		model.Status = uint8(product.Status)
	}
	if product.Length != 0 {
		model.Length = product.Length
	}
	if product.Width != 0 {
		model.Width = product.Width
	}
	if product.Height != 0 {
		model.Height = product.Height
	}
	if product.Weight != 0 {
		model.Weight = product.Weight
	}
	if product.DistanceUnit != "" {
		model.DistanceUnit = product.DistanceUnit
	}
	if product.MassUnit != "" {
		model.MassUnit = product.MassUnit
	}
	return model, nil
}
func (model *Product) ToProtoBuf() (*pb.Product, error) {
	var product = &pb.Product{}
	product.Id = uint32(model.ID)
	product.Status = uint32(model.Status)
	product.CreatedAt = model.CreatedAt.Format("2006-01-02 15:04:05")
	product.UpdatedAt = model.UpdatedAt.Format("2006-01-02 15:04:05")
	product.BrandId = model.BrandId
	product.Sku = model.Sku
	product.Name = model.Name
	product.Slug = model.Slug
	product.Description = model.Description
	product.Cover = model.Cover
	product.Quantity = model.Quantity
	product.Price = model.Price
	product.SalePrice = model.SalePrice
	product.Length = model.Length
	product.Width = model.Width
	product.Height = model.Height
	product.Weight = model.Weight
	product.DistanceUnit = model.DistanceUnit
	product.MassUnit = model.MassUnit
	if model.Images != nil {
		images := make([]*pb.ProductImage, len(model.Images))
		for index, value := range model.Images {
			images[index], _ = value.ToProtobuf()
		}
		product.Images = images
	}
	if model.Brand.ID != 0 {
		product.Brand, _ = model.Brand.ToProtobuf()
	}
	if model.Attributes != nil {
		//attributes := make([]*pb.Attribute,len(model.Attributes))
		for index, value := range model.Attributes {
			product.Attributes[index], _ = value.ToProtobuf()
		}

	}
	if model.Categories != nil {
		c := make([]*pb.Category, len(model.Categories))
		for index, value := range model.Categories {
			c[index], _ = value.ToProtobuf()
		}
		product.Categories = c
	}
	return product, nil
}
