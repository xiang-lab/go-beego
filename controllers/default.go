package controllers

import (
	"class/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	//"github.com/astaxie/beego/orm"
	//"class/models"
)

type MainController struct {
	beego.Controller
}

//func (c *MainController) Get() {
//	c.Data["Website"] = "beego.me"
//	c.Data["Email"] = "astaxie@gmail.com"
//	c.TplName = "index.tpl"
//}

func (c *MainController) Get() {

	/* 数据的插入 */
	///* 1. ORM对象 */
	//o := orm.NewOrm()
	///* 2. 插入数据的结构体对象 */
	//user := models.User{}
	///* 3. 对结构体赋值 */
	//user.Name = "111"
	//user.Pwd = "222"
	///* 4. 插入 */
	//_, err := o.Insert(&user)
	//if err != nil {
	//	beego.Info("插入失败", err)
	//	return
	//}


	/* 查询 */
	///* 1. ORM对象 */
	//o := orm.NewOrm()
	///* 2. 查询对象 */
	//user := models.User{}
	///* 3. 指定查询对象字段值 */
	//user.Name = "111"
	///* 4. 查询 */
	//err :=  o.Read(&user, "Name")
	//if err != nil {
	//	beego.Info("查询失败")
	//	return
	//}
	//beego.Info("查询成功", user)


	/* 更新 */
	///* 1. ORM对象 */
	// o := orm.NewOrm()
	///* 2. 需要跟新的结构对象 */
	//user := models.User{}
	///* 3. 查找到需要更新的数据 */
	//user.Id = 1
	//err := o.Read(&user)
	///* 4. 给数据重新赋值 */
	//if err == nil {
	//	user.Name = "222"
	//	user.Pwd = "333"
	//	/* 5. 更新 */
	//	_, err := o.Update(&user)
	//	if err != nil {
	//		beego.Info("更新失败", err)
	//		return
	//	}
	//}


	/* 删除 */
	///* 1. ORM对象 */
	//o := orm.NewOrm()
	///* 2. 需要删除的结构体 */
	//user := models.User{}
	///* 3. 指定删除的数据 */
	//user.Id = 1
	///* 4. 删除 */
	//_, err := o.Delete(&user)
	//if err != nil {
	//	beego.Info("删除错误", err)
	//	return
	//}


	//c.Data["data"] = "今天中午吃饺子"
	//c.TplName = "test.html"

	c.TplName = "register.html"
}

func (c *MainController) Post() {
	/* 1. 拿到数据 */
	userName := c.GetString("userName")
	pwd := c.GetString("pwd")
	/* 2. 对数据进行校验 */
	if userName == "" || pwd == "" {
		beego.Info("数据不能为空")
		c.Redirect("/register", 302)
		return
	}
	/* 3. 插入数据库 */
	o := orm.NewOrm()
	user := models.User{}
	user.Name = userName
	user.Pwd = pwd
	_, err := o.Insert(&user)
	if err != nil {
		beego.Info("插入数据失败", err)
		return
	}
	/* 4. 返回登陆界面 */
	c.TplName = "login.html"
}

func (c *MainController) ShowLogin() {
	c.TplName = "login.html"
}

/* 登陆业务处理 */
func (c *MainController) HandleLogin() {

	/* 1. 获取数据 */
 	userName := c.GetString("userName")
	pwd :=  c.GetString("pwd")
	/* 2. 判断数据是否合法 */
	if userName == "" || pwd == "" {
		beego.Info("数据不合法")
		c.TplName = "login.html"
		return
	}
	/* 3. 查询账号密码的正确性 */
	o := orm.NewOrm()
	user := models.User{}
	user.Name = userName
	err := o.Read(&user, "Name")
	if err != nil {
		beego.Info("用户名不正确")
		c.TplName = "login.html"
		return
	}
	if user.Pwd != pwd {
		beego.Info("密码错误")
		c.TplName = "login.html"
		return
	}
	/* 4. 跳转 */
	c.Redirect("/index", 302)
}

/* 显示首页内容 */
func (c *MainController) ShowIndex() {
	c.TplName = "index.html"
}