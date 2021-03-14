package controller

import (
	"forum_Anpw/common"
	"forum_Anpw/model"
	"forum_Anpw/reps"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)
//用户注册
func Register(c *gin.Context)  {

	db:=common.GetDB()
	username:=c.PostForm("username")
	password:=c.PostForm("password")
	securityCode:=c.PostForm("securitycode")

	//检查
	if len(username)==0 {
		reps.UnprocessableEntity(c,nil,"用户名不能为空")
		return
	}

	//判断用户名是否已经存在
	if isUsernameExist(db,username) {
		reps.UnprocessableEntity(c,nil,"用户名已存在")
		return
	}
	if len(password)<6 {
		reps.UnprocessableEntity(c,nil,"密码不能小于6位")
		return
	}
	if len(securityCode)==0 {
		reps.UnprocessableEntity(c,nil,"安全码不能为空")
		return
	}
	if len(securityCode)>=6 {
		reps.UnprocessableEntity(c,nil,"安全码不能超过6位")
		return

	}


	//创建用户

	//加密一波
	hasedPassword,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err!=nil {
		reps.ServerError(c,nil,"加密错误")
		return
	}

	newUser :=model.User{
		Username:username,
		Password:string(hasedPassword),
		SequrityCode: securityCode,
	}
	db.Create(&newUser)
	reps.Ok(c,nil,"注册成功")
}

func isUsernameExist(db *gorm.DB,username string) bool {
	var user model.User
	db.Where("username=?",username).First(&user)

	if user.ID!=0 {
		return true
	}

	return false

}
// 用户登录
func Login(c *gin.Context){
	//获取参数
	DB:=common.GetDB()
	username:=c.PostForm("username")
	password:=c.PostForm("password")

	if len(username)==0 {
		reps.UnprocessableEntity(c,nil,"用户名不能为空")
		return
	}
	//判断用户是否存在
	var user model.User
	DB.Where("username=?",username).First(&user)
	if user.ID==0 {
		reps.UnprocessableEntity(c,nil,"用户不存在")
		return
	}

	//判断密码是否正确
	err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))
	if err!=nil{
		reps.BadRequest(c,nil,"密码错误")
		return
	}


	//发放token
	token ,err :=common.ReleaseToken(user)
	if err!=nil{
		reps.ServerError(c,nil,"系统异常")
		return
	}


	//返回结果
	reps.Ok(c,gin.H{"token":token},"登录成功")
}

//重置密码
func ResetCode(c *gin.Context)  {
	username:=c.PostForm("username")
	newPassword:=c.PostForm("newpassword")
	newPasswordsure:=c.PostForm("newpasswordsure")
	sequrityCode:=c.PostForm("sequrityCode")

	if len(username)==0 {
		reps.UnprocessableEntity(c,nil,"用户名不能为空")
		return
	}

	if len(newPassword)==0 {
		reps.UnprocessableEntity(c,nil,"新密码不能为空")
		return
	}
	if len(newPassword)<6 {
		reps.UnprocessableEntity(c,nil,"新密码不能少于6位")
		return
	}
	if newPassword!=newPasswordsure {
		reps.UnprocessableEntity(c,nil,"两次密码输入不一致")
		return
	}
	if len(sequrityCode)==0 {
		reps.UnprocessableEntity(c,nil,"安全码不能为空")
		return
	}

	if len(sequrityCode)>=6 {
		reps.UnprocessableEntity(c,nil,"安全码不能超过6位")
		return
	}

	var user model.User
	DB:=common.GetDB()
	DB.Where("username=?",username).First(&user)
	if user.ID==0 {
		reps.UnprocessableEntity(c,nil,"用户名不存在")
		return
	}
	if user.SequrityCode!=sequrityCode {
		reps.UnprocessableEntity(c,nil,"安全码不匹配")
		return
	}
	//将得到的新密码加密
	hasednewPassword,err:=bcrypt.GenerateFromPassword([]byte(newPassword),bcrypt.DefaultCost)
	if err!=nil{
		reps.ServerError(c,nil,"加密错误")
		return
	}
	DB.Model(&user).Where("username=?",username).Update("password",hasednewPassword)
	reps.Ok(c,nil,"修改密码成功")

}
//获取当前用户信息
func CurrentUser(c *gin.Context)  {
	user,_:=c.Get("user")
	reps.Ok(c,gin.H{"user":user},"获取用户信息成功")
}
