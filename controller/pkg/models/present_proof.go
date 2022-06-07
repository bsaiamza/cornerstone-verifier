package models

type CreateProofPresentationRequest struct {
	Comment             string              `json:"comment"`
	ConnectionID        string              `json:"connection_id"`
	PresentationRequest PresentationRequest `json:"presentation_request"`
}

type PresentationRequest struct {
	Indy Indy `json:"indy"`
}

type RequestedAttributes struct {
	The0_IDNumberUUID       The0___UUID `json:"0_IDNumber_uuid"`
	The0_SurnameUUID        The0___UUID `json:"0_Surname_uuid"`
	The0_ForenamesUUID      The0___UUID `json:"0_Forenames_uuid"`
	The0_GenderUUID         The0___UUID `json:"0_Gender_uuid"`
	The0_DateOfBirthUUID    The0___UUID `json:"0_DateOfBirth_uuid"`
	The0_CountryOfBirthUUID The0___UUID `json:"0_CountryOfBirth_uuid"`
}

type The0___UUID struct {
	Name         string        `json:"name"`
	Restrictions []Restriction `json:"restrictions"`
}

type Restriction struct {
	CredDefID string `json:"cred_def_id"`
}

type RequestedPredicates struct {
}

type CreateProofPresentationResponse struct {
	PresRequest  CreateProofPresentationResponsePresRequest `json:"pres_request"`
	Trace        bool                                       `json:"trace"`
	ConnectionID string                                     `json:"connection_id"`
	State        string                                     `json:"state"`
	AutoPresent  bool                                       `json:"auto_present"`
	Initiator    string                                     `json:"initiator"`
	ByFormat     ByFormat                                   `json:"by_format"`
	Role         string                                     `json:"role"`
	UpdatedAt    string                                     `json:"updated_at"`
	CreatedAt    string                                     `json:"created_at"`
	PresExID     string                                     `json:"pres_ex_id"`
	ThreadID     string                                     `json:"thread_id"`
}

// type ByFormat struct {
// 	PresRequest ByFormatPresRequest `json:"pres_request"`
// }

type ByFormatPresRequest struct {
	Indy Indy `json:"indy"`
}

type Indy struct {
	Name                string                `json:"name"`
	Version             string                `json:"version"`
	RequestedAttributes []RequestedAttributes `json:"requested_attributes"`
	RequestedPredicates []string              `json:"requested_predicates"`
	Nonce               string                `json:"nonce,omitempty"`
}

type CreateProofPresentationResponsePresRequest struct {
	Type                       string                       `json:"@type"`
	ID                         string                       `json:"@id"`
	RequestPresentationsAttach []RequestPresentationsAttach `json:"request_presentations~attach"`
	WillConfirm                bool                         `json:"will_confirm"`
	Comment                    string                       `json:"comment"`
	Formats                    []Format                     `json:"formats"`
}

type Format struct {
	AttachID string `json:"attach_id"`
	Format   string `json:"format"`
}

type RequestPresentationsAttach struct {
	ID       string `json:"@id"`
	MIMEType string `json:"mime-type"`
	Data     Data   `json:"data"`
}

type Data struct {
	Base64 string `json:"base64"`
}

type ListProofRecordsParams struct {
	ConnectionID string `json:"connection_id"`
	Role         string `json:"role"`
	State        string `json:"state"`
	ThreadID     string `json:"thread_id"`
}

type ListProofRecordsResponse struct {
	Results []Result `json:"results"`
}

type Result struct {
	PresRequest  ResultPresRequest `json:"pres_request"`
	Trace        bool              `json:"trace"`
	ConnectionID string            `json:"connection_id"`
	State        string            `json:"state"`
	AutoPresent  bool              `json:"auto_present"`
	Initiator    string            `json:"initiator"`
	ByFormat     ByFormat          `json:"by_format"`
	ErrorMsg     *string           `json:"error_msg,omitempty"`
	Role         string            `json:"role"`
	UpdatedAt    string            `json:"updated_at"`
	CreatedAt    string            `json:"created_at"`
	PresExID     string            `json:"pres_ex_id"`
	ThreadID     string            `json:"thread_id"`
	Pres         *ResultPres       `json:"pres,omitempty"`
	Verified     *string           `json:"verified,omitempty"`
}

type ByFormat struct {
	PresRequest ByFormatPresRequestRes `json:"pres_request"`
	Pres        *ByFormatPres          `json:"pres,omitempty"`
}

type ByFormatPres struct {
	Indy PresIndy `json:"indy"`
}

type PresIndy struct {
	Proof          IndyProof      `json:"proof"`
	RequestedProof RequestedProof `json:"requested_proof"`
	Identifiers    []Identifier   `json:"identifiers"`
}

type Identifier struct {
	SchemaID  string      `json:"schema_id"`
	CredDefID CredDefID   `json:"cred_def_id"`
	RevRegID  interface{} `json:"rev_reg_id"`
	Timestamp interface{} `json:"timestamp"`
}

type IndyProof struct {
	Proofs          []ProofElement  `json:"proofs"`
	AggregatedProof AggregatedProof `json:"aggregated_proof"`
}

type AggregatedProof struct {
	CHash string    `json:"c_hash"`
	CList [][]int64 `json:"c_list"`
}

type ProofElement struct {
	PrimaryProof  PrimaryProof `json:"primary_proof"`
	NonRevocProof interface{}  `json:"non_revoc_proof"`
}

type PrimaryProof struct {
	EqProof  EqProof       `json:"eq_proof"`
	GeProofs []interface{} `json:"ge_proofs"`
}

