package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"log"
	"net/http"
	"time"
)

func main() {

	// 创建路由, 默认使用了两个中间件， Logger和 Recovery
	r := gin.Default()
	// 也可以创建不带中间件的路由
	//r := gin.New()

	// 绑定路由规则
	r.GET("/test", func(context *gin.Context) {
		fullPath := context.FullPath()
		fmt.Println("全路径:", fullPath)
		//context.Writer.WriteString("success")
		context.String(http.StatusOK, "hello world")
	})

	// 从url中获取参数(restful风格的参数)
	r.GET("/user/:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		// action的值为， name那一层往后的所有
		action := context.Param("action")
		context.String(http.StatusOK, name + "  " + action)
	})

	// 从url中获取参数(问号后面的参数)
	r.GET("/testparam", func(context *gin.Context) {
		username := context.DefaultQuery("username", "jack")
		password := context.DefaultQuery("password", "123")
		fmt.Println(">>>>", username, password)
		context.String(http.StatusOK, username)
	})

	// 表单参数, 这里应该post请求， 这里方便测试都用get
	r.GET("/formparam", func(c *gin.Context) {
		// 表单参数设置默认值
		tom := c.DefaultPostForm("username", "tom")
		fmt.Println(tom)
		// 接收其他
		gender := c.PostForm("gender")

		hobbys := c.PostFormArray("hobby")
		fmt.Println(">>>", hobbys, gender)
		c.String(http.StatusOK, tom)
	})

	// 上传文件
	r.GET("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		c.SaveUploadedFile(file, file.Filename)   // 默认项目根目录
		c.String(http.StatusOK, "upload " + file.Filename)
	})

	// 上传多个文件
	// 限制表单上传大小 8M, 默认为32M
	r.MaxMultipartMemory = 8 << 20
	r.GET("/upload2", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			fmt.Println(">>>> err :", err)
		}
		// 获取所有文件
		files := form.File["files"]
		for _, file := range files{
			// 逐个保存
			err := c.SaveUploadedFile(file, file.Filename)
			if err != nil {
				fmt.Println(">>>>err :", err)
				return
			}
		}
		c.String(http.StatusOK, "upload success")
	})


	r.POST("/dfadf")
	r.PUT("/hsdf")



	// 上面都是单个测试使用， 下面是路由组使用, gin中采用的路由库是基于httprouter做的
	// 路由组1， 处理get请求
	v1 := r.Group("/v1")
	{	// 这里大括号就是分一下组
		v1.GET("/login", login)
		v1.GET("/submit", submit)
	}

	v2 := r.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}

	// 各种绑定参数的形式，  可以去了解一下， 手动获取也是可以的

	// 下面是gin的渲染,  就是返回数据格式
	// 1.json
	r.GET("/somejson", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message":"somejson", "status":200})
	})
	// 2. 结构体响应
	r.GET("/somestruct", func(c *gin.Context) {
		var msg struct{
			Name string
			Number int
		}
		msg.Number = 1
		msg.Name = "struct"
		c.JSON(http.StatusOK, msg)
	})
	// 3.xml
	r.GET("/somexml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message":"abc"})
	})
	// 4. yaml
	r.GET("/someyaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"name":"zhangsan"})
	})
	// 5. protobuf
	r.GET("/someprotobuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		// 定义数据
		lable := "lable"
		// 传protobuf格式数据
		data := &protoexample.Test{
			Label: &lable,
			Reps: reps,
		}
		c.ProtoBuf(http.StatusOK, data)
	})


	// html模板渲染, 本质上是字符串替换
	// 加载模板
	r.LoadHTMLGlob("template/*")
	//r.LoadHTMLFiles("template/index.tmpl")
	r.GET("/index", func(c *gin.Context) {
		// 根据文件名渲染
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"title":"标题"})
	})

	// 重定向
	r.GET("/redirect", func(c *gin.Context) {
		// 支持内部和外部重定向
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	// 异步
	r.GET("/long_async", func(c *gin.Context) {
		// 需要一个副本
		copyContext := c.Copy()
		// 异步处理
		go func() {
			time.Sleep(3 * time.Second)
			log.Println(">>>异步", copyContext.Request.URL.Path)
		}()
	})
	// 同步
	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println(">>>同步", c.Request.URL.Path)
	})


	// 中间件
	// 作用于全局
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 作用于单个路由
	r.GET("/middleware", gin.Logger(), gin.Recovery())

	// 作用于某个组
	authorized := r.Group("/")
	authorized.Use()
	{
		//authorized.POST("/login", loginEndpoint)
		//authorized.POST("/submit", submitEndpoint)
	}


	// 监听端口, 默认8080
	r.Run(":8000")
}

// 自定义中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// 给context实例设置一个值
		c.Set("geektutu", "111")
		// 请求前
		c.Next()
		//请求后
		latency := time.Since(t)	// 这个返回间隔的时间
		log.Println(latency)
	}
}


// 路由哪里使用的组件
func login(c *gin.Context)  {
	c.DefaultQuery("name", "jack")
	c.String(http.StatusOK, "login success")
}

func submit(c *gin.Context)  {
	c.DefaultQuery("name", "jack")
	c.String(http.StatusOK, "submit success")
}