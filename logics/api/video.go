package api

import (
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"kilikili/logics/service"
	"kilikili/util/middleware"
	"net/http"
)

func PostVideo(c *gin.Context) {
	var createVideo service.VideoService
	chaim, _ := middleware.ParseToken(c.GetHeader("Authorization"))
	err := c.ShouldBind(&createVideo)
	if err == nil {
		res := createVideo.Create(chaim.Id)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
		logging.Error(err)
	}

}
func WatchVideo(c *gin.Context) {
	var watchVideo service.VideoService
	err := c.ShouldBind(&watchVideo)
	if err == nil {
		res := watchVideo.Watch(c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
		logging.Error(err)
	}

}
