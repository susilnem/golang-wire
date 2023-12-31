package routes

import (
	"go-wire/pkg/api/handler"

	"github.com/gin-gonic/gin"
)

func UserRoutes(api *gin.RouterGroup, userHandler *handler.UserHandler) {
	// signup routes for user
	signup := api.Group("/user")
	{
		signup.POST("/signup", userHandler.UserSignUp)
	}

}
