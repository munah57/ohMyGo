package handler

import (
	
	"vaqua/services"

	"github.com/gin-gonic/gin"
)


type UserHandler struct {
	Service *services.UserService
}

func (h *UserHandler) SignupUser(c *gin.Context) {

}