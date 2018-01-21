// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "koredata": Application User Types
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

package app

import (
	"github.com/goadesign/goa"
	"unicode/utf8"
)

// All quotes for a given user ID
type quote struct {
	// ID of the user
	ID *int `form:"ID,omitempty" json:"ID,omitempty" xml:"ID,omitempty"`
	// User ID of quoter
	Name *string `form:"Name,omitempty" json:"Name,omitempty" xml:"Name,omitempty"`
	// The actual quotes of the quoter
	Quote *string `form:"Quote,omitempty" json:"Quote,omitempty" xml:"Quote,omitempty"`
}

// Publicize creates Quote from quote
func (ut *quote) Publicize() *Quote {
	var pub Quote
	if ut.ID != nil {
		pub.ID = ut.ID
	}
	if ut.Name != nil {
		pub.Name = ut.Name
	}
	if ut.Quote != nil {
		pub.Quote = ut.Quote
	}
	return &pub
}

// All quotes for a given user ID
type Quote struct {
	// ID of the user
	ID *int `form:"ID,omitempty" json:"ID,omitempty" xml:"ID,omitempty"`
	// User ID of quoter
	Name *string `form:"Name,omitempty" json:"Name,omitempty" xml:"Name,omitempty"`
	// The actual quotes of the quoter
	Quote *string `form:"Quote,omitempty" json:"Quote,omitempty" xml:"Quote,omitempty"`
}

// quotePayload user type.
type quotePayload struct {
	Name  *string `form:"Name,omitempty" json:"Name,omitempty" xml:"Name,omitempty"`
	Quote *string `form:"Quote,omitempty" json:"Quote,omitempty" xml:"Quote,omitempty"`
}

// Validate validates the quotePayload type instance.
func (ut *quotePayload) Validate() (err error) {
	if ut.Name != nil {
		if utf8.RuneCountInString(*ut.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.Name`, *ut.Name, utf8.RuneCountInString(*ut.Name), 2, true))
		}
	}
	if ut.Quote != nil {
		if utf8.RuneCountInString(*ut.Quote) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.Quote`, *ut.Quote, utf8.RuneCountInString(*ut.Quote), 2, true))
		}
	}
	return
}

// Publicize creates QuotePayload from quotePayload
func (ut *quotePayload) Publicize() *QuotePayload {
	var pub QuotePayload
	if ut.Name != nil {
		pub.Name = ut.Name
	}
	if ut.Quote != nil {
		pub.Quote = ut.Quote
	}
	return &pub
}

// QuotePayload user type.
type QuotePayload struct {
	Name  *string `form:"Name,omitempty" json:"Name,omitempty" xml:"Name,omitempty"`
	Quote *string `form:"Quote,omitempty" json:"Quote,omitempty" xml:"Quote,omitempty"`
}

// Validate validates the QuotePayload type instance.
func (ut *QuotePayload) Validate() (err error) {
	if ut.Name != nil {
		if utf8.RuneCountInString(*ut.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`type.Name`, *ut.Name, utf8.RuneCountInString(*ut.Name), 2, true))
		}
	}
	if ut.Quote != nil {
		if utf8.RuneCountInString(*ut.Quote) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`type.Quote`, *ut.Quote, utf8.RuneCountInString(*ut.Quote), 2, true))
		}
	}
	return
}
