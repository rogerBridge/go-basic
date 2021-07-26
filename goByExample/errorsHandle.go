package main

import (
	"errors"
	"fmt"
)

func meet1(a int) (int, error){
	if a == 1 {
		return -1, errors.New("a could not be 1 :)")
	}
	return a, nil
}

func main() {
	meet1(2)
	_, err := meet1(1)
	if err!=nil {
		fmt.Println(err.Error())
	}
}
