package common

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"stamp/internal/api"
	"stamp/internal/config"
)

func GetVersionRoute(s *api.Server) *echo.Route {
	return s.Router.Management.GET("/version", getVersionHandler(s))
}

// Returns the version and build date baked into the binary.
func getVersionHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, config.GetFormattedBuildArgs())
	}
}
