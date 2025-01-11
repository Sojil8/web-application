package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const RollUser = "user"
const RollAdmin = "admin"

func Indexhandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
