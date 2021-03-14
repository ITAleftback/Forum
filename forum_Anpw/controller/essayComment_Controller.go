package controller


import (
	"forum_Anpw/common"
	"forum_Anpw/model"
	"forum_Anpw/reps"
	"github.com/gin-gonic/gin"
	"log"
)

func EssayComment(c *gin.Context)  {
	username:=c.PostForm("username")
	comment:=c.PostForm("comment")
	essayID:=c.PostForm("essayID")
	db:=common.GetDB()
	if len(comment)==0 {
		reps.UnprocessableEntity(c,nil,"评论不能为空")
		return
	}

	newComment:=model.Essay_Comment{
		Username:username,
		Comment:comment,
		EssayID:essayID,
	}

	db.Create(&newComment)
	reps.Ok(c,nil,"评论成功！")
}
func CurrentComment(c *gin.Context){
	essayID:=c.PostForm("essayID")
	db:=common.GetDB()
	var comment []model.Essay_Comment
	commentlist:=db.Where("essay_id=?",essayID).Order("created_at desc").Find(&comment)
	reps.Ok(c,gin.H{"commentlist":commentlist},"")
}

func DelComment(c *gin.Context){
	ID:=c.PostForm("ID")
	db:=common.GetDB()
	var comment model.Essay_Comment
	err:=db.Where("id=?",ID).Delete(&comment).Error
	if err!=nil {
		log.Println("Delete error:",err)
	}
	reps.Ok(c,nil,"删除成功")
}
