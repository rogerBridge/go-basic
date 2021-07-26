package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//fmt.Println(len([]byte("安")), []byte("安瑞峰"))
	reqBody, err := json.Marshal(map[string]string{
		"name": "安 瑞峰",
		"email": "ttxs.an@gmail.com",
	})
	if err!=nil{
		log.Fatalln(err)
	}

	resp, err := http.Post("http://localhost/post", "application/json", bytes.NewBuffer(reqBody))
	if err!=nil{
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		log.Fatalln(body)
	}
	log.Println(string(body))

}
