package acapy

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"cornerstone_verifier/pkg/log"
)

// Client ...
type Client struct {
	acapyURL                   string
	apiKey                     string
	tracing                    bool
	preserveExchangeRecords    bool
	autoRespondCredentialOffer bool
	HTTPClient                 http.Client
}

// NewClient ...
func NewClient(acapyURL string) *Client {
	return &Client{
		acapyURL:   strings.TrimRight(acapyURL, "/"),
		HTTPClient: http.Client{},
	}
}

// SetAPIKey ...
func (c *Client) SetAPIKey(apiKey string) *Client {
	c.apiKey = apiKey
	return c
}

// EnableTracing ...
func (c *Client) EnableTracing() *Client {
	c.tracing = true
	return c
}

// DisableTracing ...
func (c *Client) DisableTracing() *Client {
	c.tracing = false
	return c
}

// PreserveExchangeRecords ...
func (c *Client) PreserveExchangeRecords() *Client {
	c.preserveExchangeRecords = true
	return c
}

// AutoRespondCredentialOffer ...
func (c *Client) AutoRespondCredentialOffer() *Client {
	c.autoRespondCredentialOffer = true
	return c
}

func (c *Client) post(path string, queryParam map[string]string, body interface{}, response interface{}) error {
	return c.request(http.MethodPost, c.acapyURL+path, queryParam, body, response)
}

func (c *Client) get(path string, queryParams map[string]string, response interface{}) error {
	return c.request(http.MethodGet, c.acapyURL+path, queryParams, nil, response)
}

// func (c *Client) patch(path string, queryParams map[string]string, body interface{}, response interface{}) error {
// 	return c.request(http.MethodPatch, c.acapyURL+path, queryParams, body, response)
// }

// func (c *Client) put(path string) error {
// 	return c.request(http.MethodPut, c.acapyURL+path, nil, nil, nil)
// }

// func (c *Client) delete(url string) error {
// 	return c.request(http.MethodDelete, url, nil, nil, nil)
// }

func (c *Client) request(method string, url string, queryParams map[string]string, body interface{}, responseObject interface{}) error {
	var input io.Reader
	var err error

	if body != nil {
		jsonInput, err := json.Marshal(body)
		if err != nil {
			return err
		}
		input = bytes.NewReader(jsonInput)
	}

	r, err := http.NewRequest(method, url, input)
	if err != nil {
		return err
	}
	if c.apiKey != "" {
		r.Header.Add("X-API-KEY", c.apiKey)
	}
	r.Header.Add("Content-Type", "application/json")

	q := r.URL.Query()
	for k, v := range queryParams {
		if k != "" && v != "" {
			q.Add(k, v)
		}
	}
	r.URL.RawQuery = q.Encode()

	response, err := c.HTTPClient.Do(r)
	if err != nil || response.StatusCode >= 300 {
		if response != nil {
			log.Error.Printf("ACA-py Request failed: %s", response.Status)
			if body, err := ioutil.ReadAll(response.Body); err != nil {
				log.Error.Printf("ACA-py Response body: %s", body)
			}
			return errors.New("ACA-py Request failed " + err.Error())
		}
		return err
	}

	if responseObject != nil {
		err = json.NewDecoder(response.Body).Decode(responseObject)
		if err != nil {
			return err
		}
	}
	return nil
}
