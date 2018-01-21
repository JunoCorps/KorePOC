// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "koredata": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/hegemone/kore-poc/koredata-goa/design
// --out=$(GOPATH)/src/github.com/hegemone/kore-poc/koredata-goa
// --version=v1.3.1

package app

// All quotes for a given user ID (default view)
//
// Identifier: vnd.application.io/quote; view=default
type Quote struct {
	// ID of the user
	ID *int `form:"ID,omitempty" json:"ID,omitempty" xml:"ID,omitempty"`
	// User ID of quoter
	Name *string `form:"Name,omitempty" json:"Name,omitempty" xml:"Name,omitempty"`
	// The actual quotes of the quoter
	Quote *string `form:"Quote,omitempty" json:"Quote,omitempty" xml:"Quote,omitempty"`
}

// A quote from the user database (default view)
//
// Identifier: vnd.application.io/quotes; view=default
type Quotes struct {
	// Quote
	Quotes []*Quote `form:"Quotes,omitempty" json:"Quotes,omitempty" xml:"Quotes,omitempty"`
}

// Suggestions media type (default view)
//
// Identifier: vnd.application.io/suggestions; view=default
type Suggestions struct {
	// Suggestion
	Suggestions []*Suggestion `form:"Suggestions,omitempty" json:"Suggestions,omitempty" xml:"Suggestions,omitempty"`
}
