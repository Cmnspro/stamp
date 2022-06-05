package dto

import (
	"stamp/internal/types"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/strfmt/conv"
	"github.com/go-openapi/swag"
)

type DomainStampFilter struct {
	DomainID string
	StampID  string
	Approved bool
	Rating   *int64
}

type DomainStamp struct {
	DomainID  string
	StampName string
	StampID   string
	Votes     VoteSlice
}

type DomainStampSlice []DomainStamp

func (d DomainStamp) ToTypes() *types.Stamp {
	return &types.Stamp{
		StampID:   conv.UUID4(strfmt.UUID4(d.StampID)),
		StampName: swag.String(d.StampName),
		Votes:     d.Votes.ToTypes(),
	}
}

func (ds DomainStampSlice) ToTypes() []*types.Stamp {
	res := make([]*types.Stamp, len(ds))
	for i, d := range ds {
		res[i] = d.ToTypes()
	}
	return res
}
