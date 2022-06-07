package models

type PrepareProofPresentationData struct {
	IDNumber       bool `json:"id_number"`
	Surname        bool `json:"surname"`
	Forenames      bool `json:"forenames"`
	Gender         bool `json:"gender"`
	DateOfBirth    bool `json:"date_of_birth"`
	CountryOfBirth bool `json:"country_of_birth"`
}
