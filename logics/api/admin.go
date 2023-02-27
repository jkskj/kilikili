package api

import (
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"kilikili/logics/service"
	"net/http"
)

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
