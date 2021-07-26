package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func FailOnError(err error, msg string) {
	if err!=nil{
		log.Fatalln(err, msg)
	}
}

func main() {
	resp, err := http.Get("http://localhost/get")
	if err!=nil{
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		log.Fatalln(err)
	}
	log.Println(string(body))
}
