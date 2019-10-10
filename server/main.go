package main

import (
	"github.com/er1c-zh/go-now/log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	ToQuit = make(chan struct{})
	FinishQuit = make(chan struct{})
)

func main() {
	defer log.Flush()
	log.Info("xcross init...")

	exitChan := make(chan os.Signal)
	signal.Notify(exitChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		t := time.NewTicker(time.Second)

		for {
			select {
			case <-t.C:
				log.Info("tick...")
			case <-ToQuit:
				log.Info("quit tick.")
				FinishQuit <- struct{}{}
				return
			}
		}
	}()

	signal := <-exitChan
	log.Info("try to quit gracefully, by %+v", signal)

	ToQuit <- struct{}{}

	<-FinishQuit
	log.Info("quit gracefully!")
}
