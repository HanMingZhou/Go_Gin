package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() //1.创建路由
	/*
		r.GET("/ping", func(ctx *gin.Context) { //2.绑定路由规则，执行的函数
			ctx.JSON(200,
				gin.H{
					"message": "pong",
				})
		})
	*/

	/*
		r.GET("/hello/*name", func(ctx *gin.Context) {
			name := ctx.Param("name")

			ctx.String(http.StatusOK, "hello %s ", name)

				ctx.JSON(200, gin.H{
					"user id is ": id,
				})
		})
	*/
	// api := r.Group("/api")
	// {
	// 	api.GET("/users", func(ctx *gin.Context) {
	// 		ctx.String(200, "hello, i am get")
	// 	})

	// 	api.POST("/users", func(ctx *gin.Context) {
	// 		ctx.String(200, "hello, i am post")
	// 	})

	// 	// api.DELETE("/users", func(ctx *gin.Context) {
	// 	// 	ctx.String(404, "not found")
	// 	// })

	// 	api.GET("/users", func(ctx *gin.Context) {
	// 		ctx.Redirect(http.StatusMovedPermanently, "/redirect_users")
	// 	})
	// 	api.GET("/redirect_users", func(ctx *gin.Context) {
	// 		ctx.String(200, "hello, i am redirect_users")
	// 	})

	// }

	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		// 根据ID查询用户
		user := getUserById(id)
		c.JSON(http.StatusOK, user)
	})

	r.Run() //3.监听端口，默认8080
}

func getUserById(id string) (str string) {
	switch id {
	case "HUAWEI":
		return "HUAWEI"
	case "XIAOMI":
		return "XIAOMI"
	default:
		return "NO PHONE BRAND"
	}
}
