package client

import (
	"fmt"

	"iamza_verifier/pkg/log"
	"iamza_verifier/pkg/models"
)

func (c *Client) ListConnections() (models.ListConnectionsResponse, error) {
	queryParams := map[string]string{
		"alias": "IAMZA Verifier",
		"state": "active",
	}
	var connections models.ListConnectionsResponse

	arg := models.AcapyGetRequest{
		Endpoint:    "/connections",
		QueryParams: queryParams,
		Response:    &connections,
	}

	err := c.get(arg)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py /connections: ", err)
		return models.ListConnectionsResponse{}, err
	}
	return connections, nil
}

func (c *Client) CreateInvitation(request models.CreateInvitationRequest) (models.CreateInvitationResponse, error) {
	queryParams := map[string]string{
		"alias": "IAMZA Verifier",
		// "public": "true",
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

func (c *Client) CreateMultiUseInvitation(request models.CreateInvitationRequest) (models.CreateInvitationResponse, error) {
	queryParams := map[string]string{
		"alias":     "IAMZA Verifier",
		"multi_use": "true",
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

// NEW

func (c *Client) GetConnection(connID string) (models.Connection, error) {
	endpoint := fmt.Sprintf("/connections/%s", connID)
	var connection models.Connection

	arg := models.AcapyGetRequest{
		Endpoint: endpoint,
		Response: &connection,
	}

	err := c.get(arg)
	if err != nil {
		log.Error.Printf("Failed on ACA-py /connections/{conn_id}: %s", err.Error())
		return models.Connection{}, err
	}
	return connection, nil
}
