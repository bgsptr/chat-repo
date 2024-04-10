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

var (
	ErrUserNotFound = errors.New("failed to find user")
	ErrInvalidPassword = errors.New("passwords do not match")
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
        log.Println("Failed to find user:", err)
        return "", ErrUserNotFound
    }

    // Compare passwords securely
    if !comparePasswords(userFetched.Pass, user.Pass) {
        return "", ErrInvalidPassword
    }

    expirationTime := time.Now().Add(LOGIN_EXPIRATION_DURATION)
    expire := jwt.NewNumericDate(expirationTime)

    claims := &model.Claims{
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: expire,
        },
        Username: userFetched.Username,
    }

    token := jwt.NewWithClaims(JWT_SIGNING_METHOD, claims)
    signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
    if err != nil {
        log.Println("Failed to sign JWT token:", err)
        return "", errors.New("failed to sign JWT token")
    }

    // Generate and store refresh token
    // refreshToken, err := generateRefreshToken()
    // if err != nil {
    //     return "", errors.New("failed to generate refresh token")
    // }

    // Store refresh token securely

    return signedToken, nil
}

func comparePasswords(hashedPassword, password string) bool {
    // Implement a secure password comparison method, such as bcrypt.CompareHashAndPassword
    // Example:
    // err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    // return err == nil
    return hashedPassword == password // Temporary insecure comparison for demonstration
}

func (u *UserService) ErrUserNotFound() error {
	return ErrUserNotFound
}

func (u *UserService) ErrInvalidPassword() error {
	return ErrInvalidPassword
}

// func (u *UserService) Refresh() (string, error) {
	
// }

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