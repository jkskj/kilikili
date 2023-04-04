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
			//上传视频
			v1.POST("/video", api.PostVideo)
			//观看视频
			v1.GET("/video/:id", api.WatchVideo)
			//修改个人信息
			v1.PUT("/change/personal", api.UserChangePersonal)
			//修改密码
			v1.PUT("/change/password", api.UserChangePassword)
			//发弹幕
			v1.POST("/comment/bullet/:id", api.PostBullet)
			//获取弹幕
			v1.GET("/comments/bullet/:id", api.ListBullet)
			//发评论
			v1.POST("/comment/normal/:id", api.PostComment)
			//获取评论
			v1.GET("/comments/normal/:id", api.ListComment)
			//发回复
			v1.POST("/comment/reply/:id", api.PostReply)
			//互动
			v1.POST("/video/interaction", api.Interact)
			//取消评论
			v1.PUT("/video/interaction", api.CancelInteraction)

			admin := authed.Group("/admin")

			admin.Use(middleware.Check())
			{
				//拉黑用户
				admin.PUT("/block", api.BlockUser)
				//获取未审核视频
				admin.GET("/video", api.ExamineVideo)
				//上传审核结果
				admin.PUT("/video", api.IsPassVideo)
				//删除评论
				admin.DELETE("/comment", api.DeleteComment)
			}

		}
	}
	return r
}
