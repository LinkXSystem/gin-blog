package routers

import (
	"gin-blog/middleware/jwt"
	"gin-blog/pkg/setting"
	"gin-blog/routers/api"
	"gin-blog/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	router.GET("/auth", api.GetAuth)

	api := router.Group("/api/v1")
	api.Use(jwt.JWT())
	{
		// 标签
		api.GET("tags", v1.GetTags)
		api.POST("/tags", v1.AddTags)
		api.PUT("/tags/:id", v1.EditTag)
		api.DELETE("/tags/:id", v1.DeleteTag)
		// 文章
		api.GET("/articles", v1.GetArticle)
		api.GET("/articles/:id", v1.GetArticle)
		api.POST("/articles", v1.AddArticle)
		api.PUT("/articles/:id", v1.EditArticle)
		api.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return router
}
