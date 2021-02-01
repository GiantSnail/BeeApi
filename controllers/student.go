package controllers
//
//import (
//	"BeeApi/models"
//	"fmt"
//	"github.com/astaxie/beego"
//	"github.com/beego/beego/v2/client/orm"
//)
//
//type StudentController struct {
//	beego.Controller
//}
//
//func (s *StudentController) Post(){
//	o := orm.NewOrm()
//	profile := new(models.Profile)
//	profile.Age = 30
//
//	stu := new(models.Student)
//	stu.Profile = profile
//	stu.Name = "misaka mikoto"
//
//	_,errP := o.Insert(profile)
//	if errP != nil {
//		fmt.Print(errP)
//	}
//	_,errS := o.Insert(stu)
//	if errS != nil {
//		fmt.Print(errS)
//	}
//}
////@router
//func (s *StudentController) Get(){
//	//o := orm.NewOrm()
//	//profile := new(models.Profile)
//	//profile.Age = 30
//	//
//	//stu := new(models.Student)
//	//stu.Profile = profile
//	//stu.Name = "misaka mikoto"
//	//
//	//_,errP := o.Insert(profile)
//	//if errP != nil {
//	//	//fmt.Print(errP)
//	//	s.Data["json"] = errP.Error()
//	//}
//	//_,errS := o.Insert(stu)
//	//if errS != nil {
//	//	//fmt.Print(errS)
//	//	s.Data["json"] = errS.Error()
//	//}
//	o := orm.NewOrm()
//	tar := models.Student{Id: 1}
//	err := o.Read(&tar)
//	if err != nil {
//		//fmt.Print(err)
//		s.Data["json"] = err.Error()
//	}else{
//		s.Data["json"] = tar
//	}
//	s.ServeJSON()
//}
