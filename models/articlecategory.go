package models

import (
	"fmt"
	"gocms/utils"
	"github.com/astaxie/beego/orm"
)

type Articlecategory struct {
	Id             int
	Categoryname   string
	Parentid       int
	Sortnum        int
	Innercode      string
	Seotitle       string
	Seokeywords    string
	Seodescription string
}

func (this *Articlecategory) TableName() string {
	return "articlecategory"
}

func init() {
	orm.RegisterModel(new(Articlecategory))
}

func (this *Articlecategory) GetList() ([]*Articlecategory) {
	var list []*Articlecategory
	o:=orm.NewOrm()
	_, err := o.QueryTable("Articlecategory").OrderBy("innercode").All(&list)
	if err != nil {
		fmt.Println(err)
	}
	return list
}

func (this *Articlecategory) GetListEncludesSelf(id int) ([]Articlecategory) {
	var categories []Articlecategory
	//注意：查询语句的结果字段要和数据库表里的字段顺序相同，大小写完全相同（也可能得全部是小写）
	sql := "SELECT id,categoryname,parentid,sortnum,innercode,seotitle,seokeywords,seodescription FROM articlecategory WHERE Id NOT IN (SELECT b.id FROM articlecategory a INNER JOIN articlecategory b ON a.id=? AND b.innercode LIKE CONCAT(a.innercode, '%')) order by innercode"
	o := orm.NewOrm()
	_, err := o.Raw(sql, id).QueryRows(&categories)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(categories)

	return categories
}

func (this *Articlecategory) GetCategoryInfo(id int) (err error) {
	o := orm.NewOrm()
	err = o.QueryTable(this).Filter("id", id).One(this)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func (this *Articlecategory) Save() error {
	var err error
	o := orm.NewOrm()
	category := new(Articlecategory)
	err = category.GetCategoryInfo(this.Id)
	if err != nil {
		_, err = o.Insert(this)
	} else {
		_, err = o.Update(this)
	}
	
	this.ResetInnerCode(0, "")	//重置InnerCode
	
	return err
}

func (this *Articlecategory) Delete(id int) AjaxResult {
	ajaxResult := new(AjaxResult)
	o := orm.NewOrm()
	sql := "DELETE FROM articlecategory WHERE id IN (SELECT id FROM (SELECT b.* FROM articlecategory a INNER JOIN articlecategory b ON a.id=? AND b.innercode LIKE CONCAT(a.innercode, '%')) v)"
	r := o.Raw(sql, id)
	_, err := r.Exec()
	if err != nil {
		ajaxResult.Success = false
		ajaxResult.Error = err.Error()
	} else {
		ajaxResult.Success = true
		this.ResetInnerCode(0, "")	//重置InnerCode
	}
	
	return *ajaxResult
}

//重置InnerCode
func (this *Articlecategory) ResetInnerCode(parentId int, parentInnerCode string) {
	var categories []Articlecategory
	//注意：查询语句的结果字段要和数据库表里的字段顺序相同，大小写完全相同（也可能得全部是小写）
	sql := "SELECT id,categoryname,parentid,sortnum,innercode,seotitle,seokeywords,seodescription FROM articlecategory WHERE parentid=? order by sortnum"
	o := orm.NewOrm()
	_, err := o.Raw(sql, parentId).QueryRows(&categories)
	if err != nil {
		fmt.Println(err)
	} else {
		for i, v := range categories {
			code := utils.NumberFormat(i+1, "000")
			innerCode := code
			if len(parentInnerCode) > 0 {
				innerCode = parentInnerCode + "." + code
			}
			
			sqlChild := "update articleCategory set innerCode=? where id=?"
			_, err := o.Raw(sqlChild, innerCode, v.Id).Exec()
			if err != nil {
				fmt.Println(err)
			} else {
				this.ResetInnerCode(v.Id, innerCode)
			}
		}
	}
}