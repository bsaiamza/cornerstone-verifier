package server

import (
	"os"
	"os/signal"
	"syscall"

	"cornerstone_verifier/pkg/log"
)

// GracefulExit listens for program termination.
func GracefulExit(cb func()) {
	sigs := make(chan os.Signal, 1)
	terminate := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		log.ServerError.Printf("Exited unexpectedly due to: %s", sig)
		terminate <- true
	}()

	<-terminate
	cb()
	log.ServerInfo.Print("Terminating on ctrl+c request...")
}
