package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-wechat-bot/bot"
)

func httpServer()  {
	r := gin.Default()
	r.GET("/face-get", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"code":200,
			"data": bot.SelectRecordForKeysInfo(""),
		})
	})

	err := r.Run(":19004")
	if err != nil {
		fmt.Println("http 服务异常")
		return
	}

}
func main() {
	err := bot.InitDb()
	if err != nil {
		return
	}
	bot.StartBot()

}
