package acapy

import (
	"cornerstone_verifier/pkg/log"
	"cornerstone_verifier/pkg/models"
)

// CreateProofPresentation sends a proof request.
func (c *Client) CreateProofPresentation(request models.CreateProofPresentationRequest) (models.CreateProofPresentationResponse, error) {
	var proofPresentation models.CreateProofPresentationResponse
	err := c.post("/present-proof-2.0/send-request", nil, request, &proofPresentation)
	if err != nil {
		log.Error.Printf("Failed on ACA-py /present-proof-2.0/send-request: %s", err.Error())
		return models.CreateProofPresentationResponse{}, err
	}
	return proofPresentation, nil
}

// ListProofRecords returns all proof exchange records.
func (c *Client) ListProofRecords(params *models.ListProofRecordsParams) (models.ListProofRecordsResponse, error) {
	var proofRecords models.ListProofRecordsResponse

	var queryParams = map[string]string{}
	if params != nil {
		queryParams = map[string]string{
			"connection_id": params.ConnectionID,
			"role":          params.Role,
			"state":         params.State,
			"thread_id":     params.ThreadID,
		}
	}

	err := c.get("/present-proof-2.0/records", queryParams, &proofRecords)
	if err != nil {
		log.Error.Printf("Failed on ACA-py /present-proof-2.0/records: %s", err.Error())
		return models.ListProofRecordsResponse{}, err
	}
	return proofRecords, nil
}
