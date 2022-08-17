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
