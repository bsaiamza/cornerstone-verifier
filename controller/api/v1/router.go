package api

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	acapy "cornerstone_verifier/pkg/acapy_client"
	"cornerstone_verifier/pkg/config"
	"cornerstone_verifier/pkg/util"
)

//go:embed build
var embeddedFiles embed.FS

func NewRouter(config *config.Config, acapyClient *acapy.Client, cache *util.BigCache) *http.ServeMux {
	r := http.NewServeMux()

	apiBaseURL := config.GetAPIBaseURL()

	// health
	r.HandleFunc(apiBaseURL+"/cornerstone/verifier/health", health(config))
	// connection
	r.HandleFunc(apiBaseURL+"/cornerstone/verifier/connection/invitation", invitation(config, acapyClient))
	r.HandleFunc(apiBaseURL+"/cornerstone/verifier/connections", listConnections(config, acapyClient))
	// proof
	r.HandleFunc(apiBaseURL+"/cornerstone/verifier/topic/connections/", presentProof(config, acapyClient, cache))
	r.HandleFunc(apiBaseURL+"/cornerstone/verifier/proof", displayProofRequest(config, acapyClient, cache))
	r.HandleFunc(apiBaseURL+"/cornerstone/verifier/email-proof", emailProofRequest(config, acapyClient, cache))
	r.HandleFunc(apiBaseURL+"/cornerstone/verifier/presentations", listProofRecords(config, acapyClient))

	r.Handle("/", http.FileServer(getFileSystem()))

	return r
}

func getFileSystem() http.FileSystem {
	// Get the build subdirectory as the
	// root directory so that it can be passed
	// to the http.FileServer
	fsys, err := fs.Sub(embeddedFiles, "build")
	if err != nil {
		fmt.Println(err)
	}

	return http.FS(fsys)
}
