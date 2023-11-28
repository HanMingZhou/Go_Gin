package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID   string `json:"id" form:"id" xml:"id"`
	Name string `json:"name" form:"name" xml:"name"`
}

type UserJson struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

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

	r.GET("/users/:id/:name", func(c *gin.Context) {
		id := c.Param("id")
		name := c.Param("name")
		// 根据ID查询用户
		user := getUserById(id)
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"id":   user,
		})
	})

	r.GET("/user", func(ctx *gin.Context) {
		/*
			http://127.0.0.1:8080/user?name=iphone
		*/
		name := ctx.Query("name")
		age := ctx.DefaultQuery("age", "18")
		ctx.String(200, "%s,%s", name, age)
	})

	r.POST("/post", func(ctx *gin.Context) {
		ctx.Request.ParseForm()
		name := ctx.PostForm("name")
		age := ctx.PostForm("age")
		ctx.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})

	})

	r.POST("/postUser", func(ctx *gin.Context) {
		var u user
		ctx.Bind(&u)
		ctx.String(http.StatusOK, "name=%s,id=%s 嘻嘻嘻", u.Name, u.ID)

	})

	// shouldBindJson
	r.POST("/postjson", func(ctx *gin.Context) {
		var u UserJson
		if err := ctx.ShouldBindJSON(&u); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}
		fmt.Printf("%#v\n", u)
		// 返回响应体给client
		ctx.JSON(http.StatusOK, gin.H{
			"email": u.Email,
			"name":  u.Name,
		})
	})

	// shouldbind
	/*
		说明shouBind()方法可以根据请求中contentType的不同类型，采用不同的方式进行处理。
	*/
	r.POST("/postShouldBind", func(ctx *gin.Context) {
		var U user
		if err := ctx.ShouldBind(&U); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			fmt.Printf("%v\n", U)
			ctx.JSON(http.StatusOK, gin.H{"id": U.ID, "Name": U.Name})
		}
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
