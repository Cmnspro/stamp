package local

import (
	"context"
	"database/sql"
	"errors"
	"stamp/internal/data/dto"
	"stamp/internal/models"
	"stamp/internal/util"
	"stamp/internal/util/db"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (s Service) DeleteVote(ctx context.Context, userID, voteID string) error {
	log := util.LogFromContext(ctx).With().Str("voteID", voteID).Logger()

	vote, err := models.Votes(
		models.VoteWhere.ID.EQ(voteID),
		models.VoteWhere.UserID.EQ(userID),
	).One(ctx, s.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Trace().Msg("Vote no existent or already deleted")
			return nil
		}
		log.Err(err).Msg("Failed to find vote")
		return err
	}

	_, err = vote.Delete(ctx, s.db)
	if err != nil {
		log.Err(err).Msg("Failed to delete vote")
		return err
	}
	return nil
}

func (s *Service) StampDomain(ctx context.Context, user *models.User, filter dto.DomainStampFilter) (dto.Domain, error) {
	log := util.LogFromContext(ctx)

	if err := db.WithTransaction(ctx, s.db, func(exec boil.ContextExecutor) error {
		domainStamp, err := models.DomainStamps(
			models.DomainStampWhere.DomainID.EQ(filter.DomainID),
			models.DomainStampWhere.StampID.EQ(filter.StampID),
		).One(ctx, exec)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				log.Trace().Msg("No domain stamp found, creating new")
				domainStamp := &models.DomainStamp{
					DomainID: filter.DomainID,
					StampID:  filter.StampID,
				}

				if err := domainStamp.Insert(ctx, exec, boil.Infer()); err != nil {
					log.Err(err).Msg("Failed to insert new domain stamp")
					return err
				}

				vote := models.Vote{
					DomainStampID: domainStamp.ID,
					UserID:        user.ID,
					Approved:      filter.Approved,
					Rating:        null.Int64FromPtr((filter.Rating)),
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
			Approved:      filter.Approved,
			Rating:        null.Int64FromPtr((filter.Rating)),
		}

		if err := vote.Insert(ctx, exec, boil.Infer()); err != nil {
			log.Err(err).Msg("Failed to insert new vote")
			return err
		}
		return nil
	}); err != nil {
		log.Err(err).Msg("Failed to stamp domain")
		return dto.Domain{}, err
	}

	domain, err := s.GetDomain(ctx, filter.DomainID)
	if err != nil {
		log.Debug().Err(err).Msg("Failed to get domain")
		return dto.Domain{}, err
	}

	return domain, nil
}
