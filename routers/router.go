package routers

import (
	"blog/controller"
	"blog/pkg/setting"
	"github.com/gin-gonic/gin"
	jwt "blog/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/auth", controller.GetAuth)

	apiv1 := r.Group("api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取文章列表
		apiv1.GET("/articles", controller.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", controller.GetArticle)
		//新建文章
		apiv1.POST("/articles", controller.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", controller.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", controller.DeleteArticle)

		apiv1.GET("/tags", controller.GetTag)
		apiv1.POST("/tags", controller.AddTag)
		apiv1.PUT("/tags/:id", controller.EditTag)
		apiv1.DELETE("/tags/:id", controller.DeleteTag)
	}

	return r
}
