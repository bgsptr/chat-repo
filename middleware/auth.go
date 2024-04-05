package middleware

import (
	"errors"
	"net/http"
	"os"

	// "strconv"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

type CustomClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

var (
	secretKey = os.Getenv("JWT_SECRET_KEY")
	// jwtMethod = os.Getenv("JWT_METHOD")
	jwtMethod = jwt.SigningMethodHS256
)

var (
	errSignMehtod = errors.New("error method signed")
)

type JwtContext struct {
	echo.Context
}

func GetJwtValidate() *JwtContext {
	return &JwtContext{}
}

func (j *JwtContext) ValidateJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		// tokenChan := make(chan string, 1)

		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" || !strings.Contains(authHeader, "Bearer") {
			http.Error(c.Response(), "Invalid Token", http.StatusBadRequest)
			c.Logger().Error("Invalid token")
			// echo.NewHTTPError(http.StatusBadRequest, "Invalid Token")
			return errors.New("Invalid token")
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errSignMehtod
			} else if method != jwtMethod {
				return nil, errSignMehtod
			}
		
			return secretKey, nil
		})
	
		if err != nil {
			return err
		}

		claims, ok := token.Claims.(*CustomClaims)
		if !ok || !token.Valid {
			http.Error(c.Response(), "error parse token", http.StatusBadRequest)
			return err
		}

		c.Set("username", claims.Username)
		jCtx := &JwtContext{c}
		return next(jCtx)
	}
}
