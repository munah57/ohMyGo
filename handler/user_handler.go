package handler

import (
	"fmt"
	"net/http"
	"net/mail"
	"vaqua/models"
	"vaqua/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service *services.UserService
}

func (h *UserHandler) SignUpNewUserAcct(c *gin.Context) {
	var newUser models.SignUpRequest
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	fmt.Println("Signup request received:", newUser.Email)

	if _, err := mail.ParseAddress(newUser.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	if len(newUser.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password should contain minimum 6 characters"})
		return
	}

	// call the service layer
	err := h.Service.SignUpNewUserAcct(&newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// response
	c.JSON(http.StatusCreated, gin.H {
		"message": "User created successfully",
		"email": newUser.Email,
	})

}

func (h *UserHandler) LoginUser(c *gin.Context) {
    var request models.LoginRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
        return
    }
	fmt.Println("error logging in")
    token, err := h.Service.LoginUser(request)
	fmt.Println("error logging in 2")
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *UserHandler) LogoutUser(c *gin.Context) {
	err := h.Service.LogoutUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}