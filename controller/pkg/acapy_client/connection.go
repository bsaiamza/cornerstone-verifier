package acapy

import (
	"fmt"
	"strconv"

	"cornerstone_verifier/pkg/log"
	"cornerstone_verifier/pkg/models"
)

// CreateInvitation creates an invitation.
func (c *Client) CreateInvitation(request models.CreateInvitationRequest, params *models.CreateInvitationParams) (models.CreateInvitationResponse, error) {
	var invitation models.CreateInvitationResponse

	if params.Alias == "" {
		params.Alias = request.MyLabel
	}

	if !params.AutoAccept {
		params.AutoAccept = true
	}

	var queryParams = map[string]string{}
	if params != nil {
		queryParams = map[string]string{
			"alias":       params.Alias,
			"auto_accept": strconv.FormatBool(params.AutoAccept),
			"multi_use":   strconv.FormatBool(params.MultiUse),
			"public":      strconv.FormatBool(params.Public),
		}
	}

	err := c.post("/connections/create-invitation", queryParams, request, &invitation)
	if err != nil {
		log.Error.Printf("Failed on ACA-py /connections/create-invitation: %s", err.Error())
		return models.CreateInvitationResponse{}, err
	}
	return invitation, nil
}

// CreateMultiInvitation creates a multi connect invitation.
func (c *Client) CreateMultiInvitation(request models.CreateInvitationRequest, params *models.CreateInvitationParams) (models.CreateInvitationResponse, error) {
	var invitation models.CreateInvitationResponse

	if params.Alias == "" {
		params.Alias = request.MyLabel
	}

	if !params.AutoAccept {
		params.AutoAccept = true
	}

	if !params.MultiUse {
		params.MultiUse = true
	}

	var queryParams = map[string]string{}
	if params != nil {
		queryParams = map[string]string{
			"alias":       params.Alias,
			"auto_accept": strconv.FormatBool(params.AutoAccept),
			"multi_use":   strconv.FormatBool(params.MultiUse),
			"public":      strconv.FormatBool(params.Public),
		}
	}

	err := c.post("/connections/create-invitation", queryParams, request, &invitation)
	if err != nil {
		log.Error.Printf("Failed on ACA-py /connections/create-invitation: %s", err.Error())
		return models.CreateInvitationResponse{}, err
	}
	return invitation, nil
}

// ListConnections returns all connections.
func (c *Client) ListConnections(params *models.ListConnectionsParams) (models.ListConnectionsResponse, error) {
	var connections models.ListConnectionsResponse

	var queryParams = map[string]string{}
	if params != nil {
		queryParams = map[string]string{
			"alias":               params.Alias,
			"connection_protocol": params.ConnectionProtocol,
			"invitation_key":      params.InvitationKey,
			"my_did":              params.MyDID,
			"state":               params.State,
			"their_did":           params.TheirDID,
			"their_public_did":    params.TheirPublicDID,
			"their_role":          params.TheirRole,
		}
	}

	err := c.get("/connections", queryParams, &connections)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py /connections: ", err)
		return models.ListConnectionsResponse{}, err
	}
	return connections, nil
}

// PingConnection pings a connection.
func (c *Client) PingConnection(connectionID string) (models.PingConnectionResponse, error) {
	ping := struct {
		Comment string `json:"comment"`
	}{
		Comment: "ping",
	}

	var thread models.PingConnectionResponse

	err := c.post(fmt.Sprintf("/connections/%s/send-ping", connectionID), nil, ping, &thread)
	if err != nil {
		log.Error.Printf("Failed on ACA-py /connections/{conn_id}/send-ping: %s", err.Error())
		return models.PingConnectionResponse{}, err
	}
	return thread, nil
}
