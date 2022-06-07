package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	acapy "cornerstone_verifier/pkg/acapy_client"
	"cornerstone_verifier/pkg/config"
	"cornerstone_verifier/pkg/log"
	"cornerstone_verifier/pkg/models"
	"cornerstone_verifier/pkg/server"
	"cornerstone_verifier/pkg/util"
)

func prepareProofData(config *config.Config, acapyClient *acapy.Client, cache *util.BigCache) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(prepareProofDataHandler(config, acapyClient, cache), mdw...)
}
func prepareProofDataHandler(config *config.Config, acapyClient *acapy.Client, cache *util.BigCache) http.HandlerFunc {
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

		log.Info.Println("Preparing proof request data...")

		var data models.PrepareProofPresentationData
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
		// qrCodePng, err := qrcode.Encode(invitation.InvitationURL, qrcode.Medium, 256)
		// if err != nil {
		// 	log.Warning.Print("Failed to create QR code: ", err)
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	res := server.Res{
		// 		"success": false,
		// 		"msg":     "Failed to create QR code: " + err.Error(),
		// 	}
		// 	json.NewEncoder(w).Encode(res)
		// 	return
		// }

		// Step 3: Cache user data
		prepareProofData := models.PrepareProofPresentationData{
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

		// Step 4: Send email

		// var prefix string
		// if data.Gender == "Female" {
		// 	prefix = "Ms/Mrs "
		// }
		// if data.Gender == "Male" {
		// 	prefix = "Mr "
		// }

		// err = util.SendEmail(prefix+data.Surname, data.Email, invitation.Invitation.RecipientKeys[0], qrCodePng)
		// if err != nil {
		// 	log.Warning.Print("Failed to send credential email: ", err)
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	res := server.Res{
		// 		"success": false,
		// 		"msg":     "Failed to send credential email: " + err.Error(),
		// 	}
		// 	json.NewEncoder(w).Encode(res)
		// 	return
		// }

		// err = os.Remove("./" + invitation.Invitation.RecipientKeys[0] + ".png")
		// if err != nil {
		// 	log.Warning.Print("Failed to remove QR code: ", err)
		// w.WriteHeader(http.StatusInternalServerError)
		// res := server.Res{
		// 	"success": false,
		// 	"msg":     "Failed to remove QR code: " + err.Error(),
		// }
		// json.NewEncoder(w).Encode(res)
		// }

		log.Info.Println("Proof data prepared!")

		// w.Write(qrCodePng)
		// w.Header().Set("Content-Type", "image/png")
		res := server.Res{
			"success":      true,
			"proofRequest": invitation.InvitationURL,
		}
		json.NewEncoder(w).Encode(res)
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

			credDefID := config.GetCredDefID()
			if credDefID == "" {
				credDefID = "BER7WwiAMK9igkiRjPYpEp:3:CL:40479:cornerstone_1.2"
			}

			var idProof models.The0___UUID
			var forenamesProof models.The0___UUID
			var surnameProof models.The0___UUID
			var genderProof models.The0___UUID
			var dobProof models.The0___UUID
			var cobProof models.The0___UUID

			if proofData.IDNumber {
				idProof = models.The0___UUID{
					Name: "IDNumber",
					Restrictions: []models.Restriction{
						{
							CredDefID: "BER7WwiAMK9igkiRjPYpEp:3:CL:40479:cornerstone_1.2",
						},
					},
				}
			}

			if proofData.Forenames {
				forenamesProof = models.The0___UUID{
					Name: "Forenames",
					Restrictions: []models.Restriction{
						{
							CredDefID: "BER7WwiAMK9igkiRjPYpEp:3:CL:40479:cornerstone_1.2",
						},
					},
				}
			}

			if proofData.Surname {
				surnameProof = models.The0___UUID{
					Name: "Surname",
					Restrictions: []models.Restriction{
						{
							CredDefID: "BER7WwiAMK9igkiRjPYpEp:3:CL:40479:cornerstone_1.2",
						},
					},
				}
			}

			if proofData.Gender {
				genderProof = models.The0___UUID{
					Name: "Gender",
					Restrictions: []models.Restriction{
						{
							CredDefID: "BER7WwiAMK9igkiRjPYpEp:3:CL:40479:cornerstone_1.2",
						},
					},
				}
			}

			if proofData.DateOfBirth {
				dobProof = models.The0___UUID{
					Name: "DateOfBirth",
					Restrictions: []models.Restriction{
						{
							CredDefID: "BER7WwiAMK9igkiRjPYpEp:3:CL:40479:cornerstone_1.2",
						},
					},
				}
			}

			if proofData.CountryOfBirth {
				cobProof = models.The0___UUID{
					Name: "CountryOfBirth",
					Restrictions: []models.Restriction{
						{
							CredDefID: "BER7WwiAMK9igkiRjPYpEp:3:CL:40479:cornerstone_1.2",
						},
					},
				}
			}

			sendProofRequest := models.CreateProofPresentationRequest{
				Comment:      "Please provide proof you are who you say you are",
				ConnectionID: request.ConnectionID,
				PresentationRequest: models.PresentationRequest{
					Indy: models.Indy{
						Name:    "Proof of Identity",
						Version: "1.0",
						RequestedAttributes: []models.RequestedAttributes{
							{
								The0_IDNumberUUID: idProof,
							},
							{
								The0_ForenamesUUID: forenamesProof,
							},
							{
								The0_SurnameUUID: surnameProof,
							},
							{
								The0_GenderUUID: genderProof,
							},
							{
								The0_DateOfBirthUUID: dobProof,
							},
							{
								The0_CountryOfBirthUUID: cobProof,
							},
						},
						RequestedPredicates: []string{},
					},
				},
			}

			_, err = acapyClient.CreateProofPresentation(sendProofRequest)
			if err != nil {
				log.Error.Printf("Failed to create proof presentation: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				res := server.Res{
					"success": false,
					"msg":     "Failed to create proof presentation: " + err.Error(),
				}
				json.NewEncoder(w).Encode(res)
				return
			}

			cache.DeleteDataCache(request.InvitationKey)

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
