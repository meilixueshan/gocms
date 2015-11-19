package models

import (
	"encoding/json"
	"fmt"
	"gocms/utils"
	"github.com/astaxie/beego/orm"
)

type Site struct {
	Id      int
	Content string
}

func (this *Site) TableName() string {
	return "site"
}

func init() {
	orm.RegisterModel(new(Site))
}

func (this *Site) GetSiteInfo() map[string]interface{} {
	jsonMap := make(map[string]interface{})

	o := orm.NewOrm()
	site := Site{Id: 1}
	err := o.Read(&site)
	if err == nil {
		err = json.Unmarshal([]byte(site.Content), &jsonMap)
	} else {
		fmt.Println(err)
	}
	return jsonMap
}

func (this *Site) SaveSiteInfo(siteInfo map[string]interface{}) AjaxResult {
	ajaxResult := new(AjaxResult)

	content := utils.ToJson(siteInfo)

	o := orm.NewOrm()
	site := Site{Id: 1, Content: content}

	_, err := o.Update(&site)
	if err != nil {
		ajaxResult.Success = false
		ajaxResult.Error = err.Error()
	} else {
		ajaxResult.Success = true
	}
	return *ajaxResult
}
