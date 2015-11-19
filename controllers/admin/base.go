package admin

import (
	"github.com/astaxie/beego"
	"gocms/models"
	"encoding/json"
	"fmt"
)

type AuthenPreparer interface {
	AuthenPrepare()
}

type baseController struct {
	beego.Controller
	UserName string
	User *models.User
}

func (this *baseController) Prepare() {
	this.authen()
	
	this.Data["appname"] = beego.AppConfig.String("appname")
	this.Data["user"] = this.User
	
	if app, ok := this.AppController.(AuthenPreparer); ok {
		app.AuthenPrepare()
	}
}

//验证登录身份
func (this *baseController) authen() {
	v := this.Ctx.Input.CruSession.Get("adminAuthen")
	if v == nil {
		this.Redirect("/admin/login", 302)
		return
	} else {
		user := new(models.User)
		str := v.(string)
		err := json.Unmarshal([]byte(str), &user)
		if err != nil {
		    fmt.Println("error:", err)
		}
		this.User = user
		this.UserName = user.Username
	}
}
