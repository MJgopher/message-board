package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespInternalError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"info": "服务器出错",
	})
}

func RespErrorWithData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"info": "失败",
		"data": data,
	})
}

func RespSuccessful(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"info": "成功",
	})
}

func RespSuccessfulWithData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"info": "成功",
		"data": data,
	})
}
func RespAllComments(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"info":     "成功",
		"Comments": data,
	})
}
func RespAllPosts(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"info":  "成功",
		"Posts": data,
	})
}
