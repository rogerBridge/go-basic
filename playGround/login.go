package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"time"
)

// 下面是一个没有考虑异常情况的
func login(user string, pass string) string {
	// 构造请求
	url := "http://gateway.smart.lierda.com/auth"
	requestBody, err := json.Marshal(map[string]interface{}{
		"username": user,
		"password": pass,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(requestBody, reflect.TypeOf(requestBody))

	// 发送请求给服务器&&处理服务器返回的请求
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		// 捕捉这个错误并且中止程序
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	type tokenData struct {
		User  string `json:"user"`
		Token string `json:"token"`
	}
	type loginData struct {
		Code string    `json:"code"`
		Data tokenData `json:"data"`
		Msg  string    `json:"msg"`
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	//var objmap map[string]*json.RawMessage
	objmap := new(loginData)
	json.Unmarshal(respBody, &objmap)
	// 输出objmap结构体里面的token的👈值
	fmt.Println(objmap.Data.Token)
	return objmap.Data.Token
}

// 详细的http方法
func login2(user string, pass string, ch chan int) {
	// 1 表示正常, 0 表示异常
	fmt.Println("开始时间是: ", time.Now())
	// 构造http post 请求
	reqBody, err := json.Marshal(map[string]interface{}{
		"username": user,
		"password": pass,
	})
	if err != nil {
		log.Fatalln(err)
	}
	url := "http://gateway.smart.lierda.com/auth"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatalln("构造请求时出现错误", err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	// 发送请求给服务器&&处理服务器返回的请求
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("发送请求给服务器时出现错误", err)
		ch <- 0
	}
	if resp.StatusCode == 200 {
		log.Println("服务区返回200状态码")
		ch <- 1
	} else {
		log.Println("服务器返回其他状态码", resp.StatusCode)
		ch <- 0
	}
	//respBody, _ := ioutil.ReadAll(resp.Body) //从resp.Body里面将[]byte 解析出来
	//defer resp.Body.Close()
	//var objmap map[string]*json.RawMessage
	//err = json.Unmarshal(respBody, &objmap)
	//if err != nil {
	//	log.Println("解析json数据时出了问题", err)
	//	ch <- 0
	//}
	//type data struct {
	//	User  string `json:"user"`
	//	Token string `json:"token"`
	//}
	//var userData data
	//// *objmap["data"] 就是 json.RawMessage
	//err = json.Unmarshal(*objmap["data"], &userData)
	//if err != nil {
	//	log.Println("解析json[data]的时候出了问题", err)
	//	ch <- 0
	//}
	//token := userData.Token
	//if token != "" {
	//	ch <- 1
	//} else {
	//	ch <- 0
	//}
}

func getLocalhost() {
	// 构造http请求
	req, err := http.NewRequest("GET", "http://127.0.0.1/get", nil)
	if err != nil {
		log.Fatalln(err)
	}
	// 给req的header里面添加键值对
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "Go net/http")
	// 给req的url里面添加键值对
	q := req.URL.Query()
	q.Add("habit", "basketball")
	q.Add("lover", "Yuan")
	req.URL.RawQuery = q.Encode()
	fmt.Println("构造的get请求的url是:", req.URL)

	// 发送http请求并且接收从服务端返回的数据
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("返回的字节数组是:", respBody) // 返回字节数组类型的body
	//fmt.Println(string(respBody))
}

func customer(ch chan int, concurrentNum int) {
	//r := make(chan int, 1000)
	r := make([]int, concurrentNum)
	for i := 0; i < concurrentNum; i++ {
		r[i] = <-ch
	}
	fmt.Printf("并发%d, 成功数%d, 成功率%f", len(r), sum(r), float64(sum(r))/float64(len(r)))
}

func sum(s []int) int {
	r := 0
	for _, v := range s {
		r += v
	}
	return r
}

func producer(ch chan int, concurrentNum int) {
	for i := 0; i < concurrentNum; i++ {
		go login2("anruifeng", "sOk+qRNn2DirShe0j7lK+g==", ch)
	}
}

func main() {
	concurrentNum := 100
	ch := make(chan int, concurrentNum)
	go producer(ch, concurrentNum)
	go customer(ch, concurrentNum)
	time.Sleep(10*time.Second)
	//fmt.Scanln()
	//fmt.Println("done")
	//for i := 0; i < 1000; i++ {
	//	go login2("anruifeng", "sOk+qRNn2DirShe0j7lK+g==")
	//}
	//fmt.Scanln()
}
