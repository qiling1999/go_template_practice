package main

import "github.com/gin-gonic/gin"

func sayHello(c *gin.Context)  {//这个函数里面返回的是一个json格式的数据  状态码200
	c.JSON(200, gin.H{
		"message": "hello",   //返回给前端浏览器上的值
	})
}

func main() {
	r := gin.Default()//返回默认的路由引擎

	r.GET("/hello", sayHello)//用户访问127.0.0.1:8080/hello时会去执行sayHello函数
	//启动服务,默认在0.0.0.0：8080启动服务，或者设置为其他的端口
	r.Run(":9090")
}
