package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"projecttimer/db"
	"projecttimer/frontend"
	"projecttimer/utils"
	"strconv"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func RegisterRouters() {
	Gin.Use(Cors())
	//Gin.Use(static.Serve("/", static.LocalFile("./flutter/", true))) //使用 contrib 避免与get path "/"a 冲突
	//Gin.Use(static.Serve("/", static.LocalFile("./flutter/", true))) //使用 contrib 避免与get path "/"a 冲突
	Gin.StaticFS("/timer/", http.FS(frontend.Static))
	//web.Register(Gin, "/dataflow/")
	//frontend.Register(Gin, "/dataflow/")
	//tg := Gin.Group("/")
	v1 := Gin.Group("/v1")
	v1.Use(ErrorHandler())
	handlerHoldWS(v1)
	handlerPing(v1)
	handlerGroup(v1)

}
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		// 检查是否有错误发生
		if len(c.Errors) > 0 {
			// 获取最后一个错误
			err := c.Errors.Last()
			// 将错误信息转换为标准的响应格式
			c.JSON(http.StatusOK, RespError(err.Error()))

			// 阻止其他中间件和处理函数继续执行
			c.Abort()
		}
	}
}

func handlerGroup(r *gin.RouterGroup) {
	r.GET("/customers", func(c *gin.Context) {
		offset := c.Query("offset")
		limit := c.Query("limit")
		var count int64
		offsetInt, err := strconv.Atoi(offset)
		if err != nil {
			fmt.Println("转换失败:", err)
			return
		}
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			fmt.Println("转换失败:", err)
			return
		}
		customer := db.Customer{}
		var cs []db.Customer
		err = customer.ListPage(&count, &cs, offsetInt, limitInt)
		if err != nil {
			utils.Log.Error(err)
			c.JSON(http.StatusOK, RespError(err.Error()))
		}
		c.JSON(http.StatusOK, RespSuccess(Page{Count: count, Data: cs}))
	})
	r.POST("/addCustomer", func(c *gin.Context) {
		customer := db.Customer{}
		c.BindJSON(&customer)
		customer.Add()
		c.JSON(http.StatusOK, RespSuccess(nil))
	})
	r.POST("/updateCustomer", func(c *gin.Context) {
		customer := db.Customer{}
		err := c.BindJSON(&customer)
		if err != nil {
			c.JSON(http.StatusOK, RespError(err.Error()))
		}
		err = customer.Update()
		if err != nil {
			c.JSON(http.StatusOK, RespError(err.Error()))
		}
		c.JSON(http.StatusOK, RespSuccess(nil))
	})
	r.GET("/deleteCustomer", func(c *gin.Context) {
		_id := c.Query("id")
		db.Customer{}.Delete(_id)
	})
}

func handlerPing(r *gin.RouterGroup) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, RespSuccess("Pong."))
	})
}
