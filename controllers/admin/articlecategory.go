package admin

import (
	"strings"
	"fmt"
	"strconv"
	"encoding/json"
	models "gocms/models"
)

type ArtilcecategoryController struct {
	baseController
}

func (this *ArtilcecategoryController) Get() {
	articleCategory := new(models.Articlecategory)
	categories := articleCategory.GetList()
	this.Data["categories"] = categories
	this.TplNames = "admin/articleCategory.tpl"
}

func (this *ArtilcecategoryController) Edit() {
	id := this.Ctx.Input.Param(":id")
	articleCategory := new(models.Articlecategory)
	
	if strings.Compare(id, "") == 0 {
		this.Data["categories"] = articleCategory.GetList()
	} else {
		categoryId, err := strconv.Atoi(id)
		if err !=nil {
			fmt.Println(err)
		}
		categories := articleCategory.GetListEncludesSelf(categoryId)
		fmt.Println(categories)
		this.Data["categories"] = categories
	}

	this.Data["id"] = id
	this.TplNames = "admin/articleCategoryEdit.tpl"
}

func (this *ArtilcecategoryController) Save() {
	ajaxResult := new(models.AjaxResult)
	category := new(models.Articlecategory)
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &category)
	if err != nil {
		ajaxResult.Error = err.Error()
	} else {
		err = category.Save()
		if err != nil {
			ajaxResult.Error = err.Error()
		} else {
			ajaxResult.Success = true
			ajaxResult.Msg = &category
		}	
	}
	
	this.Data["json"] = ajaxResult
	this.ServeJson()
}

func (this *ArtilcecategoryController) GetData() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	category := new(models.Articlecategory)
	err := category.GetCategoryInfo(id)
	if err != nil {
		fmt.Println(err)
	}
	this.Data["json"] = category
	this.ServeJson()
}

func (this *ArtilcecategoryController) GetDataList() {
	articleCategory := new(models.Articlecategory)
	categories := articleCategory.GetList()
	this.Data["json"] = categories
	this.ServeJson()
}

func (this *ArtilcecategoryController) Delete() {
	var ajaxResult models.AjaxResult
	id := this.Ctx.Input.Param(":id")
	categoryId, err := strconv.Atoi(id)
	if err != nil {
		ajaxResult = models.AjaxResult{Success:false}
		ajaxResult.Error = err.Error()
	} else {
		category := new(models.Articlecategory)
		ajaxResult =  category.Delete(categoryId)
	}
	
	this.Data["json"] = ajaxResult
	this.ServeJson()
}