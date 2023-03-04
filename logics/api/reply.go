package api

import (
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"kilikili/logics/service"
	"kilikili/util/middleware"
	"net/http"
)

// PostReply 上传回复
func PostReply(c *gin.Context) {
	var reply service.ReplyService
	chaim, _ := middleware.ParseToken(c.GetHeader("Authorization"))
	err := c.ShouldBind(&reply)
	if err == nil {
		res := reply.Create(chaim.Id, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
		logging.Error(err)
	}
}
