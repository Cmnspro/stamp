package dto

import (
	"stamp/internal/types"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/strfmt/conv"
	"github.com/go-openapi/swag"
	"github.com/volatiletech/null/v8"
)

type DomainResponse struct {
	ID       string
	Domain   string
	ParentID null.String
}

func (d DomainResponse) ToTypes() *types.DomainResponse {
	res := &types.DomainResponse{
		ID:     conv.UUID4(strfmt.UUID4(d.ID)),
		Domain: swag.String(d.Domain),
	}
	if d.ParentID.Valid {
		res.ParentDomain = d.ParentID.String
	}
	return res
}

type Domain struct {
	DomainResponse
	Stamps DomainStampSlice
}

func (d Domain) ToTypes() *types.Domain {
	res := &types.Domain{
		DomainID: conv.UUID4(strfmt.UUID4(d.ID)),
		Domain:   swag.String(d.Domain),
		Stamps:   d.Stamps.ToTypes(),
	}
	if d.ParentID.Valid {
		res.ParentDomain = d.ParentID.String
	}
	return res
}
