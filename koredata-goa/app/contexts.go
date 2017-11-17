// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "koredata": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/thefirstofthe300/kore-poc/koredata-goa/design
// --out=$(GOPATH)/src/github.com/thefirstofthe300/kore-poc/koredata-goa
// --version=v1.3.0

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// ListQuoteContext provides the quote list action context.
type ListQuoteContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListQuoteContext parses the incoming request URL and body, performs validations and creates the
// context used by the quote controller list action.
func NewListQuoteContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListQuoteContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListQuoteContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListQuoteContext) OK(r *JSON) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// ListByIDQuoteContext provides the quote list by ID action context.
type ListByIDQuoteContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	UserID string
}

// NewListByIDQuoteContext parses the incoming request URL and body, performs validations and creates the
// context used by the quote controller list by ID action.
func NewListByIDQuoteContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListByIDQuoteContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListByIDQuoteContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramUserID := req.Params["userId"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		rctx.UserID = rawUserID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListByIDQuoteContext) OK(r *JSON) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListByIDQuoteContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}
