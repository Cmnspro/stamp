package mapper

import (
	"stamp/internal/data/dto"
	"stamp/internal/models"
)

func LocalDomainStampToDTO(domainStamp *models.DomainStamp) dto.DomainStamp {
	res := dto.DomainStamp{
		DomainID: domainStamp.DomainID,
		StampID:  domainStamp.StampID,
	}
	if domainStamp.R != nil && len(domainStamp.R.Votes) > 0 {
		res.Votes = LocalVoteSliceToDTO(domainStamp.R.Votes)
	}
	return res
}
