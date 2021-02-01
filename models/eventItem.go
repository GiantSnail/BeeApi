package models

import "github.com/beego/beego/v2/client/orm"

type DatePage struct {
	Id int
	Date int
	EventItem []*EventItem	`orm:"reverse(many)"`
}

type EventItem struct {
	Id int
	Title string
	Content string
	Status string
	DatePage *DatePage	`orm:"rel(fk)"`
}


func init(){
	orm.RegisterModel(new(DatePage),new(EventItem))
}


