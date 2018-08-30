package main

import (
	"net/http"
	"fmt"
	"gin-blog/pkg/setting"
	"gin-blog/routers"
)

func main() {
	router := routers.InitRouter()
	//router.GET("/test", func(context *gin.Context) {
	//	context.JSON(200, gin.H{
	//		"message": "api of blog !",
	//	})
	//})

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
