package token

import (
	"errors"
	"time"
	pb "user-service/genproto/userpb"

	"github.com/form3tech-oss/jwt-go"
)

var tokenKey = "secret_key"

func GenerateJwtToken(user *pb.User) (*pb.Tokens, error) {

	accessToken := jwt.New(jwt.SigningMethodHS256)
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	claims := accessToken.Claims.(jwt.MapClaims)
	claims["id"] = user.Id
	claims["username"] = user.UserName
	claims["password_hash"] = user.PasswordHash
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(60 * time.Minute).Unix()
	access, err := accessToken.SignedString([]byte(tokenKey))
	if err != nil {
		return nil, err
	}

	rclaims := refreshToken.Claims.(jwt.MapClaims)
	rclaims["id"] = user.Id
	rclaims["username"] = user.UserName
	rclaims["password_hash"] = user.PasswordHash
	rclaims["iat"] = time.Now().Unix()
	rclaims["exp"] = time.Now().Add(60 * time.Minute).Unix()
	refresh, err := refreshToken.SignedString([]byte(tokenKey))
	if err != nil {
		return nil, err
	}
	return &pb.Tokens{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

func ParseToken(tokenStr string) (jwt.MapClaims, error) {

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenKey), nil
	}
	token, err := jwt.Parse(tokenStr, keyFunc)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
