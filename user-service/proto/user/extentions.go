package laracom_service_user

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	fmt.Println("最前面哦")
	_ = scope.SetColumn("CreatedAt", time.Now().Format(time.RFC3339))
	uuid := uuid.NewV4()
	fmt.Println("laidao middleware")
	return scope.SetColumn("Id", uuid.String())
}
func (model *User) BeforeSave(scope *gorm.Scope) error {
	_ = scope.SetColumn("UpdatedAt", time.Now().Format(time.RFC3339))
	return nil
}
func (model *PasswordReset) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("CreatedAt", time.Now().Format(time.RFC3339))
	return nil
}
