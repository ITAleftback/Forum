package common


import (
	"github.com/dgrijalva/jwt-go"
	"forum_Anpw/model"
	"time"
)


var jwtKey =[]byte("Anpw")

type Claims struct {
	UserID uint
	jwt.StandardClaims
}

//颁发证书
func ReleaseToken(user model.User)(string,error){
	//token过期时间
	expirationTime :=time.Now().Add(7*24*time.Hour)
	claims:=&Claims{
		UserID:         user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt:expirationTime.Unix(),//到期时间
			IssuedAt:time.Now().Unix(),//签发时间
			Subject:"user token",//主题
			Issuer:"selfblog",//发行人

		},
	}
	//加密封装
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	//获取签名令牌
	tokenString,err:=token.SignedString(jwtKey)

	if err!=nil {
		return "",err
	}
	return tokenString,nil
}
//解析token
func PareseToken(tokenString string)(*jwt.Token,*Claims,error)  {
	claims:=&Claims{}

	token,err:=jwt.ParseWithClaims(tokenString,claims, func(token *jwt.Token) (i interface{}, e error) {
		return jwtKey,nil
	})

	return token,claims,err
}