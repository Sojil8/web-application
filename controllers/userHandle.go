package controllers

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"main.go/auth"
	"main.go/initializers"
	"main.go/models"
)

var Error string
var FeachUser models.UserModel

func UserLoginHandler(c *gin.Context) {
	c.Header("Cache-control", "no-cache,no-store,must-revalidate")

	session := sessions.Default(c)
	userCheck := session.Get(RollUser)

	if userCheck != nil {
		c.Redirect(http.StatusSeeOther, "/home")
	} else {
		c.HTML(http.StatusSeeOther, "userlogin.html", Error)
		Error = ""
	}
}

func UserLoginPostHandler(c *gin.Context) {
	err := initializers.DB.First(&FeachUser, "email=?", c.Request.FormValue("email")).Error

	if err != nil {
		Error = "Invalid Email Address"
		c.Redirect(http.StatusSeeOther, "user/login")
	} else {
		plainPassword := c.Request.FormValue("password")

		err := bcrypt.CompareHashAndPassword([]byte(FeachUser.Password), []byte(plainPassword))

		if err != nil {
			Error = "invalid password"
			c.Redirect(http.StatusSeeOther, "user/login")
		} else {
			if FeachUser.Status == "Blocked" {
				Error = "your Account has been blocked"
				c.Redirect(http.StatusSeeOther, "user/login")
			} else {
				auth.JwtTokens(c, FeachUser.Email, RollUser)
				c.Redirect(http.StatusSeeOther, "/home")
			}
		}

	}
}

func UserSignUpHandler(c *gin.Context) {
	c.Header("Cache-control", "no-cache,no-store,must-revalidate")

	session := sessions.Default(c)
	userCheck := session.Get(RollUser)

	if userCheck != nil {
		c.Redirect(http.StatusSeeOther, "/home")
	} else {
		c.HTML(http.StatusSeeOther, "usersignup.html", Error)
		Error = ""
	}
}

func UserSignUpPostHandler(c *gin.Context) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(c.Request.PostFormValue("password")), 10)
	if err != nil {
		log.Fatal("Error in password hashing", err)
	}

	error := initializers.DB.Create(&models.UserModel{
		Name:     c.Request.PostFormValue("username"),
		Email:    c.Request.PostFormValue("email"),
		Password: string(hashedPassword),
		Status:   "Active",
	})

	if error.Error != nil {
		Error = "Email already exists "
		c.Redirect(http.StatusSeeOther, "/user/signup")
	} else {
		Error = "User successfully Created "
		c.Redirect(http.StatusSeeOther, "/user/login")
	}
}

func UserHomeHandler(c *gin.Context) {
	c.Header("Cache-control", "no-cache,no-store,must-revalidate")

	session := sessions.Default(c)
	check := session.Get(RollUser)

	if check != nil {
		c.HTML(http.StatusSeeOther, "userhome.html", FeachUser)
	} else {
		c.Redirect(http.StatusSeeOther, "/user/login")
	}
}

func UserLogoutHandler(c *gin.Context) {
	c.Header("Cache-control", "no-cace,no-store,must-revalidate")

	session := sessions.Default(c)
	session.Delete(RollUser)
	session.Save()
	FeachUser = models.UserModel{}

	c.Redirect(http.StatusSeeOther, "/user/login")
}
