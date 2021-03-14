package reps


import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ok(c *gin.Context,data gin.H, msg string)  {
	c.JSON(http.StatusOK,gin.H{"code":200,"data":data,"msg":msg})
}
func ServerError(c *gin.Context,data gin.H, msg string)  {
	c.JSON(http.StatusInternalServerError,gin.H{"code":500,"data":data,"msg":msg})
}
func BadRequest(c *gin.Context,data gin.H, msg string)  {
	c.JSON(http.StatusBadRequest,gin.H{"code":400,"data":data,"msg":msg})
}
func UnprocessableEntity(c *gin.Context,data gin.H, msg string)  {
	c.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"data":data,"msg":msg})
}

func Unauthorized(c *gin.Context,data gin.H, msg string)  {
	c.JSON(http.StatusUnauthorized,gin.H{"code":401,"data":data,"msg":msg})
}
