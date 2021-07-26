package main

import "os"

func main() {
	//panic("A problem")

	_, err := os.Create("/root/hello")
	if err!=nil {
		panic(err)
	}
}
