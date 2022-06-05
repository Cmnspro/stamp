package domain

import (
	"net/http"
	"stamp/internal/api"
	"stamp/internal/types/domain"
	"stamp/internal/util"

	"github.com/labstack/echo/v4"
)

func GetDomainRoute(s *api.Server) *echo.Route {
	return s.Router.APIV1Domain.GET("/:domainId", getDomainHandler(s))
}

func getDomainHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		log := util.LogFromContext(ctx)

		params := domain.NewGetDomainRouteParams()
		if err := util.BindAndValidatePathParams(c, &params); err != nil {
			return err
		}

		domain, err := s.Local.GetDomain(ctx, params.DomainID.String())
		if err != nil {
			log.Debug().Err(err).Msg("Failed to get domain")
			return err
		}

		return util.ValidateAndReturn(c, http.StatusOK, domain.ToTypes())
	}
}
