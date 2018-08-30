package main

import (
	"context"
	"fmt"
	"gin-blog/pkg/setting"
	"gin-blog/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
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

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancal := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancal()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown", err)
	}

	log.Println("Server exiting")
}
