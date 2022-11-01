package utils

import (
	"os"
	"os/signal"
	"syscall"

	"iamza-verifier/pkg/log"
)

func GracefulServerExit(callback func()) {
	s := make(chan os.Signal, 1)
	terminate := make(chan bool, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-s // used for warning handling
		log.ServerError.Printf("Exited unexpectedly due to: %s", sig)
		terminate <- true
	}()

	<-terminate
	callback()
	log.ServerInfo.Print("Terminating on ctrl+c request...")
}
