package front

import (
	models "gocms/models"
	"github.com/astaxie/beego"
	"strconv"
	"fmt"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	filter := map[string]interface{}{"categoryId":"", "keywords": ""}	//初始化一个无用定义，防止filter为nil

	article := new(models.Article)
	list := article.GetList(15, 1, filter)
	
	this.Data["Website"] = "gocms"
	this.Data["Email"] = "44476511@qq.com"
	this.Data["articles"] = list
	this.TplNames = "front/index.tpl"
}

func (this *HomeController) ArticleList() {
	cid := this.Ctx.Input.Param(":cid")
	filter := map[string]interface{}{"categoryId": cid, "keywords": ""}	//初始化一个无用定义，防止filter为nil

	article := new(models.Article)
	list := article.GetList(15, 1, filter)

	this.Data["articles"] = list
	this.TplNames = "front/articlelist.tpl"
}

func (this *HomeController) ArticleDetail() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	article := new(models.Article)
	err := article.GetArticleInfo(id)
	if err != nil {
		fmt.Println(err)
	}
	this.Data["article"] = article
	
	this.TplNames = "front/articledetail.tpl"
}

func (this *HomeController) PageDetail() {
	id := this.Ctx.Input.Param(":id")
	page := new(models.Page)
	err := page.GetPageInfo(id)
	if err != nil {
		fmt.Println(err)
	}
	this.Data["page"] = page
	this.TplNames = "front/pagedetail.tpl"
}