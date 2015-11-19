package models

import (
	"bytes"
	"fmt"
	"gocms/utils"
	"time"
	"strconv"
	"github.com/astaxie/beego/orm"
)

type Article struct {
	Id             int `orm:"pk;column(id);"`
	Categoryid     int
	Title          string
	Description    string
	Author         string
	Comefrom       string
	Keywords       string
	Content        string
	Publishtime    string
	Createtime     string
	Readnum        int
	Picfilename    string
	Seotitle       string
	Seokeywords    string
	Seodescription string
}

func (this *Article) TableName() string {
	return "article"
}

//init初始化函数，字母必须全部是小写
func init() {
	orm.RegisterModel(new(Article))
}

func (this *Article) GetList(pageSize int, pageNo int,
	filter map[string]interface{}) []orm.Params {
	var params []interface{}
	var sql bytes.Buffer
	var list []orm.Params
	firstResult := (pageNo - 1) * pageSize

	categoryId := filter["categoryId"]
	keywords := filter["keywords"]

	sql.WriteString("select t.id, t.title, t.categoryid, c.categoryname, t.description, t.publishtime, t.createtime, t.readnum, t.author from article t, articlecategory c where t.categoryId=c.id")

	if categoryId != "" {
		sql.WriteString(" and categoryId in (SELECT b.id FROM articlecategory a INNER JOIN articlecategory b ON a.id=? AND b.innercode LIKE CONCAT(a.innercode, '%'))")
		params = append(params, categoryId)
	}

	if keywords != "" {
		sql.WriteString(" and t.title like ?")
		params = append(params, utils.Concat("%", keywords.(string), "%"))
	}

	sql.WriteString(" order by t.id desc limit ?, ?")
	params = append(params, firstResult)
	params = append(params, pageSize)

	o := orm.NewOrm()
	_, err := o.Raw(sql.String(), params).Values(&list)
	if err != nil {
		fmt.Println(err)
	}

	return list
}

func (this *Article) GetCount(filter map[string]interface{}) int {
	var params []interface{}
	var maps []orm.Params
	var sql bytes.Buffer
	totalCount := 0

	categoryId := filter["categoryId"]
	keywords := filter["keywords"]

	sql.WriteString("select count(*) as totalcount from article where 1=1")

	if categoryId != "" {
		sql.WriteString(" and categoryId in (SELECT b.id FROM articlecategory a INNER JOIN articlecategory b ON a.id=? AND b.innercode LIKE CONCAT(a.innercode, '%'))")
		params = append(params, categoryId)
	}

	if keywords != "" {
		sql.WriteString(" and title like ?")
		params = append(params, utils.Concat("%", keywords.(string), "%"))
	}

	o := orm.NewOrm()
	_, err := o.Raw(sql.String(), params).Values(&maps)
	if err != nil {
		fmt.Println(err)
	} else {
		totalCount, _ = strconv.Atoi(maps[0]["totalcount"].(string))
	}

	return totalCount
}

func (this *Article) GetArticleInfo(id int) (err error) {
	o := orm.NewOrm()
	err = o.QueryTable(this).Filter("id", id).One(this)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func (this *Article) Save() error {
	var err error
	o := orm.NewOrm()
	article := new(Article)
	err = article.GetArticleInfo(this.Id)
	if err != nil {
		this.Createtime = utils.TimeFormat(time.Now(), "yyyy-mm-dd hh:ii:ss")
		_, err = o.Insert(this)
	} else {
		_, err = o.Update(this)
	}
	return err
}

func (this *Article) Delete(id int) AjaxResult {
	ajaxResult := new(AjaxResult)
	o := orm.NewOrm()
	_, err := o.Delete(&Article{Id: id})
	if err != nil {
		ajaxResult.Success = false
		ajaxResult.Error = err.Error()
	} else {
		ajaxResult.Success = true
	}
	return *ajaxResult
}
