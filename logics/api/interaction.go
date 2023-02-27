package api

import (
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"kilikili/logics/service"
	"kilikili/util/middleware"
	"net/http"
)

func Interact(c *gin.Context) {
	var interaction service.InteractionService
	chaim, _ := middleware.ParseToken(c.GetHeader("Authorization"))
	err := c.ShouldBind(&interaction)
	if err == nil {
		res := interaction.Interact(chaim.Id)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
		logging.Error(err)
	}
}
func CancelInteraction(c *gin.Context) {
	var interaction service.InteractionService
	chaim, _ := middleware.ParseToken(c.GetHeader("Authorization"))
	err := c.ShouldBind(&interaction)
	if err == nil {
		res := interaction.Cancel(chaim.Id)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
		logging.Error(err)
	}
}
