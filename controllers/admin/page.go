package admin

import (
	"fmt"
	models "gocms/models"
	"encoding/json"
)

type PageController struct {
	baseController
}

func (this *PageController) Get() {
	this.TplNames = "admin/page.tpl"
}

func (this *PageController) Edit() {
	id := this.Ctx.Input.Param(":id")
	this.Data["id"] = id
	this.TplNames = "admin/pageEdit.tpl"
}

func (this *PageController) Save() {
	ajaxResult := new(models.AjaxResult)
	page := new(models.Page)
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &page)
	if err != nil {
		ajaxResult.Error = err.Error()
	} else {
		err = page.Save()
		if err != nil {
			ajaxResult.Error = err.Error()
		} else {
			ajaxResult.Success = true
			ajaxResult.Msg = &page
		}	
	}
	
	this.Data["json"] = ajaxResult
	this.ServeJson()
}

func (this *PageController) Delete() {
	id := this.Ctx.Input.Param(":id")
	page := new(models.Page)
	ajaxResult := page.Delete(id)
	this.Data["json"] = ajaxResult
	this.ServeJson()
}

func (this *PageController) GetData() {
	id := this.Ctx.Input.Param(":id")
	page := new(models.Page)
	err := page.GetPageInfo(id)
	if err != nil {
		fmt.Println(err)
	}

	this.Data["json"] = page
	this.ServeJson()
}

func (this *PageController) GetDataList() {
	page := new(models.Page)
	list := page.GetList()
	this.Data["json"] = list
	this.ServeJson()
}