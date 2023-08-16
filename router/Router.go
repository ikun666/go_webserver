package router

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ikun666/go_webserver/control"
	"github.com/ikun666/go_webserver/middleware"
	"github.com/spf13/viper"
)

// 注册路由函数
type IFnRegisterRoute = func(rgPublic, rgAuthor *gin.RouterGroup)

// 注册路由切片
var fnRoutes []IFnRegisterRoute

func RegisterRoutes(fn IFnRegisterRoute) {
	if fn != nil {
		fnRoutes = append(fnRoutes, fn)
	}
}
func InitRoutes() {
	RegisterRoutes(func(rgPublic *gin.RouterGroup, rgAuthor *gin.RouterGroup) {
		userControl := control.NewUserControl()
		rgPublicUser := rgPublic.Group("user")

		rgPublicUser.GET("/login", userControl.Login)

		rgAuthorUser := rgAuthor.Group("user")
		rgAuthorUser.POST("/add", userControl.AddUser)
		rgAuthorUser.GET("/id", userControl.GetUserByName)
		// rgAuthorUser.POST("/list", userControl.GetUserList)
		// rgAuthorUser.POST("/update", userControl.UpdateUser)
		rgAuthorUser.GET("/delete", userControl.DeleteUserByName)
	})
}
func InitRouter() {
	// 无优雅关闭
	//默认gin框架
	// r := gin.Default()
	// //分组
	// rgPublic := r.Group("public")
	// rgAuthor := r.Group("author")
	// InitRoutes()
	// for _, fnRoute := range fnRoutes {
	// 	fnRoute(rgPublic, rgAuthor)
	// }

	// port := viper.GetString("server.port")

	// err := r.Run(fmt.Sprintf(":%s", port))
	// if err != nil {
	// 	panic(fmt.Sprintf("开启服务错误：%s\n", err.Error()))
	// }
	// fmt.Printf("开启服务成功\n")

	//优雅关闭
	//直接在gin 文档抄过来
	//默认gin框架
	r := gin.Default()
	//分组
	rgPublic := r.Group("public")
	rgAuthor := r.Group("author")
	rgAuthor.Use(middleware.Author())
	InitRoutes()
	for _, fnRoute := range fnRoutes {
		fnRoute(rgPublic, rgAuthor)
	}

	port := viper.GetString("server.port")
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			// global.Logger.Error(fmt.Sprintf("开启服务错误：%s\n", err.Error()))
			log.Fatalf("开启服务错误：%s\n", err.Error())
			return
		}
		// global.Logger.Info(fmt.Sprintf("开启服务成功"))
		log.Println("开启服务成功")
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
