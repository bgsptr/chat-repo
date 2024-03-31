package dataservice

import (
	// "errors"
	// "net/http"
	"context"
	"fmt"
	"userservice/model"
	"log"
	"gorm.io/gorm"
	// "gorm.io/driver/mysql"

	"userservice/dataservice/usersqldb"
)

type UserDataService struct {
	DB *gorm.DB
}

// type PostRepositoryIntrfc interface {
// 	Create() (*model.User, interface{}, error)
// 	Delete() error
// 	Update(username string) (interface{}, error)
// }

var (
	ErrFindAccount = errors.New("can't find account")
)

func NewUserData() *UserDataService {
	log.Println("user")
	db, err := usersqldb.NewGorm()
	// db.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Println("Failed to initialize database:", err)
        return nil
	}

	tx := db.Begin()
	return &UserDataService{
		DB: tx,
	}
}

func (u *UserDataService) Create(ctx context.Context, user *model.User) error {
	log.Println("user")
	defer func() {
		if r := recover(); r != nil {
		  u.DB.Rollback()
		}
	  }()
	
	if err := u.DB.Create(&user).Error; err != nil {
		return err
	}
	return u.DB.Commit().Error
}

func (u *UserDataService) Find(ctx context.Context, username string) (*model.User, error) {
	log.Println(username)

	res := u.DB.Where("username = ?", username).Find(&model.User)
	if res.Err != nil {
		return nil, ErrFindAccount
	}

	return res, nil
}

// func (u *UserDataService) Update(ctx context.Context, user *model.User) error {

// }




