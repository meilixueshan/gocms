package admin

import (
	"fmt"
	models "gocms/models"
	"encoding/json"
	"gocms/utils"
)

type HomeController struct {
	baseController
}

func (this *HomeController) Get() {
	this.TplNames = "admin/index.tpl"
}

func (this *HomeController) Password() {
	this.TplNames = "admin/password.tpl"
}

func (this *HomeController) SavePassword() {
	var ajaxResult models.AjaxResult
	var data map[string]interface{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &data)
	if err != nil {
		fmt.Println(err)
	}
	
	oldpwd := data["oldpwd"].(string)
	newpwd := data["newpwd"].(string)

	if utils.ToMD5([]byte(oldpwd)) != this.User.Password {
		ajaxResult.Success = false
		ajaxResult.Error = "原密码输入错误！"
	} else {
		user := new(models.User)
		user.Id = this.User.Id
		user.Username = this.UserName
		user.Password = utils.ToMD5([]byte(newpwd))
		ajaxResult = user.SaveUserInfo()
	}

	this.Data["json"] = ajaxResult
	this.ServeJson()
}
