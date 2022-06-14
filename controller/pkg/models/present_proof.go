package models

type ProofRequest struct {
	Comment             string              `json:"comment"`
	ConnectionID        string              `json:"connection_id"`
	PresentationRequest PresentationRequest `json:"presentation_request"`
}

type ProofRequestV1 struct {
	Comment             string                `json:"comment"`
	ConnectionID        string                `json:"connection_id"`
	PresentationRequest PresentationRequestV1 `json:"proof_request"`
}

type PresentationRequest struct {
	Indy Indy `json:"indy"`
}

type Indy struct {
	Name                string              `json:"name"`
	Version             string              `json:"version"`
	RequestedAttributes RequestedAttributes `json:"requested_attributes,omitempty"`
	RequestedPredicates RequestedPredicates `json:"requested_predicates"`
}

type PresentationRequestV1 struct {
	Name                string              `json:"name"`
	Version             string              `json:"version"`
	RequestedAttributes RequestedAttributes `json:"requested_attributes,omitempty"`
	RequestedPredicates RequestedPredicates `json:"requested_predicates"`
}

type RequestedAttributesAny struct {
	Attr RequestedAttributes
}

type RequestedAttributes struct {
	The0_IDNumberUUID       interface{} `json:"0_IDNumber_uuid,omitempty"`
	The0_ForenamesUUID      interface{} `json:"0_Forenames_uuid,omitempty"`
	The0_SurnameUUID        interface{} `json:"0_Surname_uuid,omitempty"`
	The0_GenderUUID         interface{} `json:"0_Gender_uuid,omitempty"`
	The0_DateOfBirthUUID    interface{} `json:"0_DateOfBirth_uuid,omitempty"`
	The0_CountryOfBirthUUID interface{} `json:"0_CountryOfBirth_uuid,omitempty"`
}

type The0___UUID struct {
	Name         string        `json:"name,omitempty"`
	Restrictions []Restriction `json:"restrictions,omitempty"`
}

type Restriction struct {
	CredDefID string `json:"cred_def_id,omitempty"`
}

type RequestedPredicates struct {
}

// type CreateProofPresentationRequest struct {
// 	Comment             string              `json:"comment,omitempty"`
// 	ConnectionID        string              `json:"connection_id,omitempty"`
// 	PresentationRequest PresentationRequest `json:"presentation_request,omitempty"`
// 	Trace               bool                `json:"trace,omitempty"`
// }

// type PresentationRequest struct {
// 	Indy Indy `json:"indy,omitempty"`
// }

// type Indy struct {
// 	Name                string                `json:"name,omitempty"`
// 	NonRevoked          NonRevoked            `json:"non_revoked,omitempty"`
// 	Nonce               string                `json:"nonce,omitempty"`
// 	RequestedAttributes []RequestedAttributes `json:"requested_attributes,omitempty"`
// 	RequestedPredicates RequestedPredicates   `json:"requested_predicates,omitempty"`
// 	Version             string                `json:"version,omitempty"`
// }

type NonRevoked struct {
	From int32 `json:"from,omitempty"`
	To   int32 `json:"to,omitempty"`
}

// type RequestedAttribute struct {
// 	Name         string         `json:"name,omitempty"`
// 	Names        []string       `json:"names,omitempty"`
// 	NonRevoked   NonRevoked     `json:"non_revoked,omitempty"`
// 	Restrictions []Restrictions `json:"restrictions,omitempty"`
// }

type Restrictions struct {
	CredDefID string `json:"cred_def_id,omitempty"`
}

// type RequestedPredicates struct {
// 	Name         string         `json:"name,omitempty"`
// 	NonRevoked   NonRevoked     `json:"non_revoked,omitempty"`
// 	PType        string         `json:"p_type,omitempty"`
// 	PValue       []int32        `json:"p_value,omitempty"`
// 	Restrictions []Restrictions `json:"restrictions,omitempty"`
// }

