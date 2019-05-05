package controllers

import (
	"github.com/astaxie/beego"
	"github.com/proj_LP/models"
)

type MainController struct {
	beego.Controller
	usr models.User
}

func (main *MainController) Get() {
	main.TplName = "index.html"
	main.Data["prof_pic"] = models.RequestUserProfilePicture(1)
}
