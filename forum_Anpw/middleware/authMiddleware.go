package middleware


import (
	"forum_Anpw/common"
	"forum_Anpw/model"
	"forum_Anpw/reps"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthMiddleware()gin.HandlerFunc  {
	//获取authorization header
	return func(c *gin.Context) {
		tokenString :=c.GetHeader("Authorization")

		//验证token格式

		if tokenString ==""||!strings.HasPrefix(tokenString,"Bearer ") {
			reps.Unauthorized(c,nil,"权限不足")
			//中止
			c.Abort()
			return
		}
		tokenString=tokenString[7:]

		token,claims,err:=common.PareseToken(tokenString)
		//如果token无效
		if err!=nil || !token.Valid{
			reps.Unauthorized(c,nil,"权限不足")
			c.Abort()
			return
		}
		userID:=claims.UserID
		DB:=common.GetDB()
		var user model.User
		DB.First(&user,userID)

		//用户不存在
		if user.ID==0 {
			reps.Unauthorized(c,nil,"权限不足")
			c.Abort()
			return
		}
		//用户存在
		c.Set("user",user)
		c.Next()
	}
}

