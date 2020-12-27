package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

/*
Go程序唯一入口
*/
func main() {

	// 1. 安装  gin : go get -u github.com/gin-gonic/gin
	// 2. go.mod 中有对应的依赖信息

	// 获取一个路由, 使用默认中间件（logger和recovery）
	r := gin.Default()

	// get 方法 和 处理get 方法的逻辑
	r.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{ //返回一个JSON，状态码是200，gin.H是map[string]interface{}的简写
			"name": "tiaozao",
		})
	})

	// 获取参数

	// 获取参数方式1 Query -> querystring
	r.GET("/getParam", func(context *gin.Context) {
		//name := context.Query("name")
		name := context.DefaultQuery("name", "hhhhhh")
		age := context.Query("age")
		fmt.Println("get=====>", name, age)
		context.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	// 获取参数方式2 获取form参数
	r.POST("/getPost", func(context *gin.Context) {
		// DefaultPostForm取不到值时会返回指定的默认值
		name := context.PostForm("name")
		age := context.DefaultPostForm("age", "9")
		fmt.Println("post=====>", name, age)
		context.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	// 获取参数方式3 获取path参数
	r.GET("/u/s/:username/:address", func(c *gin.Context) {
		// 通过 Param 绑定获取path参数
		username := c.Param("username")
		address := c.Param("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	// 获取参数方式4 参数绑定 重点！！！
	// 绑定JSON的示例 ({"user": "wuzhuo", "password": "123456"})
	r.POST("/json", func(context *gin.Context) {
		var login Login
		if err := context.ShouldBind(&login); err == nil {
			fmt.Printf("login info:%#v\n", login)
			context.JSON(http.StatusOK, gin.H{
				"Name": login.Name,
				"Age":  login.Age,
			})
		} else {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	})

	r.GET("/get/field", func(context *gin.Context) {
		var login Login
		if err := context.ShouldBind(&login); err == nil {
			fmt.Printf("login info:%#v\n", login)
			context.JSON(http.StatusOK, gin.H{
				"Name": login.Name,
				"Age":  login.Age,
			})
		} else {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	})

	// 启动，并设置端口
	_ = r.Run(":8002")

	// 操作数据库
	//GetDB()
}

type Login struct {
	// 给字段绑定时设置一个别名，传入的参数名称必须和别名对应，否则无法绑定；若未设置别名，可以绑定字段名称（大小写都可以）
	// json:"hhhh" : json 类型的参数
	// form:"name" : form 表单类的参数
	// binding:"required" : 参数是否必须
	Name string `json:"hhhh" form:"name" binding:"required"`
	Age  int8
}

func GetDB() {
	// 操作数据库
	// 1. 安装 第三方库 go get github.com/go-sql-driver/mysql

	db, e := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/douyin")
	//err := db.Ping()
	//if err != nil {
	//	log.Fatal("ping failed: ", err)
	//}
	defer db.Close()
	if e != nil {
		panic("连接数据库出错")
		return
	}

	var id int
	var name string
	var age int
	var salary int
	var teamId int

	rows, err := db.Query("select id, name, age, salary, team_id from employees where id = ?", 1)
	if err != nil {
		log.Fatal("query failed: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name, &age, &salary, &teamId)
		if err != nil {
			log.Fatal("scan failed: ", err)
		}
		log.Printf("id: %d name:%s age:%d salary:%d teamId:%d\n", id, name, age, salary, teamId)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}

func init() {
	//sql.Register("mysql",&MySQLDriver{})
}
