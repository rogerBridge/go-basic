package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	formData := url.Values{"name": []string{"kahn"}}
	resp, err := http.PostForm("http://localhost/post", formData)
	if err!=nil{
		log.Fatalln(err)
	}
	var result map[string]interface{}
	d, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(d, &result)
	log.Println(string(d))
}
