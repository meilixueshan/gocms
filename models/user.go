package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id int
	Username string	//第一个字母大写，n不能大写，否则会被翻译为user_name
	Password string
}

func (u *User) TableName() string {
    return "user"
}


func init(){
    orm.RegisterModel(new(User))
}

func (this *User) GetUserName(userName string) (err error) {
	o := orm.NewOrm()
	err = o.QueryTable(this).Filter("username", userName).One(this)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func (this *User) SaveUserInfo() AjaxResult {
	ajaxResult := new(AjaxResult)
	o := orm.NewOrm()
	_, err := o.Update(this)
	fmt.Println(&this)
	if err != nil {
		ajaxResult.Success = false
		ajaxResult.Error = err.Error()
	} else {
		ajaxResult.Success = true
	}
	return *ajaxResult
}