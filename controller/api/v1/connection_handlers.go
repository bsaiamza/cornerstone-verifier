package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	acapy "cornerstone_verifier/pkg/acapy_client"
	"cornerstone_verifier/pkg/config"
	"cornerstone_verifier/pkg/log"
	"cornerstone_verifier/pkg/models"
	"cornerstone_verifier/pkg/server"
)

func invitation(config *config.Config, acapyClient *acapy.Client) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(invitationHandler(config, acapyClient), mdw...)
}
func invitationHandler(config *config.Config, acapyClient *acapy.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Add("Access-Control-Allow-Origin", config.GetClientURL())
		header.Add("Access-Control-Allow-Methods", "GET, OPTIONS")
		header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodGet {
			log.Warning.Print("Incorrect request method!")
			w.WriteHeader(http.StatusMethodNotAllowed)
			res := server.Res{
				"success": false,
				"msg":     "Warning: Incorrect request method!",
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		defer r.Body.Close()

		log.Info.Println("Generating invitation...")

		alias := r.URL.Query().Get("alias")
		autoAccept, _ := strconv.ParseBool(r.URL.Query().Get("auto_accept"))
		multiuse, _ := strconv.ParseBool(r.URL.Query().Get("multi_use"))
		public, _ := strconv.ParseBool(r.URL.Query().Get("public"))

		queryParams := models.CreateInvitationParams{
			Alias:      alias,
			AutoAccept: autoAccept,
			MultiUse:   multiuse,
			Public:     public,
		}

		var createInvitationRequest models.CreateInvitationRequest
		err := json.NewDecoder(r.Body).Decode(&createInvitationRequest)
		switch {
		case err == io.EOF:
			log.Warning.Print("Empty body!")
		case err != nil:
			log.Warning.Print("Failed to parse body: ", err)
			w.WriteHeader(http.StatusBadRequest)
			res := server.Res{
				"success": false,
				"msg":     "Failed to parse body: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		invitation, err := acapyClient.CreateInvitation(createInvitationRequest, &queryParams)
		if err != nil {
			log.Error.Printf("Failed to generate invitation: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to generate invitation: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		log.Info.Print("Generated invitation successfully!")

		w.WriteHeader(http.StatusOK)
		res := server.Res{
			"success":       true,
			"msg":           "Generated invitation successfully!",
			"invitation":    invitation.Invitation,
			"invitationURL": invitation.InvitationURL,
		}
		json.NewEncoder(w).Encode(res)
	}
}

func listConnections(config *config.Config, acapyClient *acapy.Client) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(listConnectionsHandler(config, acapyClient), mdw...)
}
func listConnectionsHandler(config *config.Config, acapyClient *acapy.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Add("Access-Control-Allow-Origin", config.GetClientURL())
		header.Add("Access-Control-Allow-Methods", "GET, OPTIONS")
		header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodGet {
			log.Warning.Print("Incorrect request method!")
			w.WriteHeader(http.StatusMethodNotAllowed)
			res := server.Res{
				"success": false,
				"msg":     "Warning: Incorrect request method!",
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		defer r.Body.Close()

		log.Info.Println("Listing connections...")

		alias := r.URL.Query().Get("alias")
		connectionProtocol := r.URL.Query().Get("connection_protocol")
		invitationKey := r.URL.Query().Get("invitation_key")
		myDID := r.URL.Query().Get("my_did")
		state := r.URL.Query().Get("state")
		theirDID := r.URL.Query().Get("their_did")
		theirPublicDID := r.URL.Query().Get("their_public_did")
		theirRole := r.URL.Query().Get("their_role")

		queryParams := models.ListConnectionsParams{
			Alias:              alias,
			ConnectionProtocol: connectionProtocol,
			InvitationKey:      invitationKey,
			MyDID:              myDID,
			State:              state,
			TheirDID:           theirDID,
			TheirPublicDID:     theirPublicDID,
			TheirRole:          theirRole,
		}

		connections, err := acapyClient.ListConnections(&queryParams)
		if err != nil {
			log.Error.Printf("Failed to query connections: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to query connections: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		log.Info.Print("Connections listed successfully!")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(connections.Results)
	}
}
