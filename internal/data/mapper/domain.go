package mapper

import (
	"stamp/internal/data/dto"
	"stamp/internal/models"
)

func LocalDomainToDTOResponse(domain *models.Domain) dto.DomainResponse {
	return dto.DomainResponse{
		ID:       domain.ID,
		Domain:   domain.Domain,
		ParentID: domain.ParentID,
	}
}

func LocalDomainToDTO(domain *models.Domain) dto.Domain {
	res := dto.Domain{
		DomainResponse: dto.DomainResponse{
			ID:       domain.ID,
			Domain:   domain.Domain,
			ParentID: domain.ParentID,
		},
	}
	if domain.R != nil && len(domain.R.DomainStamps) > 0 {
		res.Stamps = LocalDomainStampsToDTO(domain.R.DomainStamps)
	}
	return res
}
