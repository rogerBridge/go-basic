package main

import (
	"fmt"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"log"
	"os"
	"time"
)

const (
	workMinutes  = 25
	relaxMinutes = 5
)

type TestA struct {
	Name string
	Age  int
}

func main() {
	//runtime.GOMAXPROCS(1)
	log.Printf("%+v", deferAndReturn(0)) // first normal code, then defer, finally return
}

func printS(i int) {
	fmt.Println(i)
}

func deferAndReturn(i int) int {
	defer func() {
		i += 1
		fmt.Printf("in defer %+v\n", i)
	}()
	return i
}
func openMusic(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		log.Println(err)
		return err
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Println(err)
		return err
	}
	defer streamer.Close()
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		log.Println("Init mp3 error happend!")
		return err
	}
	speaker.Play(streamer)
	done := make(chan bool)
	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()
	<-done
	return nil
}

func A() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
