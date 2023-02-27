package v1

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"kilikili/logics/api"
	"kilikili/util/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("my_session", store))
	v1 := r.Group("/api/v1")
	{
		//用户注册
		v1.POST("/register", api.UserRegister)
		//用户登录
		v1.POST("/login", api.UserLogin)

		authed := v1.Group("/")

		authed.Use(middleware.JWT())
		{

			v1.POST("/video", api.PostVideo)

			v1.GET("/video/:id", api.WatchVideo)

			v1.PUT("/change/personal", api.UserChangePersonal)

			v1.PUT("/change/password", api.UserChangePassword)

			v1.POST("/comment/bullet/:id", api.PostBullet)

			v1.GET("/comments/bullet/:id", api.ListBullet)

			v1.POST("/comment/normal/:id", api.PostComment)

			v1.GET("/comments/normal/:id", api.ListComment)

			v1.POST("/comment/reply/:id", api.PostReply)

			v1.POST("/video/interaction", api.Interact)

			v1.PUT("/video/interaction", api.CancelInteraction)

			admin := v1.Group("/admin")

			admin.Use(middleware.Check())
			{
				admin.PUT("/block", api.BlockUser)

				admin.GET("/video", api.ExamineVideo)

				admin.PUT("/video", api.IsPassVideo)

				admin.DELETE("/comment", api.DeleteComment)
			}

		}
	}
	return r
}
