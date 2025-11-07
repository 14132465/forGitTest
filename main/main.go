package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Auth 统一验证拦截器，每个接口都需要提前验证
func Auth(w http.ResponseWriter, r *http.Request) error {
	//这里添加的是你的验证层面的信息类似于一个中间键

	queryParams := r.URL.Query() // 返回一个值的集合 map[string][]string
	id := queryParams.Get("id")  // 获取指定参数的第一个值

	//id := r.Header.Get("id")
	if id == "" {
		return errors.New("id不能为空")
	}
	_, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("id参数是数字")
	}
	fmt.Println("我是验证通过")
	return nil
}

// Check 执行正常业务逻辑
func Check(w http.ResponseWriter, r *http.Request) {
	//执行正常业务逻辑
	fmt.Println("执行check！  正常业务逻辑")
}

func InitFlag() {
	//定义命令行参数方式1
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")

	//解析命令行参数
	flag.Parse()
	fmt.Println(name, age, married, delay)
	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
}

type Person struct {
	Age       int    `json:"age,string"`
	Name      string `json:"nickName"`
	Niubility bool   `json:"niubility"`
}

func main() {

	// initFlag()

	// //1、过滤器
	// filter := common.NewFilter()
	// //注册拦截器
	// filter.RegisterFilterUri("/check", Auth)
	// //2、启动服务
	// http.HandleFunc("/check", filter.Handle(Check))

	// //fmt.Println("server runing , the port : 8083")
	// //启动服务
	// http.ListenAndServe(":8083", nil)

	r := gin.Default()

	srv := &http.Server{
		Addr:              "0.0.0.0:8888",
		Handler:           r,
		ReadHeaderTimeout: 2 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	route := r.Group("/api")
	route.GET("/", HandlerForRoot)

	srv.ListenAndServe()

}

func HandlerForRoot(c *gin.Context) {

	c.JSON(200, `{ "data":"hello !" }`)
}
