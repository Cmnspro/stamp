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
	if domainStamp.R != nil {
		if len(domainStamp.R.Votes) > 0 {
			res.Votes = LocalVoteSliceToDTO(domainStamp.R.Votes)
		}
		if domainStamp.R.Stamp != nil {
			res.StampName = domainStamp.R.Stamp.Name
		}
	}
	return res
}

func LocalDomainStampsToDTO(domainStamps models.DomainStampSlice) dto.DomainStampSlice {
	res := make(dto.DomainStampSlice, len(domainStamps))
	for i, domainStamp := range domainStamps {
		res[i] = LocalDomainStampToDTO(domainStamp)
	}
	return res
}
