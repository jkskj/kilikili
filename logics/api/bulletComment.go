package api

import (
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"kilikili/logics/service"
	"kilikili/util/middleware"
	"net/http"
)

// PostBullet 上传弹幕
func PostBullet(c *gin.Context) {
	var comment service.BulletCommentService
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

// ListBullet 获取弹幕
func ListBullet(c *gin.Context) {
	var comment service.BulletCommentService
	err := c.ShouldBind(&comment)
	if err == nil {
		res := comment.List(c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
		logging.Error(err)
	}
}
