package middleware

import (
	"net/http"
)

type JwtContext struct {
	echo.Context
}

func GetJwtValidate() *JwtContext {
	return &JwtContext{}
}

func (j *JwtContext) ValidateJWT(tokenString string) echo.HandlerFunc {
	return func(c echo.Context) error {

		// tokenChan := make(chan string, 1)

		authHeader := c.Request().Header("Authorization")

		if authHeader == "" || !strings.Contains(authHeader, "Bearer") {
			http.Error(c.Response(), "Invalid Token", http.StatusBadRequest)
			c.Logger().Error("Invalid token")
			// echo.NewHTTPError(http.StatusBadRequest, "Invalid Token")
			return errors.New("Invalid token")
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
	
		if err != nil {
			return err
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(c.Response(), "error parse token", http.StatusBadRequest)
			return err
		}

		c.Set("token", claims)
		jCtx := &JwtContext{c}
		return next(jCtx)
	}
}
