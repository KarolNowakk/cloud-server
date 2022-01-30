package main

import (
	"cloud-watcher/proto/authpb"
	"context"

	"github.com/boltdb/bolt"
	"github.com/gookit/color"
	log "github.com/sirupsen/logrus"
)

type Auth struct {
	c authpb.AuthServiceClient
}

func MakeAuth(c authpb.AuthServiceClient) Auth {
	return Auth{c}
}

func (a Auth) Register() {
	req := &authpb.RegisterRequest{
		Email:                readInput("Email: "),
		Password:             readInput("Password: "),
		PasswordConfirmation: readInput("Repeat password: "),
	}

	_, err := a.c.Register(context.Background(), req)
	if err != nil {
		log.Fatalf("Error registering: %v", err)
	}

	color.Success.Tips("Successfully registered.")
}

func (a Auth) Login(db *bolt.DB) {
	req := &authpb.LoginRequest{
		Email:    readInput("Email: "),
		Password: readInput("Password: "),
	}

	res, err := a.c.Login(context.Background(), req)
	if err != nil {
		log.Fatalf("Error login: %v", err)
	}

	if err := saveToken(db, res.Token, res.UserID, res.ExpirationTime); err != nil {
		log.Fatalf("Error saving token: %v", err)
	}

	color.Success.Tips("Login successful")
}
