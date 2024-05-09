package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	fmt.Println("start")
	interrupt := make(chan os.Signal, 1)

	signal.Notify(interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-interrupt

	fmt.Println("succeded")
	time.Sleep(10 * time.Second)
}
