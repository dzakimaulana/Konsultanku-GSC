package databases

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"firebase.google.com/go/storage"
	"google.golang.org/api/option"
)

var AuthMd *auth.Client

type Auth struct {
	auth *auth.Client
}

type Storage struct {
	storage *storage.Client
}

func FirebaseConn() (*Auth, *Storage, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile("pkg/databases/firebase-sdk.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic(err.Error())
	}

	auth, err := app.Auth(ctx)
	if err != nil {
		panic(err.Error())
	}
	AuthMd = auth

	storage, err := app.Storage(ctx)
	if err != nil {
		panic(err.Error())
	}

	return &Auth{auth: auth}, &Storage{storage: storage}, nil
}

func (a *Auth) GetAuth() *auth.Client {
	return a.auth
}

func (s *Storage) GetStorage() *storage.Client {
	return s.storage
}
