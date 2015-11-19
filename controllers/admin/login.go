package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	"gocms/utils"
	models "gocms/models"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	
	ajaxResult := new(models.AjaxResult)
	
	user := new(models.User)
	err := user.GetUserName(username)
	fmt.Println(user)
	if err != nil {
		ajaxResult.Error = "帐号不存在！"
	} else {
		if user.Password != utils.ToMD5([]byte(password)) {
			ajaxResult.Error = "密码输入错误！"
		} else {
			ajaxResult.Success = true
			this.SetSession("adminAuthen", utils.ToJson(user))
		}
	}

	this.Ctx.WriteString(ajaxResult.ToString())
}

func (this *LoginController) Get() {
	this.TplNames = "admin/login.tpl"
}

func (this *LoginController) Logout() {
	this.DelSession("adminAuthen")
	this.Redirect("/admin/login", 302)
}