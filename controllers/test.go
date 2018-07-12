package controllers

import (
	"encoding/json"

	"github.com/lflxp/boltDb/models"
)

// Operations about object
type TestController struct {
	BaseController
}

// @Title Create
// @Description create object
// @Param	body		body 	models.Test	true		"The object content"
// @Success 200 {string} models.Test.Table
// @Failure 403 body is empty
// @router / [post]
func (o *TestController) Post() {
	var ob models.Test
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	err := ob.AddTest()
	if err != nil {
		o.Data["json"] = map[string]string{"ObjectId": err.Error()}
	} else {
		o.Data["json"] = "ok"
	}
	o.ServeJSON()
}

// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (o *TestController) Get() {
	objectId := o.Ctx.Input.Param(":objectId")
	if objectId != "" {
		data := models.Test{Table: objectId}
		ob, err := data.GetTest()
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (o *TestController) GetAll() {
	obs := models.GetAll()
	o.Data["json"] = obs
	o.ServeJSON()
}

// @Title Update
// @Description update the object
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (o *TestController) Put() {
	objectId := o.Ctx.Input.Param(":objectId")
	var ob models.Object
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	err := models.Update(objectId, ob.Score)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the object
// @Param	objectId		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (o *TestController) Delete() {
	objectId := o.Ctx.Input.Param(":objectId")
	models.Delete(objectId)
	o.Data["json"] = "delete success!"
	o.ServeJSON()
}
