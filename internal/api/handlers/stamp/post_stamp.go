package stamp

import (
	"net/http"
	"stamp/internal/api"
	"stamp/internal/api/auth"
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

		domainStamp, err := s.Local.StampDomain(ctx, user, payload)
		if err != nil {
			log.Debug().Err(err).Msg("Failed to stamp domain")
			return err
		}

		return util.ValidateAndReturn(c, http.StatusOK, domainStamp.ToTypes())
	}
}
