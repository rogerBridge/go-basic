package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	requestBody, err := json.Marshal(map[string]interface{}{
		"username": "anruifeng",
		"password": "sOk+qRNn2DirShe0j7lK+g==",
		"code":     123456,
	})
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Println("json化之后的字符串是: ", requestBody)

	url := "http://gateway.smart.lierda.com/auth"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	//fmt.Println(resp)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var objmap map[string]*json.RawMessage
	// 下面的if判断, 前面是一个赋值语句, 后面才是判断
	err = json.Unmarshal(body, &objmap)
	if err != nil {
		log.Println(err)
	}
	type data struct { // 一定要大写哦 :)
		User  string
		Token string
	}
	fmt.Println(objmap)
	var userData data
	err = json.Unmarshal(*objmap["data"], &userData)
	token := userData.Token
	fmt.Println(token)
}
