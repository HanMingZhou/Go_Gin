package main

import (
	"fmt"
	"net/http"
	"path"

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

	/*
		query
	*/
	r.GET("/user", func(ctx *gin.Context) {
		/*
			http://127.0.0.1:8080/user?name=iphone
		*/
		name := ctx.Query("name")
		age := ctx.DefaultQuery("age", "18")
		ctx.String(200, "%s,%s", name, age)
	})

	r.GET("/userGetQuery", func(ctx *gin.Context) {
		//第三种方式，考虑的比较全面
		name, ok := ctx.GetQuery("query")
		if !ok {
			name = "someone"
		}
		ctx.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})

	/*
		postform
	*/
	r.LoadHTMLFiles("./login.html", "./index.html")
	r.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/login", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		fmt.Println("username,password", username, password)
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"Password": password,
		})
	})

	r.GET("/GetPostForm", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/GetPostForm", func(ctx *gin.Context) {
		username, ok := ctx.GetPostForm("username")
		if ok {
			username = "sb"
		}
		password, ok := ctx.GetPostForm("password")
		if ok {
			password = "***"
		}
		fmt.Println("username,password", username, password)
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"Password": password,
		})
	})

	/*
	   Param
	*/
	r.GET("/:name/:age", func(ctx *gin.Context) {
		name := ctx.Param("name")
		age := ctx.Param("age")
		// respond
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

	/* ShouldBind()
	它用于将请求携带的参数和后端的结构体绑定起来
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

	/*
		upload 文件
	*/
	r.LoadHTMLFiles("./upload.html")
	r.GET("/upload", GetUpload)
	r.POST("/upload", PostUpload)

	/* redirect*/
	r.GET("/redirect", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
		fmt.Println("已经重新定向bing")
	})
	r.GET("/a", func(ctx *gin.Context) {
		ctx.Request.URL.Path = "/b"
		ctx.Request.Method = "POST"
		r.HandleContext(ctx)
	})
	r.POST("/b", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "b",
		})

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

func GetUpload(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "upload.html", nil)

}
func PostUpload(ctx *gin.Context) {
	if f, err := ctx.FormFile("test"); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		//将文件保存在服务器
		//dst := fmt.Sprintf("./%s", f.Filename)//写法1
		detination := path.Join("C:/Users/hanmingzhou/Desktop/", f.Filename)
		fmt.Println(detination)
		err = ctx.SaveUploadedFile(f, detination)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
		} else {
			// respond
			ctx.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}

	}

}
