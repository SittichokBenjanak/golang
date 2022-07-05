package routes

import (
	usercontroller "example.com/gin-backend-api/controllers/user"
	"example.com/gin-backend-api/middlewares"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(rg *gin.RouterGroup) {
	// routerGroup := rg.Group("/users").Use(middlewares.AuthJWT())
	routerGroup := rg.Group("/users")

	//{{domain_url}}/api/v1/users/
	routerGroup.GET("/", usercontroller.GetAll)

	//{{domain_url}}/api/v1/users/register
	routerGroup.POST("/register", usercontroller.Register)

	//{{domain_url}}/api/v1/users/login
	routerGroup.POST("/login", usercontroller.Login)

	//{{domain_url}}/api/v1/users/3
	routerGroup.GET("/:id", usercontroller.GetById)

	//{{domain_url}}/api/v1/users/search?fullname=John&age=10
	routerGroup.GET("/search", usercontroller.SearchByFullname)

	//get Profile
	routerGroup.GET("/me", middlewares.AuthJWT(), usercontroller.GetProfile)

}
