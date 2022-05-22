package mapper

import (
	"stamp/internal/data/dto"
	"stamp/internal/models"
)

func LocalDomainToDTO(domain *models.Domain) dto.Domain {
	return dto.Domain{
		ID:       domain.ID,
		Domain:   domain.Domain,
		ParentID: domain.ParentID,
	}
}