type CreateProofPresentationResponse struct {
	AutoPresent              bool                     `json:"auto_present,omitempty"`
	ConnectionID             string                   `json:"connection_id,omitempty"`
	CreatedAt                string                   `json:"created_at,omitempty"`
	ErrorMsg                 string                   `json:"error_msg,omitempty"`
	Initiator                string                   `json:"initiator,omitempty"`
	Presentation             Presentation             `json:"presentation,omitempty"`
	PresentationExchangeID   string                   `json:"presentation_exchange_id,omitempty"`
	PresentationProposalDict PresentationProposalDict `json:"presentation_proposal_dict,omitempty"`
	PresentationRequest      PresentationRequest1     `json:"presentation_request,omitempty"`
	PresentationRequestDict  PresentationRequestDict  `json:"presentation_request_dict,omitempty"`
	Role                     string                   `json:"role,omitempty"`
	State                    string                   `json:"state,omitempty"`
	ThreadID                 string                   `json:"thread_id,omitempty"`
	Trace                    bool                     `json:"trace,omitempty"`
	UpdatedAt                string                   `json:"updated_at,omitempty"`
	Verified                 string                   `json:"verified,omitempty"`
}

type Presentation struct {
	Identifiers    []Identifiers  `json:"identifiers,omitempty"`
	Proof          Proof          `json:"proof,omitempty"`
	RequestedProof RequestedProof `json:"requested_proof,omitempty"`
}

type Identifiers struct {
	CredeDefID string `json:"cred_def_id,omitempty"`
	RevRegID   string `json:"rev_reg_id,omitempty"`
	SchemaID   string `json:"schema_id,omitempty"`
	Timestamp  int32  `json:"timestamp,omitempty"`
}

type Proof struct {
	AggregatedProof AggregatedProof `json:"aggregated_proof,omitempty"`
	Proofs          []Proofs        `json:"proofs,omitempty"`
}

type AggregatedProof struct {
	CHash string `json:"c_hash,omitempty"`
	// CList int32 `json:"c_list,omitempty"`
}
type Proofs struct {
	NonRevocProof NonRevocProof `json:"non_revoc_proof,omitempty"`
	PrimaryProof  PrimaryProof  `json:"primary_proof,omitempty"`
}

type NonRevocProof struct {
	Clist Clist `json:"c_list,omitempty"`
	Xlist Xlist `json:"x_list,omitempty"`
}

type Clist struct{}

type Xlist struct{}

type PrimaryProof struct {
	EqProof  EqProof    `json:"eq_proof,omitempty"`
	GeProofs []GeProofs `json:"ge_proofs,omitempty"`
}

type EqProof struct {
	APrime        string        `json:"a_prime,omitempty"`
	E             string        `json:"e,omitempty"`
	M             M             `json:"m,omitempty"`
	M2            string        `json:"m2,omitempty"`
	RevealedAttrs RevealedAttrs `json:"revealed_attrs,omitempty"`
	V             string        `json:"v,omitempty"`
}

type RevealedAttrs struct {
	Encoded       string `json:"encoded,omitempty"`
	Raw           string `json:"raw,omitempty"`
	SubProofIndex int32  `json:"sub_proof_index,omitempty"`
}

type M struct{}

type GeProofs struct {
	Alpha     string    `json:"alpha,omitempty"`
	MJ        string    `json:"mj,omitempty"`
	Predicate Predicate `json:"predicate,omitempty"`
	R         R         `json:"r,omitempty"`
	T         T         `json:"t,omitempty"`
	U         U         `json:"u,omitempty"`
}

type Predicate struct {
	AttrName string `json:"attr_name,omitempty"`
	PType    string `json:"p_type,omitempty"`
	Value    int32  `json:"value,omitempty"`
}

type R struct{}

type T struct{}

type U struct{}

type RequestedProof struct {
	Predicates         Predicates         `json:"predicates,omitempty"`
	RevealedAttrGroups RevealedAttrGroups `json:"revealed_attr_groups,omitempty"`
	RevealedAttrs      RevealedAttrs      `json:"revealed_attrs,omitempty"`
	SelfAttestedAttrs  SelfAttestedAttrs  `json:"self_attested_attrs,omitempty"`
	UnrevealedAttrs    UnrevealedAttrs    `json:"unrevealed_attrs,omitempty"`
}

type Predicates struct {
	SubProofIndex int32 `json:"sub_proof_index,omitempty"`
}

