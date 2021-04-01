package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/errors"
	"github.com/nonfu/laracom/user-service/model"
	pb "github.com/nonfu/laracom/user-service/proto/user"
	"github.com/nonfu/laracom/user-service/repo"
	"github.com/nonfu/laracom/user-service/service"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
)

type UserService struct {
	Repo      repo.Repository
	Token     service.Authable
	ResetRepo repo.PasswordResetInterface
	PubSub    broker.Broker
}

func (srv *UserService) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	var userModel *model.User
	var err error
	if req.Id != "" {
		id, _ := strconv.ParseInt(req.Id, 10, 64)
		userModel, err = srv.Repo.Get(uint(id))
	} else if req.Email != "" {
		userModel, err = srv.Repo.GetByEmail(req.Email)
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if userModel != nil {
		res.User, _ = userModel.ToProtobuf()
	}
	return nil
}
func (srv *UserService) Create(ctx context.Context, req *pb.User, res *pb.Response) error {

	// 对密码进行hash加密
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hashedPass)
	var userModel = &model.User{}
	userModel.ToORM(req)
	if err = srv.Repo.Create(userModel); err != nil {
		return err
	}
	res.User, _ = userModel.ToProtobuf()
	return nil
}

func (srv *UserService) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	users, err := srv.Repo.GetAll()
	if err != nil {
		return err
	}
	userItems := make([]*pb.User, len(users))
	for index, user := range users {
		userItems[index], _ = user.ToProtobuf()
	}
	res.Users = userItems
	return nil
}
func (srv *UserService) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	log.Println("Logging in with :", req.Email, req.Name)
	// 获取用户信息
	userModel, err := srv.Repo.GetByEmail(req.Email)
	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(req.Password)); err != nil {
		return err
	}
	user, _ := userModel.ToProtobuf()
	token, err := srv.Token.Encode(user)
	if err != nil {
		return err
	}
	res.Token = token
	fmt.Println(token)
	return nil
}
func (srv *UserService) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	claims, err := srv.Token.Decode(req.Token)
	if err != nil {
		return err
	}
	if claims.User.Id == "" {
		return errors.New("2", "无效的用户", 203)
	}
	res.Valid = true
	return nil
}

func (srv *UserService) Update(ctx context.Context, req *pb.User, res *pb.Response) error {
	if req.Id == "" {
		return errors.New("第一个错误", "用户ID不能为空", 200)
	}
	if req.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		req.Password = string(hash)
	}
	if req.Id == "" {
		return errors.New("", "id 不能为空", 300)
	}
	id, _ := strconv.ParseInt(req.Id, 10, 64)
	user, _ := srv.Repo.Get(uint(id))
	if err := srv.Repo.Update(user); err != nil {
		return err
	}
	res.User, _ = user.ToProtobuf()
	return nil
}
func (srv *UserService) CreatePasswordReset(ctx context.Context, req *pb.PasswordReset, res *pb.PasswordResetResponse) error {
	if req.Email == "" {
		return errors.New("", "邮箱不能为空", 444)
	}
	resetModel := new(model.PasswordReset)
	resetModel.ToORM(req)
	if err := srv.ResetRepo.Create(resetModel); err != nil {
		return err
	}
	if resetModel != nil {
		res.PasswordReset, _ = resetModel.ToProtobuf()
		err := srv.PublishEvent(res.PasswordReset)
		if err != nil {
			return err
		}
	}
	return nil
}
func (srv *UserService) ValidatePasswordResetToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	// 校验用户亲求中的token信息是否有效
	if req.Token == "" {
		return errors.New("", "Token信息不能为空", 300)
	}

	_, err := srv.ResetRepo.GetByToken(req.Token)
	if err != nil && err != gorm.ErrRecordNotFound {
		return errors.New("", "数据库查询异常", 120)
	}

	if err == gorm.ErrRecordNotFound {
		res.Valid = false
	} else {
		res.Valid = true
	}
	return nil
}
func (srv *UserService) DeletePasswordReset(ctx context.Context, req *pb.PasswordReset, res *pb.PasswordResetResponse) error {
	if req.Email == "" {
		return errors.New("", "Email信息不能为空", 300)
	}
	reset, err := srv.ResetRepo.GetByEmail(req.Email)
	if err != nil {
		return errors.New("", "数据库查询出错", 333)
	}
	err = srv.ResetRepo.Delete(reset)
	if err != nil {
		return err
	}
	res.PasswordReset = reset
	return nil
}

func (srv *UserService) PublishEvent(reset *pb.PasswordReset) error {
	body, err := json.Marshal(reset)
	if err != nil {
		return err
	}
	msg := broker.Message{
		Header: map[string]string{
			"email": reset.Email,
		},
		Body: body,
	}
	if err = srv.PubSub.Publish("password.reset", &msg); err != nil {
		log.Printf("[PUB} faield :%v", err)
	}
	return nil
}

func (srv *UserService) GetById(ctx context.Context, req *pb.User, res *pb.Response) error {
	name := "test" + req.Id
	user, _ := srv.Repo.GetByName(name)
	res.User, _ = user.ToProtobuf()
	return nil
}