type EqProof struct {
	RevealedAttrs M      `json:"revealed_attrs"`
	APrime        string `json:"a_prime"`
	E             string `json:"e"`
	V             string `json:"v"`
	M             M      `json:"m"`
	M2            string `json:"m2"`
}

type M struct {
	MasterSecret   *string `json:"master_secret,omitempty"`
	Gender         *string `json:"gender,omitempty"`
	Forenames      *string `json:"forenames,omitempty"`
	Countryofbirth *string `json:"countryofbirth,omitempty"`
	Idnumber       *string `json:"idnumber,omitempty"`
	Surname        *string `json:"surname,omitempty"`
	Dateofbirth    *string `json:"dateofbirth,omitempty"`
}

type RequestedProof struct {
	RevealedAttrs     RevealedAttrs `json:"revealed_attrs"`
	SelfAttestedAttrs PleaseACK     `json:"self_attested_attrs"`
	UnrevealedAttrs   PleaseACK     `json:"unrevealed_attrs"`
	Predicates        PleaseACK     `json:"predicates"`
}

type PleaseACK struct {
}

type RevealedAttrs struct {
	The0_DateOfBirthUUID    RevealedAttrs0_CountryOfBirthUUID `json:"0_DateOfBirth_uuid"`
	The0_SurnameUUID        RevealedAttrs0_CountryOfBirthUUID `json:"0_Surname_uuid"`
	The0_ForenamesUUID      RevealedAttrs0_CountryOfBirthUUID `json:"0_Forenames_uuid"`
	The0_GenderUUID         RevealedAttrs0_CountryOfBirthUUID `json:"0_Gender_uuid"`
	The0_IDNumberUUID       RevealedAttrs0_CountryOfBirthUUID `json:"0_IDNumber_uuid"`
	The0_CountryOfBirthUUID RevealedAttrs0_CountryOfBirthUUID `json:"0_CountryOfBirth_uuid"`
}

type RevealedAttrs0_CountryOfBirthUUID struct {
	SubProofIndex int64  `json:"sub_proof_index"`
	Raw           string `json:"raw"`
	Encoded       string `json:"encoded"`
}

type ByFormatPresRequestRes struct {
	Indy PresRequestIndy `json:"indy"`
}

type PresRequestIndy struct {
	Name                string                 `json:"name"`
	Version             string                 `json:"version"`
	RequestedAttributes RequestedAttributesRes `json:"requested_attributes"`
	RequestedPredicates PleaseACK              `json:"requested_predicates"`
	Nonce               string                 `json:"nonce"`
}

type RequestedAttributesRes struct {
	The0_IDNumberUUID       RequestedAttributes0_CountryOfBirthUUID `json:"0_IDNumber_uuid"`
	The0_SurnameUUID        RequestedAttributes0_CountryOfBirthUUID `json:"0_Surname_uuid"`
	The0_ForenamesUUID      RequestedAttributes0_CountryOfBirthUUID `json:"0_Forenames_uuid"`
	The0_GenderUUID         RequestedAttributes0_CountryOfBirthUUID `json:"0_Gender_uuid"`
	The0_DateOfBirthUUID    RequestedAttributes0_CountryOfBirthUUID `json:"0_DateOfBirth_uuid"`
	The0_CountryOfBirthUUID RequestedAttributes0_CountryOfBirthUUID `json:"0_CountryOfBirth_uuid"`
}

type RequestedAttributes0_CountryOfBirthUUID struct {
	Name         Name          `json:"name"`
	Restrictions []Restriction `json:"restrictions"`
}

type ResultPres struct {
	Type                string                `json:"@type"`
	ID                  string                `json:"@id"`
	PleaseACK           PleaseACK             `json:"~please_ack"`
	Thread              Thread                `json:"~thread"`
	Comment             string                `json:"comment"`
	Formats             []FormatElement       `json:"formats"`
	PresentationsAttach []PresentationsAttach `json:"presentations~attach"`
}

type FormatElement struct {
	AttachID string     `json:"attach_id"`
	Format   FormatEnum `json:"format"`
}

type PresentationsAttach struct {
	ID       string   `json:"@id"`
	MIMEType MIMEType `json:"mime-type"`
	Data     Data     `json:"data"`
}

type Thread struct {
	Thid           string           `json:"thid"`
	ReceivedOrders map[string]int64 `json:"received_orders"`
}

type ResultPresRequest struct {
	Type                       string                `json:"@type"`
	ID                         string                `json:"@id"`
	RequestPresentationsAttach []PresentationsAttach `json:"request_presentations~attach"`
	WillConfirm                bool                  `json:"will_confirm"`
	Comment                    string                `json:"comment"`
	Formats                    []FormatElement       `json:"formats"`
}

type CredDefID string

const (
	BER7WwiAMK9IgkiRjPYpEp3CL40479Cornerstone12 CredDefID = "BER7WwiAMK9igkiRjPYpEp:3:CL:40479:cornerstone_1.2"
	DSg3BALcDFyVnpJgvR1AMi3CL23Cornerstone      CredDefID = "DSg3bALcDFyVnpJgvR1aMi:3:CL:23:cornerstone"
)

type Name string

const (
	CountryOfBirth Name = "CountryOfBirth"
	DateOfBirth    Name = "DateOfBirth"
	Forenames      Name = "Forenames"
	Gender         Name = "Gender"
	IDNumber       Name = "IDNumber"
	Surname        Name = "Surname"
)

type FormatEnum string

const (
	HlindyProofReqV20 FormatEnum = "hlindy/proof-req@v2.0"
	HlindyProofV20    FormatEnum = "hlindy/proof@v2.0"
)

type MIMEType string

const (
	ApplicationJSON MIMEType = "application/json"
)
