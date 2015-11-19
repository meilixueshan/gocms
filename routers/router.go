package routers

import (
	"github.com/astaxie/beego"
	"gocms/controllers/admin"
	"gocms/controllers/front"
)

func init() {
	beego.Router("/admin", &admin.HomeController{})
	beego.Router("/admin/password", &admin.HomeController{}, "get:Password")
	beego.Router("/admin/password/save", &admin.HomeController{}, "post:SavePassword")
	beego.Router("/admin/login", &admin.LoginController{})
	beego.Router("/admin/logout", &admin.LoginController{}, "get:Logout")
	beego.Router("/admin/site", &admin.SiteController{})
	beego.Router("/admin/site/getdata", &admin.SiteController{}, "get:GetData")
	beego.Router("/admin/site/save", &admin.SiteController{}, "post:Save")
	
	beego.Router("/admin/page", &admin.PageController{})
	beego.Router("/admin/page/edit/:id([0-9a-zA-Z]*)", &admin.PageController{}, "get:Edit")
	beego.Router("/admin/page/save", &admin.PageController{}, "post:Save")
	beego.Router("/admin/page/delete/:id([0-9a-zA-Z]+)", &admin.PageController{}, "get:Delete")
	beego.Router("/admin/page/getdata/:id([0-9a-zA-Z]+)", &admin.PageController{}, "get:GetData")
	beego.Router("/admin/page/getdatalist", &admin.PageController{}, "get:GetDataList")

	beego.Router("/admin/article", &admin.ArticleController{})
	beego.Router("/admin/article/edit/:id([0-9]*)", &admin.ArticleController{}, "get:Edit")
	beego.Router("/admin/article/save", &admin.ArticleController{}, "post:Save")
	beego.Router("/admin/article/delete/:id([0-9]+)", &admin.ArticleController{}, "get:Delete")
	beego.Router("/admin/article/getdata/:id([0-9]+)", &admin.ArticleController{}, "get:GetData")
	beego.Router("/admin/article/getdatalist", &admin.ArticleController{}, "post:GetDataList")
	
	beego.Router("/admin/articlecategory", &admin.ArtilcecategoryController{})
	beego.Router("/admin/articlecategory/edit/:id([0-9]*)", &admin.ArtilcecategoryController{}, "get:Edit")
	beego.Router("/admin/articlecategory/save", &admin.ArtilcecategoryController{}, "post:Save")
	beego.Router("/admin/articlecategory/delete/:id([0-9]+)", &admin.ArtilcecategoryController{}, "get:Delete")
	beego.Router("/admin/articlecategory/getdata/:id([0-9]+)", &admin.ArtilcecategoryController{}, "get:GetData")
	beego.Router("/admin/articlecategory/getdatalist", &admin.ArtilcecategoryController{}, "post:GetDataList")
	
	beego.Router("/", &front.HomeController{})
	beego.Router("/article/list/:cid([0-9]*)", &front.HomeController{}, "get:ArticleList")
	beego.Router("/article/show/:id([0-9]+)", &front.HomeController{}, "get:ArticleDetail")
	beego.Router("/page/:id([0-9a-zA-Z]+)", &front.HomeController{}, "get:PageDetail")
	//建立一个配置文件，动态配置路由
}