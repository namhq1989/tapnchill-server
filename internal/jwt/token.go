package appjwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
)

var unauthorizedError = errors.New("unauthorized")

func (j JWT) GenerateAccessToken(ctx *appcontext.AppContext, userID, clientID, timezone string) (string, error) {
	if userID == "" || clientID == "" {
		ctx.Logger().Error("invalid user id or platform id", nil, appcontext.Fields{"userID": userID, "clientID": clientID})
		return "", apperrors.Common.InvalidID
	}

	accessToken, _, err := j.generateAccessToken(userID, clientID, timezone)
	if err != nil {
		ctx.Logger().Error("failed to generate access token", err, appcontext.Fields{"userID": userID, "clientID": clientID})
		return "", err
	}

	return accessToken, nil
}

func (j JWT) generateAccessToken(userID, clientID, timezone string) (string, time.Time, error) {
	exp := time.Now().Add(j.accessTokenTTL)
	claims := &Claims{
		UserID:   userID,
		ClientID: clientID,
		Timezone: timezone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	value, err := token.SignedString(j.accessTokenSecret)
	return value, exp, err
}

func (j JWT) ParseAccessToken(ctx *appcontext.AppContext, token string) (*Claims, error) {
	if token == "" {
		return nil, unauthorizedError
	}

	// parse the token
	tokenData, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			ctx.Logger().Error("check signing method", fmt.Errorf("unexpected signing method: %v", t.Header["alg"]), appcontext.Fields{"token": token})
			return nil, unauthorizedError
		}

		return j.accessTokenSecret, nil
	})

	// error
	if err != nil {
		ctx.Logger().Error("parse token", err, appcontext.Fields{"token": token})
		return nil, err
	}

	// respond
	if claims, ok := tokenData.Claims.(*Claims); ok && tokenData.Valid {
		return claims, nil
	} else {
		ctx.Logger().Error("parse claims", nil, appcontext.Fields{"token": token, "tokenData": tokenData})
		return nil, unauthorizedError
	}
}