type RevealedAttrGroups struct {
	SubProofIndex int32  `json:"sub_proof_index,omitempty"`
	Values        Values `json:"values,omitempty"`
}

type Values struct {
	Encoded string `json:"encoded,omitempty"`
	Raw     string `json:"raw,omitempty"`
}

type SelfAttestedAttrs struct{}

type UnrevealedAttrs struct{}

type PresentationProposalDict struct {
	ID                   string               `json:"@id,omitempty"`
	Type                 string               `json:"@type,omitempty"`
	Comment              string               `json:"comment,omitempty"`
	PresentationProposal PresentationProposal `json:"presentation_proposal,omitempty"`
}

type PresentationProposal struct {
	Attributes []Attributes  `json:"attributes,omitempty"`
	Predicates []Predicates1 `json:"predicates,omitempty"`
}

type Attributes struct {
	CredDefID string `json:"cred_def_id,omitempty"`
	MimeType  string `json:"mime-type,omitempty"`
	Name      string `json:"name,omitempty"`
	Referent  string `json:"referent,omitempty"`
	Value     string `json:"value,omitempty"`
}

type Predicates1 struct {
	CredDefID string `json:"cred_def_id,omitempty"`
	Name      string `json:"name,omitempty"`
	Predicate string `json:"predicate,omitempty"`
	Threshold int32  `json:"threshold,omitempty"`
}

type PresentationRequest1 struct {
	Name                string              `json:"name,omitempty"`
	NonRevoked          NonRevoked          `json:"non_revoked,omitempty"`
	Nonce               string              `json:"nonce,omitempty"`
	RequestedAttributes RequestedAttribute  `json:"requested_attributes,omitempty"`
	RequestedPredicates RequestedPredicates `json:"requested_predicates,omitempty"`
	Version             string              `json:"version,omitempty"`
}

type RequestedAttribute struct {
	Name         string         `json:"name,omitempty"`
	Names        []string       `json:"names,omitempty"`
	NonRevoked   NonRevoked     `json:"non_revoked,omitempty"`
	Restrictions []Restrictions `json:"restrictions,omitempty"`
}

type PresentationRequestDict struct {
	ID                         string                       `json:"@id,omitempty"`
	Type                       string                       `json:"@type,omitempty"`
	Comment                    string                       `json:"comment,omitempty"`
	RequestPresentationsAttach []RequestPresentationsAttach `json:"request_presentations~attach,omitempty"`
}

type ListProofRecordsParams struct {
	ConnectionID string `json:"connection_id,omitempty"`
	Role         string `json:"role,omitempty"`
	State        string `json:"state,omitempty"`
	ThreadID     string `json:"thread_id,omitempty"`
}

type ListProofRecordsResponse struct {
	Results []ResultsV1 `json:"results,omitempty"`
}

type ResultsV1 struct {
	AutoPresent     bool                     `json:"auto_present,omitempty"`
	ConnectionID    string                   `json:"connection_id,omitempty"`
	CreatedAt       string                   `json:"created_at,omitempty"`
	ErrorMsg        string                   `json:"error_msg,omitempty"`
	Initiator       string                   `json:"initiator,omitempty"`
	Presentation    Presentation             `json:"presentation,omitempty"`
	PresExID        string                   `json:"presentation_exchange_id,omitempty"`
	PresProposal    PresentationProposalDict `json:"presentation_proposal_dict,omitempty"`
	PresRequest     PresRequest              `json:"presentation_request,omitempty"`
	PresRequestDict PresentationRequestDict  `json:"presentation_request_dict,omitempty"`
	Role            string                   `json:"role,omitempty"`
	State           string                   `json:"state,omitempty"`
	ThreadID        string                   `json:"thread_id,omitempty"`
	Trace           bool                     `json:"trace,omitempty"`
	UpdatedAt       string                   `json:"updated_at,omitempty"`
	Verified        string                   `json:"verified,omitempty"`
}

