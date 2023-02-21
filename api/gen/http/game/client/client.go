// Code generated by goa v3.11.0, DO NOT EDIT.
//
// game client HTTP transport
//
// Command:
// $ goa gen github.com/comi91262/domilike/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the game service endpoint HTTP clients.
type Client struct {
	// Get Doer is the HTTP client used to make requests to the get endpoint.
	GetDoer goahttp.Doer

	// Create Doer is the HTTP client used to make requests to the create endpoint.
	CreateDoer goahttp.Doer

	// Delete Doer is the HTTP client used to make requests to the delete endpoint.
	DeleteDoer goahttp.Doer

	// GetSupplies Doer is the HTTP client used to make requests to the
	// get_supplies endpoint.
	GetSuppliesDoer goahttp.Doer

	// GetTrashes Doer is the HTTP client used to make requests to the get_trashes
	// endpoint.
	GetTrashesDoer goahttp.Doer

	// CORS Doer is the HTTP client used to make requests to the  endpoint.
	CORSDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the game service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		GetDoer:             doer,
		CreateDoer:          doer,
		DeleteDoer:          doer,
		GetSuppliesDoer:     doer,
		GetTrashesDoer:      doer,
		CORSDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Get returns an endpoint that makes HTTP requests to the game service get
// server.
func (c *Client) Get() goa.Endpoint {
	var (
		decodeResponse = DecodeGetResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("game", "get", err)
		}
		return decodeResponse(resp)
	}
}

// Create returns an endpoint that makes HTTP requests to the game service
// create server.
func (c *Client) Create() goa.Endpoint {
	var (
		decodeResponse = DecodeCreateResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCreateRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CreateDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("game", "create", err)
		}
		return decodeResponse(resp)
	}
}

// Delete returns an endpoint that makes HTTP requests to the game service
// delete server.
func (c *Client) Delete() goa.Endpoint {
	var (
		decodeResponse = DecodeDeleteResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDeleteRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("game", "delete", err)
		}
		return decodeResponse(resp)
	}
}

// GetSupplies returns an endpoint that makes HTTP requests to the game service
// get_supplies server.
func (c *Client) GetSupplies() goa.Endpoint {
	var (
		decodeResponse = DecodeGetSuppliesResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetSuppliesRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetSuppliesDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("game", "get_supplies", err)
		}
		return decodeResponse(resp)
	}
}

// GetTrashes returns an endpoint that makes HTTP requests to the game service
// get_trashes server.
func (c *Client) GetTrashes() goa.Endpoint {
	var (
		decodeResponse = DecodeGetTrashesResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetTrashesRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetTrashesDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("game", "get_trashes", err)
		}
		return decodeResponse(resp)
	}
}