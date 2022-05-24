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
	"stamp/internal/util/db"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (s *Service) StampDomain(ctx context.Context, user *models.User, payload types.PostStampPayload) (dto.DomainStamp, error) {
	log := util.LogFromContext(ctx)

	if err := db.WithTransaction(ctx, s.db, func(exec boil.ContextExecutor) error {
		domainStamp, err := models.DomainStamps(
			models.DomainStampWhere.DomainID.EQ(payload.DomainID.String()),
			models.DomainStampWhere.StampID.EQ(payload.StampID.String()),
		).One(ctx, exec)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				log.Trace().Msg("No domain stamp found, creating new")
				domainStamp := &models.DomainStamp{
					DomainID: payload.DomainID.String(),
					StampID:  payload.StampID.String(),
				}

				if err := domainStamp.Insert(ctx, exec, boil.Infer()); err != nil {
					log.Err(err).Msg("Failed to insert new domain stamp")
					return err
				}

				vote := models.Vote{
					DomainStampID: domainStamp.ID,
					UserID:        user.ID,
					Approved:      *payload.Approved,
					Rating:        null.Int64FromPtr(payload.Rating),
				}

				if err := vote.Insert(ctx, exec, boil.Infer()); err != nil {
					log.Err(err).Msg("Failed to insert new vote")
					return err
				}
				return nil
			}
			log.Err(err).Msg("Failed to get domain stamp")
			return err
		}

		vote := models.Vote{
			DomainStampID: domainStamp.ID,
			UserID:        user.ID,
			Approved:      *payload.Approved,
			Rating:        null.Int64From(*payload.Rating),
		}

		if err := vote.Insert(ctx, exec, boil.Infer()); err != nil {
			log.Err(err).Msg("Failed to insert new vote")
			return err
		}
		return nil
	}); err != nil {
		log.Err(err).Msg("Failed to stamp domain")
		return dto.DomainStamp{}, err
	}

	domainStamp, err := models.DomainStamps(
		models.DomainStampWhere.DomainID.EQ(payload.DomainID.String()),
		models.DomainStampWhere.StampID.EQ(payload.StampID.String()),
		qm.Load(models.DomainStampRels.Votes),
	).One(ctx, s.db)
	if err != nil {
		log.Err(err).Msg("Failed to get domain stamp")
		return dto.DomainStamp{}, err
	}

	return mapper.LocalDomainStampToDTO(domainStamp), nil
}
