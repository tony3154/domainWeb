package controllers

import (
	"DomainDiscoverWeb/models"
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type RegisterController struct {
	beego.Controller
}

// 注册展示页面
func (this *RegisterController) ShowRegister() {
	this.TplName = "register.html"
}

// 注册获取数据页面
func (this *RegisterController) HandleRegister() {
	// 获取浏览器传递的值，并去除两边的空格

	Name := strings.TrimSpace(this.GetString("userName"))
	Pwd := strings.TrimSpace(this.GetString("passWord"))
	// beego.Info("账号:", Name, "密码:", Pwd)
	// 数据处理
	if Name == "" || Pwd == "" {
		beego.Info("用户名或密码不能为空")
		this.TplName = "register.html"
		this.Data["errmsg"] = "用户名或密码不能为空 ！"
		return
	}
	// 插入数据库（数据库表，Users）
	//获取orm对象
	o := orm.NewOrm()
	//   获取插入对象
	user := models.Users{}
	//   插入数值
	user.Name = Name
	user.Pwd = Pwd

	_, err := o.Insert(&user)
	if err != nil {
		beego.Info("插入数据失败，用户相同或者其他错误！！！")
		this.TplName = "register.html"
		this.Data["errmsg"] = "插入数据失败，用户相同或者其他错误！！！！"
		return
	}
	// 测试返回视图
	// this.Ctx.WriteString("注册成功！！！")
	// 实际情况注册成功返回到登录页面
	this.Redirect("login", 302)
}

type LoginController struct {
	beego.Controller
}

// 登录页面 get
func (this *LoginController) ShowLogin() {
	this.TplName = "login.html"
}

// 登录页面 post
func (this *LoginController) HandleLogin() {
	// 拿到浏览器数据，并去除两边的空格
	Name := strings.TrimSpace(this.GetString("userName"))
	Pwd := strings.TrimSpace(this.GetString("passWord"))
	fmt.Println(Name, Pwd)
	beego.Info("账号:", Name, "密码:", Pwd)

	//数据处理
	if Name == "" || Pwd == "" {
		beego.Info("登录失败！！")
		this.TplName = "login.html"
		this.Data["errmsg"] = "登录失败！！！！！"
		return
	}
	// 查找数据库
	//获取orm对象
	o := orm.NewOrm()
	//获取查询对象
	var user models.Users
	// 查询
	user.Name = Name
	err := o.Read(&user, "Name")
	if err != nil {
		beego.Info("用户名登录失败！！！")
		this.TplName = "login.html"
		this.Data["errmsg"] = "用户名登录失败！！！！！"
		return
	}
	// 判断密码是否一致
	if user.Pwd != Pwd {
		beego.Info("密码登录失败！！！")
		this.TplName = "login.html"
		this.Data["errmsg"] = "密码登录失败！！"
		return
	}
	// 测试返回视图
	this.Ctx.WriteString("登录成功")
	// 实际情况注册成功返回到登录页面
	// this.Redirect("index", 302)
}
