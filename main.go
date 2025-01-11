package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"main.go/initializers"
	"main.go/routers"
)

var r *gin.Engine

func init() {
	r = gin.Default()
	r.LoadHTMLGlob("views/*")
	r.Static("/assets", "./assets")
	initializers.LoadEnv()
	initializers.DBconnection()

}

func main() {
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	routers.Routes(r)
	r.Run()
}
