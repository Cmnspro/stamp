package stamp

import (
	"net/http"
	"stamp/internal/api"
	"stamp/internal/api/auth"
	"stamp/internal/types/stamp"
	"stamp/internal/util"

	"github.com/labstack/echo/v4"
)

func DeleteVoteRoute(s *api.Server) *echo.Route {
	return s.Router.APIV1Stamp.DELETE("/:voteId", deleteVoteHandler(s))
}

func deleteVoteHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		log := util.LogFromContext(ctx)
		user := auth.UserFromContext(ctx)

		params := stamp.NewDeleteVoteRouteParams()
		if err := util.BindAndValidatePathParams(c, &params); err != nil {
			return err
		}

		err := s.Local.DeleteVote(ctx, user.ID, params.VoteID.String())
		if err != nil {
			log.Debug().Err(err).Msg("Failed to delete vote")
			return err
		}

		return c.NoContent(http.StatusNoContent)
	}
}
