package controllers

import (
	"BeeApi/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/beego/beego/v2/client/orm"
	"strconv"
)

type EventItemCtrl struct {
	beego.Controller
}
// swagger注解配置
// @Title Post
// @Description insert event
// @Success 200 {string} models.EventItem.Id
// @Failure 403 error
// @router / [post]
func (this *EventItemCtrl) AddEvent()  {
	var eventInfo models.EventItem
	var err error
	if err =json.Unmarshal(this.Ctx.Input.RequestBody,&eventInfo); err == nil{
		o := orm.NewOrm()
		id,errI := o.Insert(eventInfo)
		if errI == nil{
			this.Data["json"] = "{\"id\":\"" + strconv.FormatInt(id,10) + "\"}"
		}else{
			this.Data["json"] = errI.Error()
		}
	}else{
		this.Data["json"] = err.Error()
	}
	this.ServeJSON()
}

// swagger注解配置
// @Title Get
// @Description get one event
// @Success 200 {EventItem} models.EventItem
// @Failure 403 error
// @router / [get]
func (this *EventItemCtrl) GetOne(){
	//var id := this.Ctx.Input.RequestBody
	fmt.Print(this.Ctx.Input.RequestBody)
}