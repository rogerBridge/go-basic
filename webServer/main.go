package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/test", sayHelloName) //设置访问的路由
	http.HandleFunc("/foo", foo)
	err := http.ListenAndServe(":8000", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	//randStatement()
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	// 可以获取进入的请求的方法
	fmt.Println(r.Method)
	// 首先, 对请求进行一系列处理
	err := r.ParseForm()
	if err != nil {
		fmt.Println("参数解析错误!")
	}
	fmt.Println("解析url参数的结果是: ", r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path:  ", r.URL.Path)
	fmt.Println("scheme:  ", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	//if _, ok := (r.Form)["name"]; ok {
	//	name := (r.Form)["name"][0]
	//	fmt.Fprintf(w, "Hello "+name+" "+randStatement()) //这个写入到w的是输出到客户端的
	//} else {
	//	fmt.Fprintf(w, "Hello User "+randStatement())
	//}
	_, _ = w.Write([]byte("安瑞峰"))
}

// 返回处理后的json字符串
func foo(w http.ResponseWriter, r *http.Request) {
	//err := r.ParseForm()
	//if err != nil {
	//	fmt.Println("参数解析错误!")
	//}
	//err = r.ParseMultipartForm(0)
	//fmt.Println("解析multiPartForm: ", r.MultipartForm.Value)
	//// 开始处理请求
	//fmt.Println("解析url参数的结果是: ", r.Form) //这些信息是输出到服务器端的打印信息
	err := r.ParseForm()
	fmt.Println("postform is:", r.Form)

	// 解析post请求的body部分
	v, err := ioutil.ReadAll(r.Body)
	if err!=nil{
		log.Println(err)
	}
	var input Request
	err = json.Unmarshal(v, &input)
	if err!=nil{
		log.Println("json 字节数组转map时出现错误:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println("body解析结果是: ", input)
	fmt.Println("path:  ", r.URL.Path)
	fmt.Println("method:  ", r.Method)

	// 处理传入的数据, 计算后将数据返回
	output := &Res{strconv.Itoa(input.Code), []string{},input.Property}
	jsData, err := json.Marshal(output) //将map处理成byte数组
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", "go net")
	w.WriteHeader(200)
	_, _ = w.Write(jsData)
}

// 从slice里面随机选择一个字符串
func randStatement() string {
	//statements := []string{
	//	"happy to see U :)",
	//	"Nice to meet U :)",
	//	"你好 :)",
	//}
	rand.Seed(time.Now().UnixNano())
	statements := make([]string, 10)
	statements[0] = "hello"
	statements[1] = "ko"
	statements[2] = "cc"
	n := rand.Intn(3)
	fmt.Println("函数的返回值是:", statements[n])
	return statements[n]
}
