package main
//gin框架是舞台，你自己的是才艺，gin是一套web框架，无重复造轮子
//关于web，web是基于HTTP协议进行交互的应用网络，通过使用浏览器/APP访问的各种资源
//后端的内容，gin框架是基于httprouter这个库开发的web框架，速度很快
//RESTful API 与技术无关，代表的是一种软件架构风格，"表征状态转移"
//简单来说 REST的含义就是客户端与web服务器之间进行交互的时候使用HTTP协议中的4个请求方法代表不同的动作
//GET用来获取资源
//POST用来新建资源
//PUT用来更新资源
//DELETE用来删除资源
//只要API程序遵循了REST风格，那就可以称其为RESTful API，目前在前后端分离的架构中，前后端基本都是通过RESTful API来进行交互的

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//创建一个默认的路由引擎
	r := gin.Default()
	//GET:请求方式； /hello:请求路径
	//当客户端以GET方法请求/hello路径时，会执行后面的匿名函数  即127.0.0.1:8080/hello
	r.GET("/hello", func(c *gin.Context){
		//c.JSON:返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "hello world!",
		})
	})
	//当客户端以GET方法请求/book路径时，会执行后面的匿名函数  即127.0.0.1:8080/hello
	r.POST("/book", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "POST",
		})
	})
	r.PUT("/book", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT",
		})
	})
	r.DELETE("/book", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "DELETE",
		})
	})
	/*//RESTful API风格例如
	r.GET("/book",...)
	r.POST("/create_book",...)
	r.PUT("/update_book",...)
	r.DELETE("/delete_book",...)
	*/
	//启动HTTP服务，默认在0.0.0.0：8080启动服务
	r.Run()
}