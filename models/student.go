package models

type Student struct {
	Id int
	Name string
	Profile *Profile	`orm:"rel(one)"`	//1对1
	Post []*Post `orm:"reverse(many)"`	//设置1对多的反向关系
}

type Profile struct {
	Id int
	Age int16
	User *Student	`orm:"reverse(one)"`	//设置1对1反向关系(可选)
}

type Post struct {
	Id int
	Title string
	User *Student	`orm:"rel(fk)"`		//设置1对多关系
	Tags []*Tag	`orm:"rel(m2m)"`	//设置多对多关系
}

type Tag struct {
	Id int
	Name string
	Posts []*Post	`orm:"reverse(many)"`	//设置多对多反向关系
}

func init(){
	//在init中注册定义的model
	//orm.RegisterModel(new(Student),new(Post),new(Profile),new(Tag))
}
