package common


import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"forum_Anpw/model"
)

//连接数据库


var DB *gorm.DB
func InitDB() *gorm.DB {
	driverName:="mysql"
	host :="localhost"
	port :="3306"
	database :="forum"
	username :="root"
	password :=""
	charset :="utf8"
	parseTime:="true"
	loc:="Local"
	args :=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		parseTime,
		loc,
	)

	db,err:=gorm.Open(driverName,args)
	if err!=nil {
		panic("failed to connect database,err:"+err.Error())

	}
	DB=db
	db.AutoMigrate(&model.User{},&model.Essay{},&model.Essay_Comment{})
	return DB
}

func GetDB() *gorm.DB {
	return DB
}