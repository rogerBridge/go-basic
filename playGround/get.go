package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	// 构造http get 请求
	req, err := http.NewRequest("GET", "http://127.0.0.1/get", nil)
	if err!=nil{
		log.Fatalln("制造req时出现错误", err)
	}
	req.Header.Add("User-Agent", "Chrome")
	// 构造params
	q := req.URL.Query()
	//fmt.Println(reflect.TypeOf(q), reflect.ValueOf(q))
	q.Add("key", "value")
	req.URL.RawQuery = q.Encode()

	client := &http.Client{Timeout: 5*time.Second}
	// 发送http get 请求
	resp, err := client.Do(req)
	if err!=nil{
		log.Fatalln(err)
	}
	//resp.Body 是io.ReadCloser流, 需要使用ioutil.ReadAll方法将其转为[]byte
	respBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(string(respBody))
}
