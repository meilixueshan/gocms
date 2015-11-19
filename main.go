package main
import (
	"fmt"
	"gocms/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "gocms/routers"
)

func init() {
	dataSource := beego.AppConfig.String("datasource")
	maxIdleConn, err1 := beego.AppConfig.Int("maxidleconn")
	if err1 != nil {
		maxIdleConn = 5
	}
	
	maxOpenConn, err2 := beego.AppConfig.Int("maxopenconn")
	if err2 != nil {
		maxOpenConn = 10
	}
	
	orm.Debug = true	//开启调试，控制台将会输出SQL语句
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	err := orm.RegisterDataBase("default", "mysql", dataSource, maxIdleConn, maxOpenConn)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	beego.SessionOn = true
	beego.SetStaticPath("/public", "public")
	//增加模板格式化函数标签
	beego.AddFuncMap("CategoryFormat", utils.CategoryFormat)
	beego.Run()
}