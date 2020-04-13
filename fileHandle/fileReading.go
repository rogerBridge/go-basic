package fileHandle

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func ReadingFile() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("cannot get current directory")
	}
	fmt.Println(pwd)
	data, err := ioutil.ReadFile(pwd + "/fileHandle/test.txt")
	if err != nil {
		log.Println("read file failed!")
		return
	}
	fmt.Println("content of file:", string(data))
	os.Create("hello")
}
