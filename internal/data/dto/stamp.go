package dto

import (
	"stamp/internal/types"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/strfmt/conv"
)

type DomainStamp struct {
	DomainID string
	StampID  string
	Votes    VoteSlice
}

func (d DomainStamp) ToTypes() *types.PostStampResponse {
	return &types.PostStampResponse{
		DomainID: conv.UUID4(strfmt.UUID4(d.DomainID)),
		StampID:  conv.UUID4(strfmt.UUID4(d.StampID)),
		Votes:    d.Votes.ToTypes(),
	}
}
