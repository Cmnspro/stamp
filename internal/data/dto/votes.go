package dto

import (
	"stamp/internal/types"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/strfmt/conv"
	"github.com/go-openapi/swag"
	"github.com/volatiletech/null/v8"
)

type Vote struct {
	Approved      bool
	DomainStampID string
	Rating        null.Int64
}

type VoteSlice []Vote

func (v Vote) ToTypes() *types.Vote {
	return &types.Vote{
		Approved:      swag.Bool(v.Approved),
		DomainStampID: conv.UUID4(strfmt.UUID4(v.DomainStampID)),
		Rating:        v.Rating.Ptr(),
	}
}

func (vs VoteSlice) ToTypes() []*types.Vote {
	res := make([]*types.Vote, len(vs))
	for i, v := range vs {
		res[i] = v.ToTypes()
	}
	return res
}
