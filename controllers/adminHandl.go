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

var GetAdmin models.AdminModel
var AdError string

const RoleAdmin = "admin"

func AdminHandler(c *gin.Context) {
	c.Header("Cache-control", "no-cache,no-store,must-revalidate")
	session := sessions.Default(c)
	adminCheck := session.Get(RollAdmin)
	if adminCheck != nil {
		c.Redirect(http.StatusSeeOther, "/adminpannel")
	} else {
		c.HTML(http.StatusSeeOther, "AdminLogin.html", AdError)
		AdError = ""
	}
}

func AdminLoginpostHandler(c *gin.Context) {
	err := initializers.DB.First(&GetAdmin, "email=?", c.Request.PostFormValue("email")).Error

	if err != nil {
		AdError = "Invalid Email Address"
		c.Redirect(http.StatusSeeOther, "/admin/login")
	} else {
		if c.Request.PostFormValue("password") != GetAdmin.Password {
			AdError = "Invalid password"
			c.Redirect(http.StatusSeeOther, "/admin/login")
		} else {
			auth.JwtTokens(c, GetAdmin.Email, RoleAdmin)
			c.Redirect(http.StatusSeeOther, "/adminpannel")
		}
	}
}

func AdminpannelHandler(c *gin.Context) {
	c.Header("Cache-control", "no-cache,no-store,must-revalidate")
	var UserData []models.UserModel
	session := sessions.Default(c)
	check := session.Get(RoleAdmin)

	if check != nil {
		initializers.DB.Find(&UserData)
		c.HTML(http.StatusSeeOther, "adminHome.html", gin.H{
			"UserDatas": UserData,
			"Admin":     GetAdmin.Name,
			"Error":     AdError,
		})
		AdError = ""
	} else {
		c.Redirect(http.StatusSeeOther, "/admin/login")
	}
}

func AdminSearchHandler(c *gin.Context) {
	c.Header("Cache-control", "no-cache,no-store,must-revalidate")
	session := sessions.Default(c)
	check := session.Get(RollAdmin)

	if check != nil {
		query := c.Query("query")

		if query == "" {
			AdError = " Search can't be empty"
			c.Redirect(http.StatusSeeOther, "/adminpannel")
			return
		}

		var matchedUsers []models.UserModel
		searchQuery := query + "%"
		err := initializers.DB.Where("name iLike ? OR email iLike ?", searchQuery, searchQuery).Find(&matchedUsers).Error

		if err != nil {
			AdError = "Error Feaching Search Result"
			c.Redirect(http.StatusSeeOther, "/adminpannel")
			return
		}

		c.HTML(http.StatusOK, "adminHome.html", gin.H{
			"UserDatas": matchedUsers,
			"Admin":     GetAdmin.Name,
			"Error":     AdError,
		})
		AdError = ""
	} else {
		c.Redirect(http.StatusSeeOther, "/admin/login")
	}
}

func UserCreateHandler(c *gin.Context) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(c.Request.PostFormValue("password")), 10)

	if err != nil {
		log.Fatal("Error in passwrod hashing", err)
	}

	DBerror := initializers.DB.Create(&models.UserModel{
		Name:     c.Request.PostFormValue("username"),
		Email:    c.Request.PostFormValue("email"),
		Password: string(hashedPassword),
		Status:   "Active",
	})

	if DBerror.Error != nil {
		AdError = "User Not Created Email Already Existing"
	} else {
		AdError = "User Successfully Created"
	}
	c.Redirect(http.StatusSeeOther, "/adminpannel")
}

func EditHandler(c *gin.Context) {
	c.Header("Cache-control", "no-cache,no-store,must-revalidate")
	session := sessions.Default(c)
	check := session.Get(RollAdmin)

	if check != nil {
		UserId := c.Param("ID")
		c.HTML(http.StatusSeeOther, "edit.html", UserId)
	} else {
		c.Redirect(http.StatusSeeOther, "/admin/login")
	}
}

func EditPostHandler(c *gin.Context) {
	var UserUpdate models.UserModel
	UserId := c.Param("ID")
	initializers.DB.Find(&UserUpdate, "ID=?", UserId)
	UserUpdate.Name = c.Request.FormValue("username")
	UserUpdate.Email = c.Request.FormValue("email")

	err := initializers.DB.Save(&UserUpdate).Error
	if err != nil {
		AdError = "Update Not Saved Email Already Existing"
	} else {
		AdError = "Update Save"
	}
	UserUpdate = models.UserModel{}
	c.Redirect(http.StatusSeeOther, "/adminpannel")
}

func UserSatusHandler(c *gin.Context) {
	c.Header("Cache-control", "no-cache,no-store,must-revalidate")
	var UserUpdate models.UserModel
	session := sessions.Default(c)
	check := session.Get(RoleAdmin)

	if check != nil {
		User := c.Param("ID")
		initializers.DB.First(&UserUpdate, "ID=?", User)

		if UserUpdate.Status == "Active" {
			UserUpdate.Status = "Blocked"
			AdError = "you have successfully Blocked the user"
		} else {
			UserUpdate.Status = "Active"
			AdError = "you have successfully Activated the User"
		}

		initializers.DB.Save(&UserUpdate)
		UserUpdate = models.UserModel{}
		c.Redirect(http.StatusSeeOther, "/adminpannel")
	} else {
		c.Redirect(http.StatusSeeOther, "/admin/login")
	}
}
func UserDeletHandler(c *gin.Context) {
	c.Header("Cache-control", "no-cache,no-store,must-revalidate")
	var UserUpdate models.UserModel
	session := sessions.Default(c)
	check := session.Get(RoleAdmin)

	if check != nil {
		User := c.Param("ID")
		initializers.DB.First(&UserUpdate, "ID=?", User)
		initializers.DB.Delete(&UserUpdate)
		UserUpdate = models.UserModel{}
		AdError = "User Successfully Deleted"
		c.Redirect(http.StatusSeeOther, "/adminpannel")
	} else {
		c.Redirect(http.StatusSeeOther, "/admin/login")
	}
}
func LogoutHandler(c *gin.Context) {
	c.Header("Cache-control", "no-cache,no-store,must-revalidate")
	session := sessions.Default(c)
	session.Delete(RoleAdmin)
	session.Save()
	GetAdmin = models.AdminModel{}
	c.Redirect(http.StatusSeeOther, "/admin/login")
}
