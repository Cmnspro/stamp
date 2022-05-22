package local

import (
	"context"
	"database/sql"
	"errors"
	"stamp/internal/data/dto"
	"stamp/internal/data/mapper"
	"stamp/internal/models"
	"stamp/internal/util"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (s *Service) AddDomain(ctx context.Context, domain dto.Domain, user models.User) (dto.Domain, error) {
	log := util.LogFromContext(ctx)

	domainModel := new(models.Domain)
	domainModel.Domain = domain.Domain

	if domain.ParentID.String != "" {
		parentDomain, err := models.FindDomain(ctx, s.db, domain.ParentID.String)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				log.Debug().Err(err).Msg("Failed to find parent domain, inserting domain without parent domain")
			} else {
				log.Err(err).Msg("Failed to get parend domain")
				return dto.Domain{}, err
			}
		}
		if parentDomain != nil {
			domainModel.ParentID = null.StringFrom(parentDomain.ID)
		}
	}

	err := domainModel.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		log.Err(err).Msg("Failed to insert domain")
		return dto.Domain{}, err
	}

	return mapper.LocalDomainToDTO(domainModel), nil
}
