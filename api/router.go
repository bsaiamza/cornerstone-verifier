package api

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"iamza_verifier/pkg/client"
	"iamza_verifier/pkg/config"
	"iamza_verifier/pkg/utils"

	"github.com/gorilla/mux"
)

//go:embed build
var embeddedFiles embed.FS

func NewRouter(config *config.Config, client *client.Client, cache *utils.BigCache) *mux.Router {
	r := mux.NewRouter()

	path := r.PathPrefix("/api/v1/iamza-verifier").Subrouter()
	path.HandleFunc("/health", health(config))
	path.HandleFunc("/connections", listConnections(config, client))
	path.HandleFunc("/verification-records", listVerificationRecords(config, client))
	path.HandleFunc("/verify", verifyCredential(config, client, cache))
	path.HandleFunc("/verify-email", verifyCredentialByEmail(config, client, cache))
	path.HandleFunc("/verify-cornerstone", verifyCornerstoneCredential(config, client, cache))
	path.HandleFunc("/verify-cornerstone-email", verifyCornerstoneCredentialByEmail(config, client, cache))
	path.HandleFunc("/topic/{topic}/", webhookEvents(config, client, cache))

	r.PathPrefix("/").Handler(http.FileServer(getFileSystem()))

	return r
}

func getFileSystem() http.FileSystem {
	fsys, err := fs.Sub(embeddedFiles, "build")
	if err != nil {
		fmt.Println(err)
	}

	return http.FS(fsys)
}
