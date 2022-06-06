package main

import (
	"cornerstone_verifier/api/v1"
	"cornerstone_verifier/pkg/config"
	"cornerstone_verifier/pkg/log"
	"cornerstone_verifier/pkg/server"
)

func main() {
	config := config.GetConfig()

	serverAddress := config.GetServerAddress()

	srv := server.NewServer().
		WithAddress(serverAddress).
		WithRouter(api.NewRouter(config)).
		WithErrLogger(log.ServerError)

	go func() {
		log.ServerInfo.Println("-------------------------------------------------")
		log.ServerInfo.Println("|		Cornerstone Verifier		|")
		log.ServerInfo.Println("-------------------------------------------------")
		log.ServerInfo.Println("		**ENV VARS**")
		log.ServerInfo.Println("	CLIENT_URL: ", config.GetClientURL())
		log.ServerInfo.Println("	SERVER_ADDRESS: ", config.GetServerAddress())
		log.ServerInfo.Println("	API_BASE_URL: ", config.GetAPIBaseURL())
		log.ServerInfo.Println("-------------------------------------------------")
		log.ServerInfo.Println("")
		log.ServerInfo.Printf("Server started on: %s", serverAddress)
		if err := srv.Start(); err != nil {
			log.ServerError.Fatal(err)
		}
	}()

	server.GracefulExit(func() {
		if err := srv.Stop(); err != nil {
			log.ServerError.Printf("Failed to stop server: %s", err.Error())
		}
	})
}
