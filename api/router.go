package api

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"iamza-verifier/pkg/acapy"
	"iamza-verifier/pkg/config"
	"iamza-verifier/pkg/utils"

	"github.com/gorilla/mux"
)

//go:embed build
var embeddedFiles embed.FS

func NewRouter(config *config.Config, acapy *acapy.Client, cache *utils.BigCache) *mux.Router {
	r := mux.NewRouter()

	path := r.PathPrefix("/api/v2/iamza-verifier").Subrouter()
	path.HandleFunc("/verify-cornerstone", verifyCornerstoneCredential(config, acapy, cache))
	path.HandleFunc("/verify-cornerstone-email", verifyCornerstoneCredentialByEmail(config, acapy, cache))
	path.HandleFunc("/verify-contactable", verifyContactableCredential(config, acapy, cache))
	path.HandleFunc("/verify-contactable-email", verifyContactableCredentialByEmail(config, acapy, cache))
	path.HandleFunc("/verify-address", verifyAddressCredential(config, acapy, cache))
	path.HandleFunc("/verify-address-email", verifyAddressCredentialByEmail(config, acapy, cache))
	path.HandleFunc("/verify-vaccine", verifyVaccineCredential(config, acapy, cache))
	path.HandleFunc("/verify-vaccine-email", verifyVaccineCredentialByEmail(config, acapy, cache))
	path.HandleFunc("/topic/{topic}/", webhookEvents(config, acapy, cache))

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
