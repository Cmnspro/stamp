package dto

import (
	"stamp/internal/types"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/strfmt/conv"
	"github.com/go-openapi/swag"
	"github.com/volatiletech/null/v8"
)

type Domain struct {
	ID       string
	Domain   string
	ParentID null.String
}

func (d Domain) ToTypes() *types.Domain {
	res := &types.Domain{
		ID:     conv.UUID4(strfmt.UUID4(d.ID)),
		Domain: swag.String(d.Domain),
	}
	if d.ParentID.Valid {
		res.ParentDomain = d.ParentID.String
	}
	return res
}
