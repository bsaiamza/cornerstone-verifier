package models

type AcapyRequest struct {
	Method         string
	URL            string
	QueryParams    map[string]string
	Body           interface{}
	ResponseObject interface{}
}

type AcapyPostRequest struct {
	Endpoint    string
	QueryParams map[string]string
	Body        interface{}
	Response    interface{}
}

type AcapyGetRequest struct {
	Endpoint    string
	QueryParams map[string]string
	Response    interface{}
}

type DHASuccessResponse struct {
	IDNumber                   string `json:"Identity_Number"`
	Forenames                  string `json:"Names"`
	Surname                    string `json:"Surname"`
	Gender                     string `json:"Sex"`
	IssueDate                  string `json:"Issue_Date"`
	DateOfBirth                string `json:"Date_of_Birth"`
	BiometricsPhoto            string `json:"Biometrics-photo"`
	BiometricsFingerprint      string `json:"Biometrics-fingerprint"`
	BiometricsFingerprintMatch int    `json:"Biometrics-fingerprint_match"`
	Nationality                string `json:"Nationality"`
	CountryOfBirth             string `json:"Country_of_Birth"`
}

type DHAErrorResponse []DHAError

type DHAError struct {
	OriginatingDate string `json:"originatingDate"`
	ResponseCode    int64  `json:"responseCode"`
	ResponseDesc    string `json:"responseDesc"`
	ZaID            string `json:"ZA ID"`
	EndToEndID      string `json:"endToEndId"`
}
