package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Page struct {
	Id             string `orm:"pk;column(id);"`
	Title          string
	Content        string
	Seotitle       string
	Seokeywords    string
	Seodescription string
}

func (this *Page) TableName() string {
	return "page"
}

func init() {
	orm.RegisterModel(new(Page))
}

func (this *Page) GetPageInfo(id string) (err error) {
	o := orm.NewOrm()
	err = o.QueryTable(this).Filter("id", id).One(this)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func (this *Page) GetList() []orm.Params {
	var maps []orm.Params
	o := orm.NewOrm()
	_, err := o.Raw("select id, title, seotitle from page order by title").Values(&maps)
	if err != nil {
		fmt.Println(err)
	}
	return maps
}

func (this *Page) Save() error {
	var err error
	o := orm.NewOrm()
	page := new(Page)
	err = page.GetPageInfo(this.Id)
	if err != nil {
		_, err = o.Insert(this)
	} else {
		_, err = o.Update(this)
	}
	return err
}

func (this *Page) Delete(id string) AjaxResult {
	ajaxResult := new(AjaxResult)
	o := orm.NewOrm()
	_, err := o.Delete(&Page{Id: id})
	if err != nil {
		ajaxResult.Success = false
		ajaxResult.Error = err.Error()
	} else {
		ajaxResult.Success = true
	}
	return *ajaxResult
}
