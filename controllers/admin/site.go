package admin

import (
	"fmt"
	models "gocms/models"
	"encoding/json"
)

type SiteController struct {
	baseController
}

func (this *SiteController) Get() {
	this.TplNames = "admin/site.tpl"
}

func (this *SiteController) GetData() {
	site := new(models.Site)
	siteInfo := site.GetSiteInfo()
	this.Data["json"] = siteInfo
	this.ServeJson()	//输出json（this.Data的键名必须是json）
}

func (this *SiteController) Save() {
	//this.Ctx.Input.Param(":DomainName")
	//this.Input().Get("DomainName")
	var siteInfo map[string]interface{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &siteInfo)
	if err != nil {
		fmt.Println(err)
	}
	
	site := new(models.Site)
	this.Data["json"] = site.SaveSiteInfo(siteInfo)
	this.ServeJson()	//输出json（this.Data的键名必须是json）
}