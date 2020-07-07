package models

import "github.com/astaxie/beego/orm"

func init() {
	//设置数据库连接信息
	orm.RegisterDataBase("default", "mysql", "go:Asong123!@#@tcp(192.168.182.200:3306)/Beego?charset=utf8&loc=Local", 30)
	// 映射modle数据
	orm.RegisterModel(new(Users))
	orm.RegisterModel(new(Domain))
	// 生成表，第二个false要是改成true 就会强制更新表，数据全部丢失
	orm.RunSyncdb("default", false, true)

}
