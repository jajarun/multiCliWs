package main

import (
	"fmt"
	"time"
)

var testChan = make(chan string)

func main() {
	go func() {
		time.Sleep(time.Second * 1)
		testChan <- "test"
	}()
	for {
		select {
		case msg := <-testChan:
			// 建立连接事件
			fmt.Println(msg)
		default:
			fmt.Println("no message")
		}
	}
}
