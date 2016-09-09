package controllers

import (
	"github.com/astaxie/beego"
	"testBeego/models"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"testBeego/databases/mongo"
	"log"
	"testBeego/response"
	"github.com/dgrijalva/jwt-go"
	"testBeego/databases/redis"
	"fmt"
	red "github.com/garyburd/redigo/redis"
)

// operations for Test
type TestController struct {
	beego.Controller
}

func (c *TestController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Create
// @Description create Test
// @Param	body		body 	models.Test	true		"body for Test content"
// @Success 201 {object} models.Test
// @Failure 403 body is empty
// @router / [post]
func (c *TestController) Post() {

}

// @Title GetOne
// @Description get Test by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Test
// @Failure 403 :id is empty
// @router /:id [get]
func (c *TestController) GetOne() {
	id := c.GetString(":id")
	if id != "" {
		user := models.GetPersonById(bson.M{"_id": bson.ObjectIdHex(id)})
		c.Data["json"] = user
	}
	c.ServeJSON()
}

// @Title GetAll
// @Description get Test
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Test
// @Failure 403
// @router / [get]
func (c *TestController) GetAll() {
	users := models.GetAllPerson(nil, "email")
	c.Data["json"] = users
	c.ServeJSON()
}

// @Title Update
// @Description update the Test
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Test	true		"body for Test content"
// @Success 200 {object} models.Test
// @Failure 403 :id is not int
// @router /:id [put]
func (c *TestController) Put() {
	reqToken := c.GetString(":id")

	// connect redis
	session := redis.Conn()
	defer session.Close()

	m, err := red.String(session.Do("GET", reqToken))
	if err != nil {
		log.Println(err)
	}

	if m == "" {


	} else {

		// json string to struct
		var structM models.Person
		json.Unmarshal([]byte(m), &structM)

		// get body
		var p models.Person
		json.Unmarshal(c.Ctx.Input.RequestBody, &p)
	}
}

// @Title Delete
// @Description delete the Test
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *TestController) Delete() {

}

func (c *TestController) Register() {

	var p models.Person
	json.Unmarshal(c.Ctx.Input.RequestBody, &p)

	// connect mongo
	session := mongo.Conn()
	defer session.Close()

	result := models.Person{}
	coll := session.DB("test").C("person")
	err := coll.Find(bson.M{"email": p.Email}).One(&result)
	if err != nil {
		log.Println(err)
	}

	res := response.Response{}
	if p.Email == result.Email {
		res.Success = false
		res.Desc = "duplicate email"
	} else {
		// insert
		coll.Insert(&p)

		res.Success = true
		res.Desc = "success"
	}

	// response
	c.Data["json"] = res
	c.ServeJSON()
}

func (c *TestController) Login() {

	var p models.Person
	json.Unmarshal(c.Ctx.Input.RequestBody, &p)

	// check duplicate email
	session := mongo.Conn()
	defer session.Close()

	// find document by email & password
	pColl := models.Person{}
	coll := session.DB("test").C("person")
	err := coll.Find(bson.M{"email": p.Email, "password": p.Password}).One(&pColl)
	if err != nil {
		log.Println(err)
	}

	res := response.Response{}
	if pColl.Email == "" {
		res.Success = false
		res.Desc = "not found"
	} else {

		// create token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"ID": pColl.Id})
		tokenString, err := token.SignedString([]byte("AllYourBase"))
		if err != nil {
			log.Println(err)
		}

		// save token to redis
		session := redis.Conn()
		defer session.Close()

		jsonMember, _ := json.Marshal(pColl)
		session.Do("SET", tokenString, jsonMember)
		session.Do("EXPIRE", tokenString, 20)

		c.Ctx.Output.Header("token", tokenString)
		fmt.Println(tokenString)

		res.Success = true
		res.Desc = "success"
	}

	c.Data["json"] = res
	c.ServeJSON()
}