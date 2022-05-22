package domain

import (
	"net/http"
	"stamp/internal/api"
	"stamp/internal/api/auth"
	"stamp/internal/data/dto"
	"stamp/internal/types"
	"stamp/internal/util"

	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
)

func PostDomainRoute(s *api.Server) *echo.Route {
	return s.Router.APIV1Domain.POST("", postDomainHanlder(s))
}

func postDomainHanlder(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		log := util.LogFromContext(ctx)
		user := auth.UserFromContext(ctx)

		body := types.PostDomainPayload{}
		if err := util.BindAndValidateBody(c, &body); err != nil {
			return err
		}

		domain := dto.Domain{
			Domain:   *body.Domain,
			ParentID: null.StringFrom(string(body.ParentDomainID)),
		}

		domain, err := s.Local.AddDomain(ctx, domain, *user)
		if err != nil {
			log.Err(err).Msg("Failed to add domain")
			return err
		}

		return util.ValidateAndReturn(c, http.StatusOK, domain.ToTypes())
	}
}
