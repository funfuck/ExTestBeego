package controllers

import (
	"github.com/astaxie/beego"
	_ "testBeego/models"
	_ "fmt"
	"testBeego/models"
	_ "gopkg.in/mgo.v2/bson"
)

type PersonController struct {
	beego.Controller
}

func (p *PersonController) GetAll() {
	users := models.GetAllPerson(nil, "email")
	p.Data["json"] = users
	p.ServeJSON()
}

func (p *PersonController) Get() {
	users := models.GetPersonById(nil)
	p.Data["json"] = users
	p.ServeJSON()
}