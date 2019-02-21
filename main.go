package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type (
	//DimaChan ne Iaroslav
	DimaChan struct {
		ch   chan string
		name string
	}
)

var (
	chanels = [3]chan string{}
	colors  = map[int]string{0: "GREEN", 1: "YELLOW", 2: "RED"}
)

func prnt(c chan string, lightName string) {
	for {
		fmt.Println(lightName, ":", <-c)

	}
}

func snd(c chan string, t int) {
	for range time.Tick(time.Duration(t) * (time.Second)) {
		c <- colors[rand.Intn(3)]
	}
}
func main() {
	rand.Seed(time.Now().UnixNano())

	eCh := make(chan int)

	name := "FIRST"
	for i := 0; i < 3; i++ {
		switch i {
		case 1:
			name = "SECOND"
		case 2:
			name = "THIRD"
		}
		chanels[i] = make(chan string)
		go snd(chanels[i], i)
		go prnt(chanels[i], name)
	}
	os.Exit(<-eCh)
}
