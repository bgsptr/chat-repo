package usecase

import (
	"context"
	"errors"
	"time"
	"userservice/dataservice"

	// "os"
	"log"
	"userservice/model"

	jwt "github.com/golang-jwt/jwt/v4"
	// "github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/aws/session"
	// "github.com/aws/aws-sdk-go/service/s3"
	// "github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var (
	JWT_SIGNING_METHOD = jwt.SigningMethodHS256
	JWT_SIGNATURE_KEY = []byte("the secret of kalimdor")
	LOGIN_EXPIRATION_DURATION = 1 * time.Hour 
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

func (u *UserService) Find(ctx context.Context, user *model.User) (string, error) {
	
	userFetched, err := u.UserDataService.FindAcc(ctx, user.Username)
	if err != nil {
		log.Println("user")
		return "", errors.New("failed to create acc in service layer")
	}

	// bcrypt password
	if userFetched.Pass != user.Pass {
		return "", errors.New("password not match")
	}

	expirationTime := time.Now().Add(LOGIN_EXPIRATION_DURATION)
	expire := jwt.NewNumericDate(expirationTime)

	claims := &model.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expire,
		},
		Username: userFetched.Username,
	}

	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		return "", err
	}

	return signedToken, nil
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