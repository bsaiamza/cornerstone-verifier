package acapy

import (
	"fmt"

	"iamza-verifier/pkg/log"
	"iamza-verifier/pkg/models"
)

func (c *Client) CreateInvitation(request models.CreateInvitationRequest) (models.CreateInvitationResponse, error) {
	queryParams := map[string]string{
		"alias": "Physical Address Issuer",
	}
	var invitation models.CreateInvitationResponse

	arg := models.AcapyPostRequest{
		Endpoint:    "/connections/create-invitation",
		QueryParams: queryParams,
		Body:        request,
		Response:    &invitation,
	}

	err := c.post(arg)
	if err != nil {
		log.Error.Printf("Failed on ACA-py /connections/create-invitation: %s", err.Error())
		return models.CreateInvitationResponse{}, err
	}
	return invitation, nil
}

func (c *Client) PingConnection(connectionID string, request models.PingConnectionRequest) (models.PingConnectionResponse, error) {
	endpoint := fmt.Sprintf("/connections/%s/send-ping", connectionID)
	var thread models.PingConnectionResponse

	arg := models.AcapyPostRequest{
		Endpoint: endpoint,
		Body:     request,
		Response: &thread,
	}

	err := c.post(arg)
	if err != nil {
		log.Error.Printf("Failed on ACA-py /connections/{conn_id}/send-ping: %s", err.Error())
		return models.PingConnectionResponse{}, err
	}
	return thread, nil
}
