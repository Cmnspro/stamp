package stamp

import (
	"net/http"
	"stamp/internal/api"
	"stamp/internal/api/auth"
	"stamp/internal/data/dto"
	"stamp/internal/types"
	"stamp/internal/util"

	"github.com/labstack/echo/v4"
)

func PostStampRoute(s *api.Server) *echo.Route {
	return s.Router.APIV1Stamp.POST("", postStampHandler(s))
}

func postStampHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		log := util.LogFromContext(ctx)
		user := auth.UserFromContext(ctx)

		var payload types.PostStampPayload
		if err := util.BindAndValidateBody(c, &payload); err != nil {
			return err
		}

		filter := dto.DomainStampFilter{
			DomainID: payload.DomainID.String(),
			StampID:  payload.StampID.String(),
			Approved: *payload.Approved,
			Rating:   payload.Rating,
		}

		domain, err := s.Local.StampDomain(ctx, user, filter)
		if err != nil {
			log.Debug().Err(err).Msg("Failed to stamp domain")
			return err
		}

		return util.ValidateAndReturn(c, http.StatusOK, domain.ToTypes())
	}
}
