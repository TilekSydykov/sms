package utils

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"reflect"
	"solar-faza/entity"
	"solar-faza/repository"
	"strconv"
)

type AuthTokens struct {
	AccessToken string `json:"access_token"`
}

var connection repository.UserRepository

func SingIn(email string, password string) (*AuthTokens, error) {
	connection = repository.NewUserRepository()
	user := connection.GetUserByEmail(email)
	if user == nil {
		return &AuthTokens{}, errors.New("no user found")
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return &AuthTokens{}, err
	}
	tokens := ObtainTokenPair(user)
	return tokens, nil
}

func RegisterUser(user *entity.User) error {
	connection = repository.NewUserRepository()
	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(password)
	u := connection.GetUserByEmail(user.Email)
	if u != nil && u.Id > 0 {
		return errors.New("such user already exists")
	}
	connection.Create(user)
	if user.Id == 0 {
		return errors.New("such user already exists")
	}
	return nil
}

func ObtainTokenPair(user *entity.User) *AuthTokens {
	return &AuthTokens{
		AccessToken: newAccessToken(user),
	}
}

func newAccessToken(user *entity.User) string {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.FormatUint(user.Id, 10),
	})
	accessToken, _ := claims.SignedString([]byte("secretKey"))
	return accessToken
}

func GetDataFromToken(tokenstring string) uint64 {
	claims := jwt.MapClaims{}
	_, _ = jwt.ParseWithClaims(tokenstring, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secretKey"), nil
	})
	for key, val := range claims {
		if key == "iss" {
			parseUint, err := strconv.ParseUint(fmt.Sprintf("%v", val), 10, 64)
			if err != nil {
				return 0
			}
			return parseUint
		}
	}
	return 0
}

func GenerateCookies(c *gin.Context, data string) {
	// domain := env.GetEnvVars().Host
	c.SetCookie("data", data, 60*60*24*30, "/", "", false, true)
}

func GetType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}
