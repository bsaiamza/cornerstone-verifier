package api

import (
	"encoding/json"
	"net/http"
	"os"

	"iamza_verifier/pkg/client"
	"iamza_verifier/pkg/config"
	"iamza_verifier/pkg/log"
	"iamza_verifier/pkg/models"
	"iamza_verifier/pkg/server"
	"iamza_verifier/pkg/utils"

	"github.com/gorilla/mux"
	"github.com/skip2/go-qrcode"
)

func health(config *config.Config) http.HandlerFunc {
	mdw := []server.Middleware{
		server.LogAPIRequest,
	}

	return server.ChainMiddleware(healthHandler(config), mdw...)
}
func healthHandler(config *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
}

func listConnections(config *config.Config, client *client.Client) http.HandlerFunc {
	mdw := []server.Middleware{
		server.LogAPIRequest,
	}

	return server.ChainMiddleware(listConnectionsHandler(config, client), mdw...)
}
func listConnectionsHandler(config *config.Config, client *client.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Add("Access-Control-Allow-Origin", "*")
		header.Add("Access-Control-Allow-Methods", "GET, OPTIONS")
		header.Add("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodGet {
			log.Warning.Print("Incorrect request method!")
			w.WriteHeader(http.StatusMethodNotAllowed)
			res := server.Response{
				"success": false,
				"msg":     "Warning: Incorrect request method!",
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		defer r.Body.Close()

		log.Info.Println("Listing connections...")

		connections, err := client.ListConnections()
		if err != nil {
			log.Error.Printf("Failed to list connections: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Response{
				"success": false,
				"msg":     "Failed to list connections: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		log.Info.Print("Connections listed successfully!")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(connections.Results)
	}
}

func listVerificationRecords(config *config.Config, client *client.Client) http.HandlerFunc {
	mdw := []server.Middleware{
		server.LogAPIRequest,
	}

	return server.ChainMiddleware(listVerificationRecordsHandler(config, client), mdw...)
}
func listVerificationRecordsHandler(config *config.Config, client *client.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Add("Access-Control-Allow-Origin", "*")
		header.Add("Access-Control-Allow-Methods", "GET, OPTIONS")
		header.Add("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodGet {
			log.Warning.Print("Incorrect request method!")
			w.WriteHeader(http.StatusMethodNotAllowed)
			res := server.Response{
				"success": false,
				"msg":     "Warning: Incorrect request method!",
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		defer r.Body.Close()

		log.Info.Println("Listing verification records...")

		records, err := client.ListVerificationRecords()
		if err != nil {
			log.Error.Printf("Failed to list verification records: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Response{
				"success": false,
				"msg":     "Failed to list verification records: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		log.Info.Print("Verification records listed successfully!")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(records.Results)
	}
}

func verifyCredential(config *config.Config, client *client.Client, cache *utils.BigCache) http.HandlerFunc {
	mdw := []server.Middleware{
		server.LogAPIRequest,
	}

	return server.ChainMiddleware(verifyCredentialHandler(config, client, cache), mdw...)
}
func verifyCredentialHandler(config *config.Config, client *client.Client, cache *utils.BigCache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Add("Access-Control-Allow-Origin", "*")
		header.Add("Access-Control-Allow-Methods", "GET, OPTIONS")
		header.Add("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodGet {
			log.Warning.Print("Incorrect request method!")
			w.WriteHeader(http.StatusMethodNotAllowed)
			res := server.Response{
				"success": false,
				"msg":     "Warning: Incorrect request method!",
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		defer r.Body.Close()

		log.Info.Println("Creating proof request...")

		// Step 1: Create Invitation
		invitationRequest := models.CreateInvitationRequest{}

		invitation, err := client.CreateInvitation(invitationRequest)
		if err != nil {
			log.Error.Printf("Failed to create invitation: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Response{
				"success": false,
				"msg":     "Failed to create invitation: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		// Step 2: Cache IAMZA indicator
		err = cache.UpdateString(invitation.Invitation.RecipientKeys[0]+"IAMZA", "IAMZA proof")
		if err != nil {
			log.Error.Printf("Failed to cache proof data: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Response{
				"success": false,
				"msg":     "Failed to cache proof data: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		w.WriteHeader(http.StatusOK)
		res := server.Response{
			"success":      true,
			"proofRequest": invitation.InvitationURL,
		}
		json.NewEncoder(w).Encode(res)
	}
}

func verifyCredentialByEmail(config *config.Config, client *client.Client, cache *utils.BigCache) http.HandlerFunc {
	mdw := []server.Middleware{
		server.LogAPIRequest,
	}

	return server.ChainMiddleware(verifyCredentialByEmailHandler(config, client, cache), mdw...)
}
func verifyCredentialByEmailHandler(config *config.Config, client *client.Client, cache *utils.BigCache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Add("Access-Control-Allow-Origin", "*")
		header.Add("Access-Control-Allow-Methods", "POST, OPTIONS")
		header.Add("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodPost {
			log.Warning.Print("Incorrect request method!")
			w.WriteHeader(http.StatusMethodNotAllowed)
			res := server.Response{
				"success": false,
				"msg":     "Warning: Incorrect request method!",
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		defer r.Body.Close()

		log.Info.Println("Creating proof request...")

		// Step 1: Retrieve proof information
		var proofInfo models.CredentialProofRequest
		err := json.NewDecoder(r.Body).Decode(&proofInfo)
		if err != nil {
			log.Error.Printf("Failed to decode proof request data: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Response{
				"success": false,
				"msg":     "Failed to decode proof request data: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		// Step 2: Validate email address
		err = utils.ValidEmail(proofInfo.Email)
		if err != nil {
			log.Error.Printf("Failed %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Response{
				"success": false,
				"msg":     "Failed " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		// Step 3: Create Invitation
		invitationRequest := models.CreateInvitationRequest{}

		invitation, err := client.CreateInvitation(invitationRequest)
		if err != nil {
			log.Error.Printf("Failed to create invitation: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Response{
				"success": false,
				"msg":     "Failed to create invitation: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		// Step 4: Cache IAMZA indicator
		err = cache.UpdateString(invitation.Invitation.RecipientKeys[0]+"IAMZA", "IAMZA proof")
		if err != nil {
			log.Error.Printf("Failed to cache proof data: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Response{
				"success": false,
				"msg":     "Failed to cache proof data: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		// For email proof notification
		err = cache.UpdateString(invitation.Invitation.RecipientKeys[0]+"email", proofInfo.Email)
		if err != nil {
			log.Error.Printf("Failed to cache proof email: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Response{
				"success": false,
				"msg":     "Failed to cache proof email: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		// Step 5: Generate a qr code for email
		qrCodePng, err := qrcode.Encode(invitation.InvitationURL, qrcode.Medium, 256)
		if err != nil {
			log.Warning.Print("Failed to create QR code: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Response{
				"success": false,
				"msg":     "Failed to create QR code: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		// Step 6: Send email
		err = utils.SendProofRequestByEmail(proofInfo.Email, invitation.Invitation.RecipientKeys[0], qrCodePng, config)
		if err != nil {
			log.Warning.Print("Failed to send proof request by email: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Response{
				"success": false,
				"msg":     "Failed to send proof request by email",
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		// Step 7: Remove qr from os once email is sent
		err = os.Remove("./" + invitation.Invitation.RecipientKeys[0] + ".png")
		if err != nil {
			log.Warning.Print("Failed to remove QR code: ", err)
		}

		w.WriteHeader(http.StatusOK)
	}
}

func verifyCornerstoneCredential(config *config.Config, client *client.Client, cache *utils.BigCache) http.HandlerFunc {
	mdw := []server.Middleware{
		server.LogAPIRequest,
	}

	return server.ChainMiddleware(verifyCornerstoneCredentialHandler(config, client, cache), mdw...)
}
func verifyCornerstoneCredentialHandler(config *config.Config, client *client.Client, cache *utils.BigCache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Add("Access-Control-Allow-Origin", "*")
		header.Add("Access-Control-Allow-Methods", "POST, OPTIONS")
		header.Add("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodPost {
			log.Warning.Print("Incorrect request method!")
			w.WriteHeader(http.StatusMethodNotAllowed)
			res := server.Response{
				"success": false,
				"msg":     "Warning: Incorrect request method!",
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		defer r.Body.Close()

		log.Info.Println("Creating proof request...")

		// Step 1: Retrieve proof information
		var proofInfo models.CornerstoneCredentialProofRequest
		err := json.NewDecoder(r.Body).Decode(&proofInfo)
		if err != nil {
			log.Error.Printf("Failed to decode proof request data: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Response{
				"success": false,
				"msg":     "Failed to decode proof request data: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		// Step 2: Create Invitation
		invitationRequest := models.CreateInvitationRequest{}

		invitation, err := client.CreateInvitation(invitationRequest)
		if err != nil {
			log.Error.Printf("Failed to create invitation: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Response{
				"success": false,
				"msg":     "Failed to create invitation: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		// Step 3: Cache proof data for webhookEventsHandler
		err = cache.UpdateStruct(invitation.Invitation.RecipientKeys[0], proofInfo)
		if err != nil {
			log.Error.Printf("Failed to cache proof data: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Response{
				"success": false,
				"msg":     "Failed to cache proof data: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		w.WriteHeader(http.StatusOK)
		res := server.Response{
			"success":      true,
			"proofRequest": invitation.InvitationURL,
		}
		json.NewEncoder(w).Encode(res)
	}
}

func verifyCornerstoneCredentialByEmail(config *config.Config, client *client.Client, cache *utils.BigCache) http.HandlerFunc {
	mdw := []server.Middleware{
		server.LogAPIRequest,
	}

	return server.ChainMiddleware(verifyCornerstoneCredentialByEmailHandler(config, client, cache), mdw...)
}
func verifyCornerstoneCredentialByEmailHandler(config *config.Config, client *client.Client, cache *utils.BigCache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Add("Access-Control-Allow-Origin", "*")
		header.Add("Access-Control-Allow-Methods", "POST, OPTIONS")
		header.Add("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodPost {
			log.Warning.Print("Incorrect request method!")
			w.WriteHeader(http.StatusMethodNotAllowed)
			res := server.Response{
				"success": false,
				"msg":     "Warning: Incorrect request method!",
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		defer r.Body.Close()

		log.Info.Println("Creating proof request...")

		// Step 1: Retrieve proof information
		var proofInfo models.CornerstoneCredentialProofRequest
		err := json.NewDecoder(r.Body).Decode(&proofInfo)
		if err != nil {
			log.Error.Printf("Failed to decode proof request data: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Response{
				"success": false,
				"msg":     "Failed to decode proof request data: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		// Step 2: Validate email address
		err = utils.ValidEmail(proofInfo.Email)
		if err != nil {
			log.Error.Printf("Failed %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Response{
				"success": false,
				"msg":     "Failed " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		// Step 3: Create Invitation
		invitationRequest := models.CreateInvitationRequest{}

		invitation, err := client.CreateInvitation(invitationRequest)
		if err != nil {
			log.Error.Printf("Failed to create invitation: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Response{
				"success": false,
				"msg":     "Failed to create invitation: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		// Step 5: Cache proof data for webhookEventsHandler
		err = cache.UpdateStruct(invitation.Invitation.RecipientKeys[0], proofInfo)
		if err != nil {
			log.Error.Printf("Failed to cache proof data: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Response{
				"success": false,
				"msg":     "Failed to cache proof data: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		// Step 6: Generate a qr code for email
		qrCodePng, err := qrcode.Encode(invitation.InvitationURL, qrcode.Medium, 256)
		if err != nil {
			log.Warning.Print("Failed to create QR code: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Response{
				"success": false,
				"msg":     "Failed to create QR code: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		// Step 7: Send email
		err = utils.SendProofRequestByEmail(proofInfo.Email, invitation.Invitation.RecipientKeys[0], qrCodePng, config)
		if err != nil {
			log.Warning.Print("Failed to send proof request by email: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Response{
				"success": false,
				"msg":     "Failed to send proof request by email",
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		// Step 8: Remove qr from os once email is sent
		err = os.Remove("./" + invitation.Invitation.RecipientKeys[0] + ".png")
		if err != nil {
			log.Warning.Print("Failed to remove QR code: ", err)
		}

		w.WriteHeader(http.StatusOK)
	}
}

func webhookEvents(config *config.Config, client *client.Client, cache *utils.BigCache) http.HandlerFunc {
	mdw := []server.Middleware{
		server.LogAPIRequest,
	}

	return server.ChainMiddleware(webhookEventsHandler(config, client, cache), mdw...)
}
func webhookEventsHandler(config *config.Config, client *client.Client, cache *utils.BigCache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Add("Access-Control-Allow-Origin", "*")
		header.Add("Access-Control-Allow-Methods", "POST, OPTIONS")
		header.Add("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodPost {
			log.Warning.Print("Incorrect request method!")
			w.WriteHeader(http.StatusMethodNotAllowed)
			res := server.Response{
				"success": false,
				"msg":     "Warning: Incorrect request method!",
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		defer r.Body.Close()

		topic := mux.Vars(r)["topic"]

		switch topic {
		case "connections":
			var request models.Connection
			err := json.NewDecoder(r.Body).Decode(&request)
			if err != nil {
				log.Error.Printf("Failed to decode request body: %s", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if request.State == "response" {
				pingRequest := models.PingConnectionRequest{
					Comment: "Ping",
				}

				_, err := client.PingConnection(request.ConnectionID, pingRequest)
				if err != nil {
					log.Error.Printf("Failed to ping holder: %s", err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}

			iamzaProofRequest, _ := cache.ReadString(request.InvitationKey + "IAMZA")
			if iamzaProofRequest == "IAMZA proof" && request.State == "active" {
				cornerstoneCredDefID := config.GetCornerstoneCredDefID()
				addressCredDefID := config.GetAddressCredDefID()
				vaccineCredDefID := config.GetVaccineCredDefID()

				idProof := map[string]interface{}{
					"name": "ID Number",
					"restrictions": []map[string]interface{}{
						{
							"cred_def_id": cornerstoneCredDefID,
						},
						{
							"cred_def_id": addressCredDefID,
						},
						{
							"cred_def_id": vaccineCredDefID,
						},
					},
				}

				namesProof := map[string]interface{}{
					"name": "First Names",
					"restrictions": []map[string]interface{}{
						{
							"cred_def_id": cornerstoneCredDefID,
						},
						{
							"cred_def_id": addressCredDefID,
						},
						{
							"cred_def_id": vaccineCredDefID,
						},
					},
				}

				surnameProof := map[string]interface{}{
					"name": "Surname",
					"restrictions": []map[string]interface{}{
						{
							"cred_def_id": cornerstoneCredDefID,
						},
						{
							"cred_def_id": addressCredDefID,
						},
						{
							"cred_def_id": vaccineCredDefID,
						},
					},
				}

				genderProof := map[string]interface{}{
					"name": "Gender",
					"restrictions": []map[string]interface{}{
						{
							"cred_def_id": cornerstoneCredDefID,
						},
						{
							"cred_def_id": addressCredDefID,
						},
						{
							"cred_def_id": vaccineCredDefID,
						},
					},
				}

				dobProof := map[string]interface{}{
					"name": "Date of Birth",
					"restrictions": []map[string]interface{}{
						{
							"cred_def_id": cornerstoneCredDefID,
						},
						{
							"cred_def_id": addressCredDefID,
						},
						{
							"cred_def_id": vaccineCredDefID,
						},
					},
				}

				street1Proof := map[string]interface{}{
					"name": "Address Line 1",
					"restrictions": []map[string]interface{}{
						{
							"cred_def_id": cornerstoneCredDefID,
						},
						{
							"cred_def_id": addressCredDefID,
						},
						{
							"cred_def_id": vaccineCredDefID,
						},
					},
				}

				street2Proof := map[string]interface{}{
					"name": "Address Line 2",
					"restrictions": []map[string]interface{}{
						{
							"cred_def_id": cornerstoneCredDefID,
						},
						{
							"cred_def_id": addressCredDefID,
						},
						{
							"cred_def_id": vaccineCredDefID,
						},
					},
				}

				cityProof := map[string]interface{}{
					"name": "City",
					"restrictions": []map[string]interface{}{
						{
							"cred_def_id": cornerstoneCredDefID,
						},
						{
							"cred_def_id": addressCredDefID,
						},
						{
							"cred_def_id": vaccineCredDefID,
						},
					},
				}

				postalCodeProof := map[string]interface{}{
					"name": "Postal Code",
					"restrictions": []map[string]interface{}{
						{
							"cred_def_id": cornerstoneCredDefID,
						},
						{
							"cred_def_id": addressCredDefID,
						},
						{
							"cred_def_id": vaccineCredDefID,
						},
					},
				}

				vaccineTypeProof := map[string]interface{}{
					"name": "Vaccine Type",
					"restrictions": []map[string]interface{}{
						{
							"cred_def_id": cornerstoneCredDefID,
						},
						{
							"cred_def_id": addressCredDefID,
						},
						{
							"cred_def_id": vaccineCredDefID,
						},
					},
				}

				vaccineDoseProof := map[string]interface{}{
					"name": "Vaccine Dose",
					"restrictions": []map[string]interface{}{
						{
							"cred_def_id": cornerstoneCredDefID,
						},
						{
							"cred_def_id": addressCredDefID,
						},
						{
							"cred_def_id": vaccineCredDefID,
						},
					},
				}

				dateOfVaccineProof := map[string]interface{}{
					"name": "Date of Vaccination",
					"restrictions": []map[string]interface{}{
						{
							"cred_def_id": cornerstoneCredDefID,
						},
						{
							"cred_def_id": addressCredDefID,
						},
						{
							"cred_def_id": vaccineCredDefID,
						},
					},
				}

				proofRequest := models.IAMZAProofRequest{
					Comment:      "IAMZA Proof Request",
					ConnectionID: request.ConnectionID,
					PresentationRequest: models.IAMZAPresentationRequest{
						Name:    "Proof of Identity, Physical Address & Vaccination",
						Version: "1.0",
						RequestedAttributes: models.IAMZARequestedAttributes{
							idProof,
							namesProof,
							surnameProof,
							genderProof,
							dobProof,
							street1Proof,
							street2Proof,
							cityProof,
							postalCodeProof,
							vaccineTypeProof,
							vaccineDoseProof,
							dateOfVaccineProof,
						},
						RequestedPredicates: models.RequestedPredicates{},
					},
				}

				_, err = client.SendIAMZAProofRequest(proofRequest)
				if err != nil {
					log.Error.Printf("Failed to send proof request: %s", err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				// For email proof notification
				proofEmail, _ := cache.ReadString(request.InvitationKey + "email")
				err = cache.UpdateString(request.ConnectionID, proofEmail)
				if err != nil {
					log.Error.Printf("Failed to cache proof data: %s", err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				cache.DeleteString(request.InvitationKey + "IAMZA")
				cache.DeleteString(request.InvitationKey + "email")

				log.Info.Println("Proof request sent")
				w.WriteHeader(http.StatusOK)
			} else if request.State == "active" {
				proofData, err := cache.ReadStruct(request.InvitationKey)
				if err != nil {
					log.Error.Printf("Failed to read proof data: %s", err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				var idProof interface{}
				var namesProof interface{}
				var surnameProof interface{}
				var genderProof interface{}
				var dobProof interface{}
				var cobProof interface{}

				credDefID := config.GetCornerstoneCredDefID()

				if proofData.IDNumber {
					idProof = map[string]interface{}{
						"name": "ID Number",
						"restrictions": []map[string]interface{}{
							{
								"cred_def_id": credDefID,
							},
						},
					}
				} else {
					idProof = nil
				}

				if proofData.FirstNames {
					namesProof = map[string]interface{}{
						"name": "First Names",
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

				if proofData.DOB {
					dobProof = map[string]interface{}{
						"name": "Date of Birth",
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
						"name": "Country of Birth",
						"restrictions": []map[string]interface{}{
							{
								"cred_def_id": credDefID,
							},
						},
					}
				} else {
					cobProof = nil
				}

				proofRequest := models.ProofRequest{
					Comment:      "Cornerstone Proof Request",
					ConnectionID: request.ConnectionID,
					PresentationRequest: models.PresentationRequest{
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

				_, err = client.SendProofRequest(proofRequest)
				if err != nil {
					log.Error.Printf("Failed to send proof request: %s", err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				// For email proof notification
				err = cache.UpdateString(request.ConnectionID, proofData.Email)
				if err != nil {
					log.Error.Printf("Failed to cache cornerstone proof data: %s", err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				cache.DeleteStruct(request.InvitationKey)

				log.Info.Println("Proof request sent")
				w.WriteHeader(http.StatusOK)
			}

		case "present_proof":
			var request models.PresentProofWebhookResponse
			err := json.NewDecoder(r.Body).Decode(&request)
			if err != nil {
				log.Error.Printf("Failed to decode request body: %s", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			proofEmail, err := cache.ReadString(request.ConnectionID)
			if err != nil {
				log.Error.Printf("Failed to read cached user data: %s", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if request.State == "verified" && proofEmail != "" {
				presExRecord, err := client.GetPresExRecord(request.PresentationExchangeID)
				if err != nil {
					log.Error.Printf("Failed to get presentation record: %s", err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				status := ""
				if presExRecord.Verified == "true" {
					status = "successfully"
				} else if presExRecord.Verified == "false" {
					status = "unsuccessfully. Please restart the process and provide valid attributes."
				}

				err = utils.SendNotificationEmail(proofEmail, status, config)
				if err != nil {
					log.Error.Printf("Failed to send credential notification email: %s", err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				cache.DeleteString(request.ConnectionID)

				log.Info.Println("Notified user successfully about credential verification!")
				w.WriteHeader(http.StatusOK)
			}

		case "issue_credential":
		case "basicmessages":
		case "revocation_registry":
		case "problem_report":
		case "issuer_cred_rev":

		default:
			log.Warning.Printf("Unexpected topic: %s", topic)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
