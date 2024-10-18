package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user-management/entity"
)

var users []entity.User

// FindUser is a struct that is used to search for a user with the given name or email.
type FindUser struct {
	NameOrEmail string `json:"nameOrEmail"`
}

type UserLogin struct {
	NameOrEmail string `json:"nameOrEmail"`
	Password    string `json:"password"`
}

/*
CreateUser creates a new user and adds it to the list of users.
If the request is invalid, it will return a 400 status code.
If the user was created successfully, it will return a 201 status code.
*/
func PostCreateUser(c *gin.Context) {
	var userCreation entity.UserCreation

	if err := c.BindJSON(&userCreation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	var user, err = userCreation.MapToUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error})
		return
	}

	//Check if a user with the same email already exists
	for _, u := range users {
		if u.Email == user.Email {
			c.JSON(http.StatusBadRequest, gin.H{"message": "User already exists with this email"})
			return
		}
	}

	users = append(users, user)

	c.IndentedJSON(http.StatusCreated, user)
}

/*
FindUserWithNameOrEmail will search for a user with the given name or email.
If the user was found, it will return a 201 status code.
If the user was not found, it will return a 404 status code.
*/
func GetFindUserWithNameOrEmail(c *gin.Context) {
	var findUser FindUser
	if err := c.BindJSON(&findUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request" + err.Error()})
		return
	}

	for _, user := range users {
		if user.Username == findUser.NameOrEmail || user.Email == findUser.NameOrEmail {
			c.JSON(http.StatusCreated, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

func GetUserLogin(c *gin.Context) {
	var userLogin UserLogin
	if err := c.BindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request" + err.Error()})
		return
	}

	for _, user := range users {
		if (user.Username == userLogin.NameOrEmail || user.Email == userLogin.NameOrEmail) && user.Password == userLogin.Password {
			c.JSON(http.StatusCreated, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

/*
GetAListFromAllUser will return a list of all users
It will return a 201 status code. If there are no users, it will return an empty list.
*/
func GetAListFromAllUser(c *gin.Context) {
	c.JSON(http.StatusCreated, users)
}