type Results struct {
	AutoPresent  bool         `json:"auto_present,omitempty"`
	ByFormat     ByFormat     `json:"by_format,omitempty"`
	ConnectionID string       `json:"connection_id,omitempty"`
	CreatedAt    string       `json:"created_at,omitempty"`
	ErrorMsg     string       `json:"error_msg,omitempty"`
	Initiator    string       `json:"initiator,omitempty"`
	Pres         Pres         `json:"pres,omitempty"`
	PresExID     string       `json:"pres_ex_id,omitempty"`
	PresProposal PresProposal `json:"pres_proposal,omitempty"`
	PresRequest  PresRequest  `json:"pres_request,omitempty"`
	Role         string       `json:"role,omitempty"`
	State        string       `json:"state,omitempty"`
	ThreadID     string       `json:"thread_id,omitempty"`
	Trace        bool         `json:"trace,omitempty"`
	UpdatedAt    string       `json:"updated_at,omitempty"`
	Verifier     string       `json:"verifier,omitempty"`
}

type ByFormat struct {
	Pres         Pres         `json:"pres,omitempty"`
	PresProposal PresProposal `json:"pres_proposal,omitempty"`
	PresRequest  PresRequest  `json:"pres_request,omitempty"`
}

type Pres struct {
	ID                  string                `json:"@id,omitempty"`
	Type                string                `json:"@type,omitempty"`
	Comment             string                `json:"comment,omitempty"`
	Formats             []Formats             `json:"formats,omitempty"`
	PresentationsAttach []PresentationsAttach `json:"presentations~attach,omitempty"`
}

type PresProposal struct {
	ID              string            `json:"@id,omitempty"`
	Type            string            `json:"@type,omitempty"`
	Comment         string            `json:"comment,omitempty"`
	Formats         []Formats         `json:"formats,omitempty"`
	ProposalsAttach []ProposalsAttach `json:"proposals~attach,omitempty"`
}

type PresRequest struct {
	ID                         string                       `json:"@id,omitempty"`
	Type                       string                       `json:"@type,omitempty"`
	Comment                    string                       `json:"comment,omitempty"`
	Formats                    []Formats                    `json:"formats,omitempty"`
	RequestPresentationsAttach []RequestPresentationsAttach `json:"request_presentations~attach,omitempty"`
	WillConfirm                bool                         `json:"will_confirm,omitempty"`
}

type Formats struct {
	AttachID string `json:"attach_id,omitempty"`
	Format   string `json:"format,omitempty"`
}

type PresentationsAttach struct {
	ID          string `json:"@id,omitempty"`
	ByteCount   int32  `json:"byte_count,omitempty"`
	Data        Data   `json:"data,omitempty"`
	Description string `json:"description,omitempty"`
	Filename    string `json:"filename,omitempty"`
	LastmodTime string `json:"lastmod_time,omitempty"`
	MimeType    string `json:"mime-type,omitempty"`
}

type Data struct {
	Base64 string   `json:"base64,omitempty"`
	JSON   JSON     `json:"json,omitempty"`
	Jws    Jws      `json:"jws,omitempty"`
	Links  []string `json:"links,omitempty"`
	Sha256 string   `json:"sha256,omitempty"`
}

type JSON struct {
}

type Jws struct {
	Header     Header       `json:"header,omitempty"`
	Protected  string       `json:"protected,omitempty"`
	Signature  string       `json:"signature,omitempty"`
	Signatures []Signatures `json:"signatures,omitempty"`
}

type Header struct {
	Kid string `json:"kid,omitempty"`
}

type Signatures struct {
	Header    Header `json:"header,omitempty"`
	Protected string `json:"protected,omitempty"`
	Signature string `json:"signature,omitempty"`
}

type ProposalsAttach struct {
	ID          string `json:"@id,omitempty"`
	ByteCount   int32  `json:"byte_count,omitempty"`
	Data        Data   `json:"data,omitempty"`
	Description string `json:"description,omitempty"`
	Filename    string `json:"filename,omitempty"`
	LastmodTime string `json:"lastmod_time,omitempty"`
	MimeType    string `json:"mime-type,omitempty"`
}

type RequestPresentationsAttach struct {
	ID          string `json:"@id,omitempty"`
	ByteCount   int32  `json:"byte_count,omitempty"`
	Data        Data   `json:"data,omitempty"`
	Description string `json:"description,omitempty"`
	Filename    string `json:"filename,omitempty"`
	LastmodTime string `json:"lastmod_time,omitempty"`
	MimeType    string `json:"mime-type,omitempty"`
}
