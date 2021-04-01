package model

import (
	pb "github.com/hongxuandaozun/laracom/user-service/proto/user"
	"github.com/jinzhu/gorm"
	"strconv"
)

type User struct {
	Name     string `gorm:"type:varchar(255);column:name"`
	Email    string `gorm:"type:varchar(255);unique_index;column:email"`
	Password string
	gorm.Model
	Status        uint8  `gorm:"default:1;column:status"`
	StripeId      string `gorm:"column:stripe_id"`
	CardBrand     string `gorm:"column:card_brand"`
	CardLastFour  string `gorm:"column:card_last_four"`
	RememberToken string `gorm:"column:remember_token"`
}

func (model *User) ToORM(req *pb.User) (*User, error) {
	if req.Id != "" {
		id, _ := strconv.ParseInt(req.Id, 10, 64)
		model.ID = uint(id)
	}
	if req.Email != "" {
		model.Email = req.Email
	}
	if req.Name != "" {
		model.Name = req.Name
	}
	if req.Password != "" {
		model.Password = req.Password
	}
	if req.Status != "" {
		status, _ := strconv.ParseUint(req.Id, 10, 64)
		model.Status = uint8(status)
	}
	if req.StripeId != "" {
		model.StripeId = req.StripeId
	}
	if req.CardBrand != "" {
		model.CardBrand = req.CardBrand
	}
	if req.CardLastFour != "" {
		model.CardLastFour = req.CardLastFour
	}
	if req.RememberToken != "" {
		model.RememberToken = req.RememberToken
	}
	return model, nil
}
func (model *User) ToProtobuf() (*pb.User, error) {
	var user = &pb.User{}
	user.Id = strconv.FormatUint(uint64(model.ID), 10)
	user.Name = model.Name
	user.Email = model.Email
	user.Status = strconv.FormatUint(uint64(model.Status), 10)
	user.StripeId = model.StripeId
	user.CardBrand = model.CardBrand
	user.CardLastFour = model.CardLastFour
	user.CreatedAt = model.CreatedAt.Format("2006-01-02 15:04:05")
	user.UpdatedAt = model.UpdatedAt.Format("2006-01-02 15:04:05")
	return user, nil
}
