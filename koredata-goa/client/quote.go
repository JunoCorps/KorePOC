// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "koredata": quote Resource Client
//
// Command:
// $ goagen
<<<<<<< HEAD
// --design=github.com/hegemone/kore-poc/koredata-goa/design
// --out=$(GOPATH)/src/github.com/hegemone/kore-poc/koredata-goa
// --version=v1.3.1
=======
// --design=github.com/thefirstofthe300/kore-poc/koredata-goa/design
// --out=$(GOPATH)/src/github.com/thefirstofthe300/kore-poc/koredata-goa
// --version=v1.3.0
>>>>>>> upstream/master

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// CreateQuotePayload is the quote create action payload.
type CreateQuotePayload struct {
	Name  string `form:"Name" json:"Name" xml:"Name"`
	Quote string `form:"Quote" json:"Quote" xml:"Quote"`
}

// CreateQuotePath computes a request path to the create action of quote.
func CreateQuotePath() string {

	return fmt.Sprintf("/quotes")
}

// Create a quote and add it to the database
func (c *Client) CreateQuote(ctx context.Context, path string, payload *CreateQuotePayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateQuoteRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateQuoteRequest create the request corresponding to the create action endpoint of the quote resource.
func (c *Client) NewCreateQuoteRequest(ctx context.Context, path string, payload *CreateQuotePayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// ListQuotePath computes a request path to the list action of quote.
func ListQuotePath() string {

	return fmt.Sprintf("/quotes")
}

// Returns all quotes in the quote database
func (c *Client) ListQuote(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListQuoteRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListQuoteRequest create the request corresponding to the list action endpoint of the quote resource.
func (c *Client) NewListQuoteRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListByIDQuotePath computes a request path to the list by ID action of quote.
func ListByIDQuotePath(userID string) string {
	param0 := userID

	return fmt.Sprintf("/quotes/%s", param0)
}

// Returns all the quotes for a given person
func (c *Client) ListByIDQuote(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListByIDQuoteRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListByIDQuoteRequest create the request corresponding to the list by ID action endpoint of the quote resource.
func (c *Client) NewListByIDQuoteRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
<<<<<<< HEAD
=======
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
>>>>>>> upstream/master
	return req, nil
}

// LoginQuotePath computes a request path to the login action of quote.
func LoginQuotePath() string {

	return fmt.Sprintf("/quotes/login")
}

// Login to the api
func (c *Client) LoginQuote(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewLoginQuoteRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewLoginQuoteRequest create the request corresponding to the login action endpoint of the quote resource.
func (c *Client) NewLoginQuoteRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.BasicAuthSigner != nil {
		if err := c.BasicAuthSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}
