package models

type CornerstoneCredentialProofRequest struct {
	IDNumber       bool   `json:"id_number"`
	FirstNames     bool   `json:"first_names"`
	Surname        bool   `json:"surname"`
	Gender         bool   `json:"gender"`
	DOB            bool   `json:"date_of_birth"`
	CountryOfBirth bool   `json:"country_of_birth"`
	Email          string `json:"email"`
}

type CredentialProofRequest struct {
	Email string `json:"email"`
}
