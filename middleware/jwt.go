package middleware

import (
	"errors"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := ParseMapClaimsJwtHeader(ctx)
		if err != nil {
			ctx.Abort()
			ctx.String(http.StatusOK, err.Error())
			return
		}
		ctx.Next()
	}
}

var (
	secret          []byte
	ErrInvalidToken = errors.New("invalid token")
	Expiration      = time.Hour * 24 * 7
)

func init() {
	secret = make([]byte, 256)
	rand.Seed(time.Now().Unix())
	rand.Read(secret)
}

func BuildMapClaimsJwt(username, password string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"password": password,
		"expire":   time.Now().Add(Expiration).Unix(),
	})
	return token.SignedString(secret)
}

func ParseMapClaimsJwt(pJwt string) (jwt.MapClaims, error) {
	var err error
	token, err := jwt.Parse(pJwt, func(tk *jwt.Token) (interface{}, error) {
		if _, ok := tk.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return secret, nil
	})
	if err != nil {
		logrus.WithField("error", err).Debugln("failed to parse jwt!")
		return nil, ErrInvalidToken
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check expire
		if expire, ok := claims["expire"].(float64); ok && int64(expire) > time.Now().Unix() {
			return claims, nil
		}
		return nil, ErrInvalidToken
	}
	return nil, ErrInvalidToken
}

func ParseMapClaimsJwtHeader(ctx *gin.Context) (jwt.MapClaims, error) {
	if hd, ok := ctx.Request.Header["Authorization"]; ok {
		if len(hd) != 1 {
			return nil, ErrInvalidToken
		}
		token := strings.ReplaceAll(hd[0], "Bearer ", "")
		return ParseMapClaimsJwt(token)
	}
	return nil, ErrInvalidToken
}
