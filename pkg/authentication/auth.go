package authentication

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/xrexy/togo/config"
	"github.com/xrexy/togo/pkg/database"
)

type Authentication struct {
	config config.EnvVars
}

var instance *Authentication
var lock = &sync.Mutex{}

func New() *Authentication {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			config, err := config.LoadConfig()
			if err != nil {
				panic(err)
			}

			instance = &Authentication{config}
		}
	}

	return instance
}

func (a *Authentication) CreateJWT(user database.User) (string, time.Time, error) {
	tokenByte := jwt.New(jwt.SigningMethodHS256)

	now := time.Now()
	exp := now.Add(time.Hour * 24)

	claims := tokenByte.Claims.(jwt.MapClaims)

	claims["sub"] = user.UUID
	claims["iss"] = a.config.JWT_ISSUER

	claims["exp"] = exp.Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	token, err := tokenByte.SignedString([]byte(a.config.JWT_SECRET))
	if err != nil {
		return "", exp, err
	}

	return token, exp, nil
}

// VerifyJWT is a function to verify JWT token
// and return the claims
func (a *Authentication) VerifyJWT(token string) (jwt.MapClaims, error) {
	if token == "" {
		return nil, fmt.Errorf("token is empty")
	}

	tokenByte, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}

		return []byte(a.config.JWT_SECRET), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return nil, fmt.Errorf("invalid token claim")
	}

	return claims, nil
}

func (a *Authentication) GetTokenString(c *fiber.Ctx) string {
	var tokenString string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies(a.config.JWT_SECRET) != "" {
		tokenString = c.Cookies(a.config.JWT_SECRET)
	}

	return tokenString
}
