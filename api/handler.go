package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"iamza-verifier/pkg/acapy"
	"iamza-verifier/pkg/config"
	"iamza-verifier/pkg/log"
	"iamza-verifier/pkg/models"
	"iamza-verifier/pkg/server"
	"iamza-verifier/pkg/utils"

	"github.com/gorilla/mux"
	"github.com/skip2/go-qrcode"
)

func verifyCornerstoneCredential(config *config.Config, acapy *acapy.Client, cache *utils.BigCache) http.HandlerFunc {
	mdw := []server.Middleware{
		server.LogAPIRequest,
	}

	return server.ChainMiddleware(verifyCornerstoneCredentialHandler(config, acapy, cache), mdw...)
}

func verifyCornerstoneCredentialHandler(config *config.Config, acapy *acapy.Client, cache *utils.BigCache) http.HandlerFunc {
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

		invitation, err := acapy.CreateInvitation(invitationRequest)
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

		// Step 2: Cache data
		err = cache.String(invitation.Invitation.RecipientKeys[0]+"Cornerstone", "Cornerstone proof")
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

func verifyCornerstoneCredentialByEmail(config *config.Config, acapy *acapy.Client, cache *utils.BigCache) http.HandlerFunc {
	mdw := []server.Middleware{
		server.LogAPIRequest,
	}

	return server.ChainMiddleware(verifyCornerstoneCredentialByEmailHandler(config, acapy, cache), mdw...)
}

func verifyCornerstoneCredentialByEmailHandler(config *config.Config, acapy *acapy.Client, cache *utils.BigCache) http.HandlerFunc {
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
		var email models.EmailProofRequest
		err := json.NewDecoder(r.Body).Decode(&email)
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
		err = utils.ValidEmail(email.Email)
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

		invitation, err := acapy.CreateInvitation(invitationRequest)
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

		// Step 4: Cache data
		err = cache.String(invitation.Invitation.RecipientKeys[0]+"Cornerstone", "Cornerstone proof")
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

		err = cache.String(invitation.Invitation.RecipientKeys[0]+"CornerstoneEmail", email.Email)
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
		err = utils.SendProofRequestByEmail(email.Email, invitation.Invitation.RecipientKeys[0], qrCodePng, config)
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

func verifyContactableCredential(config *config.Config, acapy *acapy.Client, cache *utils.BigCache) http.HandlerFunc {
	mdw := []server.Middleware{
		server.LogAPIRequest,
	}

	return server.ChainMiddleware(verifyContactableCredentialHandler(config, acapy, cache), mdw...)
}

func verifyContactableCredentialHandler(config *config.Config, acapy *acapy.Client, cache *utils.BigCache) http.HandlerFunc {
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

		invitation, err := acapy.CreateInvitation(invitationRequest)
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

		// Step 2: Cache data
		err = cache.String(invitation.Invitation.RecipientKeys[0]+"Contactable", "Contactable proof")
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

func verifyContactableCredentialByEmail(config *config.Config, acapy *acapy.Client, cache *utils.BigCache) http.HandlerFunc {
	mdw := []server.Middleware{
		server.LogAPIRequest,
	}

	return server.ChainMiddleware(verifyContactableCredentialByEmailHandler(config, acapy, cache), mdw...)
}

func verifyContactableCredentialByEmailHandler(config *config.Config, acapy *acapy.Client, cache *utils.BigCache) http.HandlerFunc {
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
		var email models.EmailProofRequest
		err := json.NewDecoder(r.Body).Decode(&email)
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
		err = utils.ValidEmail(email.Email)
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

		invitation, err := acapy.CreateInvitation(invitationRequest)
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

		// Step 4: Cache data
		err = cache.String(invitation.Invitation.RecipientKeys[0]+"Contactable", "Contactable proof")
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

		err = cache.String(invitation.Invitation.RecipientKeys[0]+"ContactableEmail", email.Email)
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
		err = utils.SendProofRequestByEmail(email.Email, invitation.Invitation.RecipientKeys[0], qrCodePng, config)
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

func verifyAddressCredential(config *config.Config, acapy *acapy.Client, cache *utils.BigCache) http.HandlerFunc {
	mdw := []server.Middleware{
		server.LogAPIRequest,
	}

	return server.ChainMiddleware(verifyAddressCredentialHandler(config, acapy, cache), mdw...)
}

func verifyAddressCredentialHandler(config *config.Config, acapy *acapy.Client, cache *utils.BigCache) http.HandlerFunc {
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

		invitation, err := acapy.CreateInvitation(invitationRequest)
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

		// Step 2: Cache data
		err = cache.String(invitation.Invitation.RecipientKeys[0]+"Address", "Address proof")
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

func verifyAddressCredentialByEmail(config *config.Config, acapy *acapy.Client, cache *utils.BigCache) http.HandlerFunc {
	mdw := []server.Middleware{
		server.LogAPIRequest,
	}

	return server.ChainMiddleware(verifyAddressCredentialByEmailHandler(config, acapy, cache), mdw...)
}

func verifyAddressCredentialByEmailHandler(config *config.Config, acapy *acapy.Client, cache *utils.BigCache) http.HandlerFunc {
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
		var email models.EmailProofRequest
		err := json.NewDecoder(r.Body).Decode(&email)
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
		err = utils.ValidEmail(email.Email)
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

		invitation, err := acapy.CreateInvitation(invitationRequest)
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

		// Step 4: Cache data
		err = cache.String(invitation.Invitation.RecipientKeys[0]+"Address", "Address proof")
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

		err = cache.String(invitation.Invitation.RecipientKeys[0]+"AddressEmail", email.Email)
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
		err = utils.SendProofRequestByEmail(email.Email, invitation.Invitation.RecipientKeys[0], qrCodePng, config)
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

func verifyVaccineCredential(config *config.Config, acapy *acapy.Client, cache *utils.BigCache) http.HandlerFunc {
	mdw := []server.Middleware{
		server.LogAPIRequest,
	}

	return server.ChainMiddleware(verifyVaccineCredentialHandler(config, acapy, cache), mdw...)
}

func verifyVaccineCredentialHandler(config *config.Config, acapy *acapy.Client, cache *utils.BigCache) http.HandlerFunc {
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

		invitation, err := acapy.CreateInvitation(invitationRequest)
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

		// Step 2: Cache data
		err = cache.String(invitation.Invitation.RecipientKeys[0]+"Vaccine", "Vaccine proof")
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

func verifyVaccineCredentialByEmail(config *config.Config, acapy *acapy.Client, cache *utils.BigCache) http.HandlerFunc {
	mdw := []server.Middleware{
		server.LogAPIRequest,
	}

	return server.ChainMiddleware(verifyVaccineCredentialByEmailHandler(config, acapy, cache), mdw...)
}

func verifyVaccineCredentialByEmailHandler(config *config.Config, acapy *acapy.Client, cache *utils.BigCache) http.HandlerFunc {
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
		var email models.EmailProofRequest
		err := json.NewDecoder(r.Body).Decode(&email)
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
		err = utils.ValidEmail(email.Email)
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

		invitation, err := acapy.CreateInvitation(invitationRequest)
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

		// Step 4: Cache data
		err = cache.String(invitation.Invitation.RecipientKeys[0]+"Vaccine", "Vaccine proof")
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

		err = cache.String(invitation.Invitation.RecipientKeys[0]+"VaccineEmail", email.Email)
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
		err = utils.SendProofRequestByEmail(email.Email, invitation.Invitation.RecipientKeys[0], qrCodePng, config)
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

func webhookEvents(config *config.Config, acapy *acapy.Client, cache *utils.BigCache) http.HandlerFunc {
	mdw := []server.Middleware{
		server.LogAPIRequest,
	}

	return server.ChainMiddleware(webhookEventsHandler(config, acapy, cache), mdw...)
}

func webhookEventsHandler(config *config.Config, acapy *acapy.Client, cache *utils.BigCache) http.HandlerFunc {
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

				_, err := acapy.PingConnection(request.ConnectionID, pingRequest)
				if err != nil {
					log.Error.Printf("Failed to ping holder: %s", err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}

			cornerstoneProofRequest, _ := cache.ReadString(request.InvitationKey + "Cornerstone")
			contactableProofRequest, _ := cache.ReadString(request.InvitationKey + "Contactable")
			addressProofRequest, _ := cache.ReadString(request.InvitationKey + "Address")
			vaccineProofRequest, _ := cache.ReadString(request.InvitationKey + "Vaccine")

			if cornerstoneProofRequest == "Cornerstone proof" && request.State == "active" {
				cornerstoneSchemaID := config.GetCornerstoneSchemaID()

				t := time.Now()
				pvDay := strconv.Itoa(t.Day())
				pvDayLen := len([]rune(pvDay))
				if pvDayLen == 1 {
					pvDay = "0" + pvDay
				}
				pvMonth := strconv.Itoa(int(t.Month()))
				pvMonthLen := len([]rune(pvMonth))
				if pvMonthLen == 1 {
					pvMonth = "0" + pvMonth
				}
				pvYear := config.GetPValueYear()

				pValueStr := fmt.Sprintf("%s%s%s", pvYear, pvMonth, pvDay)

				pValue, err := strconv.Atoi(pValueStr)
				if err != nil {
					log.Error.Printf("Failed to convert p_value: %s", err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				namesProof := map[string]interface{}{
					"name": "names",
					"restrictions": []map[string]interface{}{
						{
							"schema_id": cornerstoneSchemaID,
						},
					},
				}

				surnameProof := map[string]interface{}{
					"name": "surname",
					"restrictions": []map[string]interface{}{
						{
							"schema_id": cornerstoneSchemaID,
						},
					},
				}

				ageProof := map[string]interface{}{
					"name":    "date_of_birth",
					"p_type":  "<=",
					"p_value": pValue,
					"restrictions": []map[string]interface{}{
						{
							"schema_id": cornerstoneSchemaID,
						},
					},
				}

				proofRequest := models.CornerstoneProofRequest{
					Comment:      "Cornerstone Proof Request",
					ConnectionID: request.ConnectionID,
					PresentationRequest: models.CornerstonePresentationRequest{
						Name:    "Proof of Identity & Age",
						Version: "1.0",
						RequestedAttributes: models.CornerstoneRequestedAttributes{
							namesProof,
							surnameProof,
						},
						RequestedPredicates: models.CornerstoneRequestedPredicates{
							ageProof,
						},
					},
				}

				_, err = acapy.SendCornerstoneProofRequest(proofRequest)
				if err != nil {
					log.Error.Printf("Failed to send proof request: %s", err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				// For email proof notification
				proofEmail, _ := cache.ReadString(request.InvitationKey + "CornerstoneEmail")
				err = cache.String(request.ConnectionID, proofEmail)
				if err != nil {
					log.Error.Printf("Failed to cache proof data: %s", err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				cache.DeleteString(request.InvitationKey + "Cornerstone")
				cache.DeleteString(request.InvitationKey + "CornerstoneEmail")

				log.Info.Println("Proof request sent")
			} else if contactableProofRequest == "Contactable proof" && request.State == "active" {
				contactableCredDefID := config.GetContactableCredDefID()

				nameProof := map[string]interface{}{
					"name": "name",
					"restrictions": []map[string]interface{}{
						{
							"cred_def_id": contactableCredDefID,
						},
					},
				}

				surnameProof := map[string]interface{}{
					"name": "surname",
					"restrictions": []map[string]interface{}{
						{
							"cred_def_id": contactableCredDefID,
						},
					},
				}

				idnumberProof := map[string]interface{}{
					"name": "idnumber",
					"restrictions": []map[string]interface{}{
						{
							"cred_def_id": contactableCredDefID,
						},
					},
				}

				profilePictureProof := map[string]interface{}{
					"name": "profilePicture",
					"restrictions": []map[string]interface{}{
						{
							"cred_def_id": contactableCredDefID,
						},
					},
				}

				addressLine1Proof := map[string]interface{}{
					"name": "addressLine1",
					"restrictions": []map[string]interface{}{
						{
							"cred_def_id": contactableCredDefID,
						},
					},
				}

				suburbProof := map[string]interface{}{
					"name": "suburb",
					"restrictions": []map[string]interface{}{
						{
							"cred_def_id": contactableCredDefID,
						},
					},
				}

				cityProof := map[string]interface{}{
					"name": "city",
					"restrictions": []map[string]interface{}{
						{
							"cred_def_id": contactableCredDefID,
						},
					},
				}

				provinceProof := map[string]interface{}{
					"name": "province",
					"restrictions": []map[string]interface{}{
						{
							"cred_def_id": contactableCredDefID,
						},
					},
				}

				countryProof := map[string]interface{}{
					"name": "country",
					"restrictions": []map[string]interface{}{
						{
							"cred_def_id": contactableCredDefID,
						},
					},
				}

				postalCodeProof := map[string]interface{}{
					"name": "postalCode",
					"restrictions": []map[string]interface{}{
						{
							"cred_def_id": contactableCredDefID,
						},
					},
				}

				identityDocumentProof := map[string]interface{}{
					"name": "identityDocument",
					"restrictions": []map[string]interface{}{
						{
							"cred_def_id": contactableCredDefID,
						},
					},
				}

				proofRequest := models.ContactableProofRequest{
					Comment:      "Contactable Proof Request",
					ConnectionID: request.ConnectionID,
					PresentationRequest: models.ContactablePresentationRequest{
						Name:    "Contactable Proof of Identity",
						Version: "1.0",
						RequestedAttributes: models.ContactableRequestedAttributes{
							nameProof,
							surnameProof,
							idnumberProof,
							profilePictureProof,
							addressLine1Proof,
							suburbProof,
							cityProof,
							provinceProof,
							countryProof,
							postalCodeProof,
							identityDocumentProof,
						},
						RequestedPredicates: models.ContactableRequestedPredicates{},
					},
				}

				_, err = acapy.SendContactableProofRequest(proofRequest)
				if err != nil {
					log.Error.Printf("Failed to send proof request: %s", err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				// For email proof notification
				proofEmail, _ := cache.ReadString(request.InvitationKey + "ContactableEmail")
				err = cache.String(request.ConnectionID, proofEmail)
				if err != nil {
					log.Error.Printf("Failed to cache proof data: %s", err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				cache.DeleteString(request.InvitationKey + "Contactable")
				cache.DeleteString(request.InvitationKey + "ContactableEmail")

				log.Info.Println("Proof request sent")
			} else if addressProofRequest == "Address proof" && request.State == "active" {
				addressSchemaID := config.GetAddressSchemaID()

				pValueStr := time.Now().Format("20060102")

				pValue, err := strconv.Atoi(pValueStr)
				if err != nil {
					log.Error.Printf("Failed to convert p_value: %s", err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				statementDateProof := map[string]interface{}{
					"name": "statement_date",
					"restrictions": []map[string]interface{}{
						{
							"schema_id": addressSchemaID,
						},
					},
				}

				selfAttestedProof := map[string]interface{}{
					"name": "self_attested",
					"restrictions": []map[string]interface{}{
						{
							"schema_id": addressSchemaID,
						},
					},
				}

				expirationProof := map[string]interface{}{
					"name":    "expiry_date",
					"p_type":  ">",
					"p_value": pValue,
					"restrictions": []map[string]interface{}{
						{
							"schema_id": addressSchemaID,
						},
					},
				}

				proofRequest := models.AddressProofRequest{
					Comment:      "Physical Address Proof Request",
					ConnectionID: request.ConnectionID,
					PresentationRequest: models.AddressPresentationRequest{
						Name:    "Physical Address Proof",
						Version: "1.0",
						RequestedAttributes: models.AddressRequestedAttributes{
							statementDateProof,
							selfAttestedProof,
						},
						RequestedPredicates: models.AddressRequestedPredicates{
							expirationProof,
						},
					},
				}

				_, err = acapy.SendAddressProofRequest(proofRequest)
				if err != nil {
					log.Error.Printf("Failed to send proof request: %s", err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				// For email proof notification
				proofEmail, _ := cache.ReadString(request.InvitationKey + "AddressEmail")
				err = cache.String(request.ConnectionID, proofEmail)
				if err != nil {
					log.Error.Printf("Failed to cache proof data: %s", err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				cache.DeleteString(request.InvitationKey + "Address")
				cache.DeleteString(request.InvitationKey + "AddressEmail")

				log.Info.Println("Proof request sent")
			} else if vaccineProofRequest == "Vaccine proof" && request.State == "active" {
				vaccineSchemaID := config.GetVaccineSchemaID()

				pValueStr := time.Now().Format("20060102")

				pValue, err := strconv.Atoi(pValueStr)
				if err != nil {
					log.Error.Printf("Failed to convert p_value: %s", err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				vaccineTypeProof := map[string]interface{}{
					"name": "vaccine_type",
					"restrictions": []map[string]interface{}{
						{
							"schema_id": vaccineSchemaID,
						},
					},
				}

				expirationProof := map[string]interface{}{
					"name":    "expiry_date",
					"p_type":  ">",
					"p_value": pValue,
					"restrictions": []map[string]interface{}{
						{
							"schema_id": vaccineSchemaID,
						},
					},
				}

				proofRequest := models.VaccineProofRequest{
					Comment:      "Vaccination Proof Request",
					ConnectionID: request.ConnectionID,
					PresentationRequest: models.VaccinePresentationRequest{
						Name:    "Vaccination Proof",
						Version: "1.0",
						RequestedAttributes: models.VaccineRequestedAttributes{
							vaccineTypeProof,
						},
						RequestedPredicates: models.VaccineRequestedPredicates{
							expirationProof,
						},
					},
				}

				_, err = acapy.SendVaccineProofRequest(proofRequest)
				if err != nil {
					log.Error.Printf("Failed to send proof request: %s", err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				// For email proof notification
				proofEmail, _ := cache.ReadString(request.InvitationKey + "VaccineEmail")
				err = cache.String(request.ConnectionID, proofEmail)
				if err != nil {
					log.Error.Printf("Failed to cache proof data: %s", err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				cache.DeleteString(request.InvitationKey + "Vaccine")
				cache.DeleteString(request.InvitationKey + "VaccineEmail")

				log.Info.Println("Proof request sent")
			}

		case "present_proof":
			var request models.PresentProofWebhookResponse
			err := json.NewDecoder(r.Body).Decode(&request)
			if err != nil {
				log.Error.Printf("Failed to decode request body: %s", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			presExRecord, err := acapy.GetPresExRecord(request.PresentationExchangeID)
			if err != nil {
				log.Error.Printf("Failed to get presentation record: %s", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if request.State == "verified" && presExRecord.Verified == "true" {
				log.Info.Println("Credential Verification: Successful!")
			} else if request.State == "verified" && presExRecord.Verified == "false" {
				log.Info.Println("Credential Verification: Unsuccessful!")
			}

			proofEmail, err := cache.ReadString(request.ConnectionID)
			if err != nil {
				log.Error.Printf("Failed to read cached user data: %s", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if request.State == "verified" && proofEmail != "" {
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
			}

			if request.State == "verified" {
				txnCounterSwitch := config.GetTxnCounterSwitch()

				if txnCounterSwitch == "0" {
				} else if txnCounterSwitch == "1" {
					log.Info.Println("Calling Transaction Counter")

					txnID := utils.RandomTxnID(12)

					p1 := "%7B%22WriterDID%22%3A%20%22BER7WwiAMK9igkiRjPYpEp%22%2C%22WriterDomain%22%3A%20%22IAMZA%20Verifier%22%2C%22WriterMetaData%22%3A%20%7B%22txnid%22%3A%20%22"
					p2 := "%22%7D%7D"

					payload := p1 + txnID + p2

					url := config.GetTxnCounterAPI() + payload

					req, err := http.NewRequest("POST", url, nil)
					if err != nil {
						log.Error.Printf("Failed to create new Transaction Counter API request: %s", err)
						return
					}

					req.Header.Add("Content-Type", "application/json")

					res, err := http.DefaultClient.Do(req)
					if err != nil {
						log.Error.Printf("Failed on Transaction Counter API call: %s", err)
						return
					}

					defer res.Body.Close()
					body, _ := ioutil.ReadAll(res.Body)

					fmt.Println("\n")
					fmt.Println("Transaction Counter API response: ", string(body))
				}
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
