package api

import (
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"kilikili/logics/service"
	"kilikili/util/middleware"
	"net/http"
)

// 上传评论
func PostComment(c *gin.Context) {
	var comment service.CommentService
	chaim, _ := middleware.ParseToken(c.GetHeader("Authorization"))
	err := c.ShouldBind(&comment)
	if err == nil {
		res := comment.Create(chaim.Id, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
		logging.Error(err)
	}
}

// 获取评论
func ListComment(c *gin.Context) {
	var comment service.CommentService
	err := c.ShouldBind(&comment)
	if err == nil {
		res := comment.List(c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
		logging.Error(err)
	}
}
