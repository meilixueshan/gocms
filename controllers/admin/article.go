package admin

import (
	"fmt"
	"strconv"
	"encoding/json"
	models "gocms/models"
)

type ArticleController struct {
	baseController
}

func (this *ArticleController) Get() {
	articleCategory := new(models.Articlecategory)
	categories := articleCategory.GetList()
	this.Data["categories"] = categories
	this.TplNames = "admin/article.tpl"
}

func (this *ArticleController) Edit() {
	id := this.Ctx.Input.Param(":id")
	
	articleCategory := new(models.Articlecategory)
	categories := articleCategory.GetList()
	
	this.Data["id"] = id
	this.Data["categories"] = categories
	this.TplNames = "admin/articleEdit.tpl"
}

func (this *ArticleController) Save() {
	ajaxResult := new(models.AjaxResult)
	article := new(models.Article)
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &article)
	if err != nil {
		ajaxResult.Error = err.Error()
	} else {
		err = article.Save()
		if err != nil {
			ajaxResult.Error = err.Error()
		} else {
			ajaxResult.Success = true
			ajaxResult.Msg = &article
		}	
	}
	
	this.Data["json"] = ajaxResult
	this.ServeJson()
}

func (this *ArticleController) GetData() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	article := new(models.Article)
	err := article.GetArticleInfo(id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(article)
	this.Data["json"] = article
	this.ServeJson()
}

func (this *ArticleController) GetDataList() {
	var filter map[string]interface{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &filter)
	if err != nil {
		fmt.Println(err)
	}
	pageSize, err := strconv.Atoi(filter["pageSize"].(string))
	pageNo, err := strconv.Atoi(filter["pageNo"].(string))
	result := map[string]interface{}{"totalCount":0, "data":nil}
	article := new(models.Article)
	totalCount := article.GetCount(filter)
	list := article.GetList(pageSize, pageNo, filter)
	result["totalCount"] = totalCount
	result["data"] = list
	this.Data["json"] = result
	fmt.Println(result)
	this.ServeJson()
}

func (this *ArticleController) Delete() {
	var ajaxResult models.AjaxResult
	id := this.Ctx.Input.Param(":id")
	articleId, err := strconv.Atoi(id)
	if err != nil {
		ajaxResult = models.AjaxResult{Success:false}
		ajaxResult.Error = err.Error()
	} else {
		article := new(models.Article)
		ajaxResult =  article.Delete(articleId)
	}
	
	this.Data["json"] = ajaxResult
	this.ServeJson()
}