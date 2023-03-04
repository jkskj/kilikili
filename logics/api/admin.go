package api

import (
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"kilikili/logics/service"
	"net/http"
)

// BlockUser 拉黑用户
func BlockUser(c *gin.Context) {
	var admin service.AdminService
	err := c.ShouldBind(&admin)
	if err == nil {
		res := admin.Block()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
		logging.Error(err)
	}
}

// ExamineVideo 获取审核视频
func ExamineVideo(c *gin.Context) {
	var admin service.AdminService
	err := c.ShouldBind(&admin)
	if err == nil {
		res := admin.Examine()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
		logging.Error(err)
	}
}

// IsPassVideo 返回审核结果
func IsPassVideo(c *gin.Context) {
	var admin service.AdminService
	err := c.ShouldBind(&admin)
	if err == nil {
		res := admin.IsPass()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
		logging.Error(err)
	}
}

// DeleteComment 删除评论
func DeleteComment(c *gin.Context) {
	var admin service.AdminService
	err := c.ShouldBind(&admin)
	if err == nil {
		res := admin.Delete()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
		logging.Error(err)
	}
}
