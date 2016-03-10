package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var doneCh = make(chan struct{})

func main() {
	go burn()

	ch := make(chan os.Signal)

	signal.Notify(ch,
		syscall.SIGINT,
		syscall.SIGKILL,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGSTOP,
	)

	fmt.Printf("Syscall: %s\n", <-ch)
	die()
}

func burn() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("tick")
		case <-doneCh:
			return
		}
	}
}

func die() {
	close(doneCh)
}
