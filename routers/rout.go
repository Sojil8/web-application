package routers

import (
	"github.com/gin-gonic/gin"
	"main.go/controllers"
)

func Routes(r *gin.Engine) {
	//index
	r.GET("/", controllers.Indexhandler)

	//user
	r.GET("/user/login", controllers.UserLoginHandler)
	r.POST("/login", controllers.UserLoginPostHandler)
	r.GET("/user/signup", controllers.UserSignUpHandler)
	r.POST("/signup", controllers.UserSignUpPostHandler)
	r.GET("/home", controllers.UserHomeHandler)
	r.GET("/logout", controllers.UserLogoutHandler)

	//admin
	r.GET("/admin/login", controllers.AdminHandler)
	r.POST("/adminlogin", controllers.AdminLoginpostHandler)
	r.GET("/adminpannel", controllers.AdminpannelHandler)
	r.GET("/admin/search", controllers.AdminSearchHandler)
	r.POST("/admin/create-user", controllers.UserCreateHandler)
	r.GET("/admin/edit-user/:ID", controllers.EditHandler)
	r.POST("/admin/edit-user/:ID", controllers.EditPostHandler)
	r.GET("/admin/block-user/:ID", controllers.UserSatusHandler)
	r.GET("admin/delete-user/:ID", controllers.UserDeletHandler)
	r.GET("adminlogout", controllers.LogoutHandler)
}
