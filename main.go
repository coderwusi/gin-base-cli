package main

import (
	"fmt"
	"gin-base-cli/middleware"
	"gin-base-cli/model"
	"gin-base-cli/setting"
	"github.com/gin-gonic/gin"
)

func main() {
	setting.ViperInit()                                        //读取环境配置
	setting.LogInit(setting.Conf.LogConfig, setting.Conf.Mode) //日志加载成功
	model.MysqlInit(setting.Conf.MySQLConfig)                  //连接mysql
	//setting.RedisClient()                     //连接redis
	r := gin.Default()       //默认使用Logger(), Recovery()
	r.Use(middleware.Cors()) //启用跨域
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	if err := r.Run(); err != nil {
		fmt.Println("启动服务失败：", err)
	} // 监听并在 0.0.0.0:8080 上启动服务

}
