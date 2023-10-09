package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/yousefzinsazk78/simple_note_api/internal/models"
	"github.com/yousefzinsazk78/simple_note_api/internal/utils"
)

// Register user
func (h *handler) Register(context *gin.Context) {
	var input models.Register

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
		RoleID:   3,
	}

	err := h.userStorer.InsertUser(user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"result": "user inserted successfully."})

}

// User Login
func (h *handler) Login(context *gin.Context) {
	var input models.Login

	if err := context.ShouldBindJSON(&input); err != nil {
		var errorMessage string
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			validationError := validationErrors[0]
			if validationError.Tag() == "required" {
				errorMessage = fmt.Sprintf("%s not provided", validationError.Field())
			}
		}
		context.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	user, err := h.userStorer.GetUserByUsername(input.Username)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidatePassword(input.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := utils.GenerateJWT(*user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": jwt, "username": input.Username, "message": "Successfully logged in"})

}

// get all users
func (h *handler) GetUsers(context *gin.Context) {
	var user []models.User
	user, err := h.userStorer.GetUsers()
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, user)
}

// get user by id
func (h *handler) GetUser(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var user *models.User
	user, err := h.userStorer.GetUserByID(id)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, user)
}

// update user
func (h *handler) UpdateUser(c *gin.Context) {
	//var input model.Update
	var User *models.User
	var userRole models.UserRole
	id, _ := strconv.Atoi(c.Param("id"))

	User, err := h.userStorer.GetUserByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&userRole)
	err = h.userStorer.UpdateUser(userRole)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, User)
}
