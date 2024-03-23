package usecase

import (
	"context"
	"errors"
	"userservice/dataservice"
	// "os"
	"userservice/model"
	"log"
	// "github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/aws/session"
	// "github.com/aws/aws-sdk-go/service/s3"
	// "github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type UserService struct {
	UserDataService *dataservice.UserDataService
}

// func (u *UserService) UploadImage() {

// }

func NewUserService(userDataService *dataservice.UserDataService) *UserService {
	return &UserService{
		UserDataService: userDataService,
	}
}

func (u *UserService) CreateAccount(ctx context.Context, user *model.User) error {
	err := u.UserDataService.Create(ctx, user)
	if err != nil {
		log.Println("user")
		return errors.New("failed to create acc in service layer")
	}

	return nil
}

// func (u *UserService) UploadFile(uploader *s3manager.Uploader, filePath string, bucketName string, fileName string) error {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		return err
// 	}

// 	defer file.Close()

// 	_, err = uploader.Upload(&s3manager.UploadInput{
// 		Bucket: aws.String(bucketName),
// 		Key:    aws.String(fileName),
// 		Body:   file,
// 	})

// 	return err
// }