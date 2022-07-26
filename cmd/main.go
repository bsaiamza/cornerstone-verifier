package main

import (
	"iamza_verifier/api"
	"iamza_verifier/pkg/client"
	"iamza_verifier/pkg/config"
	"iamza_verifier/pkg/log"
	"iamza_verifier/pkg/server"
	"iamza_verifier/pkg/utils"
)

func main() {
	runServer()
}

func runServer() {
	config := config.LoadConfig()
	acapyClient := client.NewClient(config.GetAcapyURL())
	cache := utils.NewBigCache()

	srv := server.NewServer().
		WithAddress(config.GetServerAddress()).
		WithRouter(api.NewRouter(config, acapyClient, cache)).
		WithErrorLogger(log.ServerError)

	go func() {
		log.ServerInfo.Println("-----------------------------------------")
		log.ServerInfo.Println("|		IAMZA Verifier		|")
		log.ServerInfo.Println("-----------------------------------------")
		log.ServerInfo.Println("")
		log.ServerInfo.Printf("Server started on: %s", config.GetServerAddress())
		if err := srv.Start(); err != nil {
			log.ServerError.Fatal(err)
		}
	}()

	utils.GracefulServerExit(func() {
		if err := srv.Stop(); err != nil {
			log.ServerError.Printf("Failed to stop server: %s", err.Error())
		}
	})
}
