package dataservice

import (
	// "errors"
	// "net/http"
	"context"
	"errors"
	"fmt"
	"log"
	"userservice/model"

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
	ErrUserNotFound = errors.New("can't find account")
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
		u.DB.Rollback()
		return err
	}
	return u.DB.Commit().Error
}

func (u *UserDataService) FindAcc(ctx context.Context, username string) (*model.User, error) {
	defer func() {
		if r := recover(); r != nil {
		  u.DB.Rollback()
		}
	  }()

	log.Println("username")

	user := &model.User{}

    res := u.DB.Where("username = ?", username).Find(user)
    if res.Error != nil {
        log.Println("Error finding user:", res.Error)
		u.DB.Rollback()
        return nil, res.Error
    }

    if res.RowsAffected == 0 {
        return nil, ErrUserNotFound
    }

	return user, u.DB.Commit().Error
}

// func (u *UserDataService) Update(ctx context.Context, user *model.User) error {

// }




