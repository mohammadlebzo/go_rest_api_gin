package model

import (
	"strings"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

var user User = User{ID: uint(0), Username: "raven walt", Password: "testRaven621"}

func TestBeforeSave(t *testing.T) {
	user.BeforeSave()

	if !strings.Contains(user.Password, "$2a$10$") {
		t.Error("The password should be encrypted. Recived value:", user.Password)
	}
}

func TestPrepareUserPublicly(t *testing.T) {
	user.PrepareUserPublicly()

	if user.Password != "" {
		t.Error("The user password should be an empty string. Recived value:", user.Password)
	}
}

func TestVerifyPassword(t *testing.T) {
	unencryptedUserPassword := user.Password
	user.BeforeSave()

	if VerifyPassword(unencryptedUserPassword, user.Password) != nil && VerifyPassword(unencryptedUserPassword, user.Password) == bcrypt.ErrMismatchedHashAndPassword {
		t.Error("The provided password does not match the encrypted one. Recived value:", user.Password, unencryptedUserPassword)
	}
}

// func TestGetUserByID(t *testing.T) {
// 	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
// 	// defer mt.

// 	mt.Run("find & transform", func(mt *mtest.T) {
// 		mt.AddMockResponses(mtest.CreateCursorResponse(0, "user.id", mtest.FirstBatch, bson.D{
// 			{Key: "id", Value: user.ID},
// 			{Key: "username", Value: user.Username},
// 			{Key: "password", Value: user.Password},
// 		}))

// 		response, err := GetUserByID(user.ID)
// 		if err != nil {
// 			t.Error(err)
// 		}

// 		if response.ID != user.ID {
// 			t.Error("Returned user id does not match the expected. Recived value:", response.ID)
// 		}
// 	})

// }
