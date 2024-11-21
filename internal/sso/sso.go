package sso

import (
	"context"
	"encoding/base64"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/namhq1989/go-utilities/appcontext"
	"google.golang.org/api/option"
)

type Operations interface {
	VerifyGoogleToken(ctx *appcontext.AppContext, idToken string) (*UserInfo, error)
}

type SSO struct {
	firebase *auth.Client
}

func NewSSOClient(firebaseSAEncoded string) *SSO {
	ctx := context.Background()

	sa, err := decodeBase64(firebaseSAEncoded)
	if err != nil {
		panic(err)
	}

	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsJSON(sa))
	if err != nil {
		panic(err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		panic(err)
	}

	return &SSO{firebase: client}
}

func decodeBase64(encoded string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}
