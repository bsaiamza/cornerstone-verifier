package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	acapy "cornerstone_verifier/pkg/acapy_client"
	"cornerstone_verifier/pkg/config"
	"cornerstone_verifier/pkg/log"
	"cornerstone_verifier/pkg/models"
	"cornerstone_verifier/pkg/server"
	"cornerstone_verifier/pkg/util"

	"github.com/skip2/go-qrcode"
)

func displayProofRequest(config *config.Config, acapyClient *acapy.Client, cache *util.BigCache) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(displayProofRequestHandler(config, acapyClient, cache), mdw...)
}
func displayProofRequestHandler(config *config.Config, acapyClient *acapy.Client, cache *util.BigCache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Add("Access-Control-Allow-Origin", config.GetClientURL())
		header.Add("Access-Control-Allow-Methods", "POST, OPTIONS")
		header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodPost {
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

		log.Info.Println("Creating proof request...")

		var data models.ProofRequestData
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			log.Error.Printf("Failed to decode credential data: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to decode credential data: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		// Step 1: Create invitation
		request := models.CreateInvitationRequest{}

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

		invitation, err := acapyClient.CreateInvitation(request, &queryParams)
		if err != nil {
			log.Error.Printf("Failed to prepare proof data: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to prepare proof data: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		// Step 2: Cache user data
		prepareProofData := models.ProofRequestData{
			IDNumber:       data.IDNumber,
			Surname:        data.Surname,
			Forenames:      data.Forenames,
			Gender:         data.Gender,
			DateOfBirth:    data.DateOfBirth,
			CountryOfBirth: data.CountryOfBirth,
		}

		err = cache.UpdateDataCache(invitation.Invitation.RecipientKeys[0], prepareProofData)
		if err != nil {
			log.Error.Printf("Failed to cache presentation data: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to cache presentation data: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		log.Info.Println("Proof request created!")

		w.WriteHeader(http.StatusOK)
		res := server.Res{
			"success":      true,
			"proofRequest": invitation.InvitationURL,
		}
		json.NewEncoder(w).Encode(res)
	}
}

func emailProofRequest(config *config.Config, acapyClient *acapy.Client, cache *util.BigCache) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(emailProofRequestHandler(config, acapyClient, cache), mdw...)
}
func emailProofRequestHandler(config *config.Config, acapyClient *acapy.Client, cache *util.BigCache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Add("Access-Control-Allow-Origin", config.GetClientURL())
		header.Add("Access-Control-Allow-Methods", "POST, OPTIONS")
		header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodPost {
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

		log.Info.Println("Create proof request...")

		var data models.ProofRequestData
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			log.Error.Printf("Failed to decode credential data: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to decode credential data: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		// Step 1: Create invitation
		request := models.CreateInvitationRequest{}

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

		invitation, err := acapyClient.CreateInvitation(request, &queryParams)
		if err != nil {
			log.Error.Printf("Failed to prepare proof data: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to prepare proof data: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		// Step 2: Generate qr code
		qrCodePng, err := qrcode.Encode(invitation.InvitationURL, qrcode.Medium, 256)
		if err != nil {
			log.Warning.Print("Failed to create QR code: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to create QR code: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		// Step 3: Cache user data
		prepareProofData := models.ProofRequestData{
			IDNumber:       data.IDNumber,
			Surname:        data.Surname,
			Forenames:      data.Forenames,
			Gender:         data.Gender,
			DateOfBirth:    data.DateOfBirth,
			CountryOfBirth: data.CountryOfBirth,
			Email:          data.Email,
		}

		err = cache.UpdateDataCache(invitation.Invitation.RecipientKeys[0], prepareProofData)
		if err != nil {
			log.Error.Printf("Failed to cache presentation data: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to cache presentation data: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		// Step 4: Send email
		err = util.SendProofRequestEmail(data.Email, invitation.Invitation.RecipientKeys[0], qrCodePng)
		if err != nil {
			log.Warning.Print("Failed to send credential email: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to send credential email: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		err = os.Remove("./" + invitation.Invitation.RecipientKeys[0] + ".png")
		if err != nil {
			log.Warning.Print("Failed to remove QR code: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to remove QR code: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
		}

		log.Info.Println("Proof request created and sent via email!")

		w.WriteHeader(http.StatusOK)
	}
}

func presentProof(config *config.Config, acapyClient *acapy.Client, cache *util.BigCache) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(presentProofHandler(config, acapyClient, cache), mdw...)
}
func presentProofHandler(config *config.Config, acapyClient *acapy.Client, cache *util.BigCache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Add("Access-Control-Allow-Origin", config.GetClientURL())
		header.Add("Access-Control-Allow-Methods", "POST, OPTIONS")
		header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodPost {
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

		var request models.Connection
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			log.Error.Printf("Fail to decode request body: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			res := server.Res{
				"success": false,
				"msg":     "Failed to decode request body: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		if request.State == "response" {
			_, err := acapyClient.PingConnection(request.ConnectionID)
			if err != nil {
				log.Error.Printf("Failed to ping connection: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				res := server.Res{
					"success": false,
					"msg":     "Failed to ping connection: " + err.Error(),
				}
				json.NewEncoder(w).Encode(res)
				return
			}
		}

		if request.State == "active" {
			credDefID := config.GetCredDefID()
			if credDefID == "" {
				credDefID = "BER7WwiAMK9igkiRjPYpEp:3:CL:40479:cornerstone_1.2"
			}

			proofData, err := cache.ReadDataCache(request.InvitationKey)
			if err != nil {
				log.Error.Printf("Failed to read proof data: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				res := server.Res{
					"success": false,
					"msg":     "Failed to read proof data: " + err.Error(),
				}
				json.NewEncoder(w).Encode(res)
				return
			}

			var idProof interface{}
			var namesProof interface{}
			var surnameProof interface{}
			var genderProof interface{}
			var dobProof interface{}
			var cobProof interface{}

			if proofData.IDNumber {
				idProof = map[string]interface{}{
					"name": "IDNumber",
					"restrictions": []map[string]interface{}{
						{
							"cred_def_id": credDefID,
						},
					},
				}
			} else {
				idProof = nil
			}

			if proofData.Forenames {
				namesProof = map[string]interface{}{
					"name": "Forenames",
					"restrictions": []map[string]interface{}{
						{
							"cred_def_id": credDefID,
						},
					},
				}
			} else {
				namesProof = nil
			}

			if proofData.Surname {
				surnameProof = map[string]interface{}{
					"name": "Surname",
					"restrictions": []map[string]interface{}{
						{
							"cred_def_id": credDefID,
						},
					},
				}
			} else {
				surnameProof = nil
			}

			if proofData.Gender {
				genderProof = map[string]interface{}{
					"name": "Gender",
					"restrictions": []map[string]interface{}{
						{
							"cred_def_id": credDefID,
						},
					},
				}
			} else {
				genderProof = nil
			}

			if proofData.DateOfBirth {
				dobProof = map[string]interface{}{
					"name": "DateOfBirth",
					"restrictions": []map[string]interface{}{
						{
							"cred_def_id": credDefID,
						},
					},
				}
			} else {
				dobProof = nil
			}

			if proofData.CountryOfBirth {
				cobProof = map[string]interface{}{
					"name": "CountryOfBirth",
					"restrictions": []map[string]interface{}{
						{
							"cred_def_id": credDefID,
						},
					},
				}
			} else {
				cobProof = nil
			}

			// values := &models.ProofRequest{
			// 	Comment:      "Proof Request",
			// 	ConnectionID: request.ConnectionID,
			// 	PresentationRequest: models.PresentationRequest{
			// 		Indy: models.Indy{
			// 			Name:    "Proof of Identity",
			// 			Version: "1.0",
			// 			RequestedAttributes: models.RequestedAttributes{
			// 				idProof,
			// 				namesProof,
			// 				surnameProof,
			// 				genderProof,
			// 				dobProof,
			// 				cobProof,
			// 			},
			// 			RequestedPredicates: models.RequestedPredicates{},
			// 		},
			// 	},
			// }

			values := &models.ProofRequestV1{
				Comment:      "Proof Request",
				ConnectionID: request.ConnectionID,
				PresentationRequest: models.PresentationRequestV1{
					Name:    "Proof of Identity",
					Version: "1.0",
					RequestedAttributes: models.RequestedAttributes{
						idProof,
						namesProof,
						surnameProof,
						genderProof,
						dobProof,
						cobProof,
					},
					RequestedPredicates: models.RequestedPredicates{},
				},
			}

			json_data, err := json.Marshal(values)
			if err != nil {
				log.Error.Printf("Failed to marshal proof request: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				res := server.Res{
					"success": false,
					"msg":     "Failed to marshal proof request: " + err.Error(),
				}
				json.NewEncoder(w).Encode(res)
				return
			}

			resp, err := http.Post(config.GetAcapyURL()+"/present-proof/send-request", "application/json", bytes.NewBuffer(json_data))
			// resp, err := http.Post(config.GetAcapyURL()+"/present-proof-2.0/send-request", "application/json", bytes.NewBuffer(json_data))
			if err != nil {
				log.Error.Printf("Failed to send proof request: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				res := server.Res{
					"success": false,
					"msg":     "Failed to send proof request: " + err.Error(),
				}
				json.NewEncoder(w).Encode(res)
				return
			}
			defer resp.Body.Close()

			log.Info.Println("Proof request sent successfully!")

			w.WriteHeader(http.StatusOK)
			res := server.Res{
				"success": true,
				"msg":     "Proof request sent successfully!",
			}
			json.NewEncoder(w).Encode(res)
			return
		}
	}
}

func listProofRecords(config *config.Config, acapyClient *acapy.Client) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(listProofRecordsHandler(config, acapyClient), mdw...)
}
func listProofRecordsHandler(config *config.Config, acapyClient *acapy.Client) http.HandlerFunc {
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

		log.Info.Println("Listing presentations...")

		connectionID := r.URL.Query().Get("connection_id")
		role := r.URL.Query().Get("role")
		state := r.URL.Query().Get("state")
		threadID := r.URL.Query().Get("thread_id")

		queryParams := models.ListProofRecordsParams{
			ConnectionID: connectionID,
			Role:         role,
			State:        state,
			ThreadID:     threadID,
		}

		proofRecords, err := acapyClient.ListProofRecords(&queryParams)
		if err != nil {
			log.Error.Printf("Failed to list proof requests: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to list proof requests: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		log.Info.Print("Proof requests listed successfully!")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(proofRecords.Results)
	}
}
