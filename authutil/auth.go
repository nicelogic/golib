package authutil

import (
	"context"
	"crypto/rsa"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Auth struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

type contextKey struct {
	name string
}

var userCtxKey = &contextKey{name: "user"}
var errorCtxKey = &contextKey{name: "error"}

type User struct {
	Id    string `json:"id"`
	Token string
}

func GetUser(ctx context.Context) (*User, error) {
	user, _ := ctx.Value(userCtxKey).(*User)
	err, _ := ctx.Value(errorCtxKey).(error)
	return user, err
}

func (auth *Auth) UserFromJwt(tokenString string) (user *User, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("recovered error(%w)", err)
		}
	}()

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method(%v)", token.Header["alg"])
		}
		return auth.publicKey, nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return
	}
	userMap, ok := claims["user"].(map[string]interface{})
	if !ok {
		err = errors.New("claims[user] is not map[string]interface{}")
		return
	}
	id, ok := userMap["id"].(string)
	if !ok {
		err = errors.New("userMap[id] is not string")
		return
	}
	user = &User{Id: id,
		Token: tokenString}
	return
}

func (auth *Auth) Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			reqToken := request.Header.Get("Authorization")
			splitToken := strings.Split(reqToken, "Bearer ")
			var ctx context.Context
			if len(splitToken) != 2 {
				ctx = context.WithValue(request.Context(), errorCtxKey, errors.New(AuthUtilErrorHttpHeaderAuthorizationInvalid))
			} else {
				jwtToken := splitToken[1]
				user, err := auth.UserFromJwt(jwtToken)
				if err != nil {
					fmt.Printf("userFromJwt(%v)\n", err)
					ctx = context.WithValue(request.Context(), errorCtxKey, err)
				} else {
					ctx = context.WithValue(request.Context(), userCtxKey, user)
				}
			}
			request = request.WithContext(ctx)
			next.ServeHTTP(writer, request)
		})
	}
}

func (auth *Auth) InitForAuth(publicKeyFilePath string) error {
	return auth.Init(publicKeyFilePath, "")
}

func (auth *Auth) Init(publicKeyFilePath string, privateKeyFilePath string) error {
	// publicKey, err := os.ReadFile("/etc/app-0/secret-jwt/jwt-publickey")
	publicKey, err := os.ReadFile(publicKeyFilePath)
	if err != nil {
		log.Printf("read public key file error(%v)\n", err)
		return err
	}
	auth.publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		log.Printf("ParseRSAPublicKeyFromPEM error(%v)\n", err)
		return err
	}
	// privateKey, err := os.ReadFile("/etc/app-0/secret-jwt/jwt-privatekey")
	if privateKeyFilePath == "" {
		log.Printf("private key file path is empty, will not support sign token\n")
	} else {
		privateKey, err := os.ReadFile(privateKeyFilePath)
		if err != nil {
			log.Printf("read private key file error(%v)\n", err)
			return err
		}
		auth.privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateKey)
		if err != nil {
			log.Printf("ParseRSAPrivateKeyFromPEM error(%v)\n", err)
			return err
		}
	}
	return nil
}

func (auth *Auth) SignToken(userId string, expireDuration time.Duration) (jwtToken string, err error) {
	user := User{Id: userId}
	mapClaims := make(jwt.MapClaims)
	mapClaims["user"] = user
	mapClaims["exp"] = time.Now().Add(expireDuration).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, mapClaims)
	jwtToken, err = token.SignedString(auth.privateKey)
	if err != nil {
		log.Printf("error signing token(%v)\n", err)
		return "", err
	}
	log.Println(jwtToken)
	return
}
