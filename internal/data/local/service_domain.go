package local

import (
	"context"
	"database/sql"
	"errors"
	"stamp/internal/data/dto"
	"stamp/internal/data/mapper"
	"stamp/internal/models"
	"stamp/internal/types"
	"stamp/internal/util"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (s *Service) AddDomain(ctx context.Context, payload types.PostDomainPayload, user models.User) (dto.DomainResponse, error) {
	log := util.LogFromContext(ctx)

	existingDomain, err := models.Domains(
		models.DomainWhere.Domain.EQ(*payload.Domain),
	).One(ctx, s.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			domain := &models.Domain{
				Domain: *payload.Domain,
			}

			if payload.ParentDomainID.String() != "" {
				parentDomain, err := models.FindDomain(ctx, s.db, payload.ParentDomainID.String())
				if err != nil {
					if errors.Is(err, sql.ErrNoRows) {
						log.Debug().Err(err).Msg("Failed to find parent domain, inserting domain without parent domain")
					} else {
						log.Err(err).Msg("Failed to get parend domain")
						return dto.DomainResponse{}, err
					}
				}
				if parentDomain != nil {
					domain.ParentID = null.StringFrom(parentDomain.ID)
				}
			}

			err = domain.Insert(ctx, s.db, boil.Infer())
			if err != nil {
				log.Err(err).Msg("Failed to insert domain")
				return dto.DomainResponse{}, err
			}

			return mapper.LocalDomainToDTOResponse(domain), nil
		}
		log.Err(err).Msg("Failed to get domain")
	}
	return mapper.LocalDomainToDTOResponse(existingDomain), nil
}

func (s *Service) GetDomain(ctx context.Context, domainID string) (dto.Domain, error) {
	log := util.LogFromContext(ctx)

	domain, err := models.Domains(
		models.DomainWhere.ID.EQ(domainID),
		qm.Load(qm.Rels(models.DomainRels.DomainStamps, models.DomainStampRels.Stamp)),
		qm.Load(qm.Rels(models.DomainRels.DomainStamps, models.DomainStampRels.Votes)),
	).One(ctx, s.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Debug().Err(err).Msg("Domain not found")
			return dto.Domain{}, err
		}
		log.Err(err).Msg("Failed to get domain")
		return dto.Domain{}, err
	}

	return mapper.LocalDomainToDTO(domain), nil
}
