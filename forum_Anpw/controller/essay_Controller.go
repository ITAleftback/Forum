package controller


import (
	"github.com/gin-gonic/gin"
	"log"
	"forum_Anpw/common"
	"forum_Anpw/model"
	"forum_Anpw/reps"
)

func WriteEssay(c *gin.Context)  {
	username:=c.PostForm("username")
	title:=c.PostForm("title")
	essay:=c.PostForm("essay")
	db:=common.GetDB()
	if len(title)==0 {
		reps.UnprocessableEntity(c,nil,"文章标题不能为空")
		return
	}
	if len(essay)==0 {
		reps.UnprocessableEntity(c,nil,"文章不能为空")
		return
	}


	newEssay:=model.Essay{
		Title:  title,
		Essay:  essay,
		Author: username,
	}

	db.Create(&newEssay)

	reps.Ok(c,nil,"发布成功！")
}
//查看文章
func CurrentEssay(c *gin.Context){

	db:=common.GetDB()
	var essay []model.Essay
	essaylist:=db.Order("created_at desc").Find(&essay)
	reps.Ok(c,gin.H{"essaylist":essaylist},"")
}

func YourEssay(c *gin.Context)  {
	username:=c.PostForm("username")
	db:=common.GetDB()
	var essay []model.Essay
	essaylist:=db.Where("author=?",username).Order("created_at desc").Find(&essay)
	reps.Ok(c,gin.H{"essaylist":essaylist},"")

}

func DelEssay(c *gin.Context){
	ID:=c.PostForm("ID")
	db:=common.GetDB()
	var essay model.Essay
	err:=db.Where("id=?",ID).Delete(&essay).Error
	if err!=nil {
		log.Println("Delete error:",err)
	}
	reps.Ok(c,nil,"删除成功")
}

func EditEssay(c *gin.Context){
	ID:=c.PostForm("ID")
	newTitle:=c.PostForm("newtitle")
	newEssay:=c.PostForm("newessay")
	if len(newTitle)==0 {
		reps.UnprocessableEntity(c,nil,"修改的标题不能为空")
	}
	if len(newEssay)==0 {
		reps.UnprocessableEntity(c,nil,"修改的文章不能为空")
	}
	db:=common.DB
	var essay model.Essay
	db.Model(&essay).Where("id=?",ID).Update(map[string]interface{}{"title":newTitle,"essay":newEssay})
	reps.Ok(c,nil,"修改成功")
}


