package controller

import (
	"gin_REST_API_ex/src/config"
	"gin_REST_API_ex/src/model"
	token "gin_REST_API_ex/src/util"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Login(c *gin.Context) {
	user := model.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := model.LoginCheck(user.Username, user.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Logout(c *gin.Context) {
	var bToken struct{ Token string }
	tokenString := token.ExtractToken(c)

	bToken.Token = tokenString

	_, err := config.MongoClient.Database(os.Getenv("DB_NAME")).Collection("blocklist").InsertOne(config.CTX, bToken)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logout Success"})
}

func CreateUser(c *gin.Context) {
	user := model.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := user.SaveUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration Success"})
}

func GetCurrentAuthUser(c *gin.Context) {
	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := model.GetUserByID(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	u64, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		log.Println(err)
	}

	user, err := model.GetUserByID(uint(u64))

	if err != nil {
		return
	}

	user.PrepareUserPublicly()

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}

func GetUsers(c *gin.Context) {
	var users []model.User
	cursor, err := config.MongoClient.Database(os.Getenv("DB_NAME")).Collection("user").Find(config.CTX, bson.M{})

	if err != nil {
		log.Fatal(err)
		return
	}

	if err = cursor.All(config.CTX, &users); err != nil {
		log.Fatal(err)
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}
