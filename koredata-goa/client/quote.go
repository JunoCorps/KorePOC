// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "koredata": quote Resource Client
//
// Command:
// $ goagen
// --design=github.com/thefirstofthe300/kore-poc/koredata-goa/design
// --out=$(GOPATH)/src/github.com/thefirstofthe300/kore-poc/koredata-goa
// --version=v1.3.0

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

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
	return req, nil
}