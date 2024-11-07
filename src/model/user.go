package model

import (
	"gin_REST_API_ex/src/config"
	token "gin_REST_API_ex/src/util"
	"html"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (user *User) BeforeSave() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	user.Password = string(hashedPassword)

	return nil
}

func GetUserByID(uid uint) (User, error) {
	var user User

	err := config.MongoClient.Collection("user").FindOne(config.CTX, bson.M{"id": uid}).Decode(&user)

	if err != nil {
		return User{}, err
	}

	user.PrepareUserPublicly()

	return user, nil
}

func LoginCheck(username string, password string) (string, error) {
	var err error
	user := User{}

	err = config.MongoClient.Collection("user").FindOne(config.CTX, bson.M{"username": username}).Decode(&user)

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (user *User) PrepareUserPublicly() {
	user.Password = ""
}

func (user *User) SaveUser() (*User, error) {
	user.BeforeSave()
	_, err := config.MongoClient.Collection("user").InsertOne(config.CTX, user)

	if err != nil {
		return &User{}, err
	}

	return user, nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
