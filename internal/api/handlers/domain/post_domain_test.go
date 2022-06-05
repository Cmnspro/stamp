package domain_test

import (
	"net/http"
	"stamp/internal/api"
	"stamp/internal/test"
	"stamp/internal/types"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPostDomain(t *testing.T) {
	test.WithTestServer(t, func(s *api.Server) {
		fix := test.Fixtures()

		payload := test.GenericPayload{
			"domain": "duckduckgo.com",
		}

		res := test.PerformRequest(t, s, "POST", "/api/v1/domains", payload, test.HeadersWithAuth(t, fix.User1AccessToken1.Token))
		require.Equal(t, http.StatusOK, res.Result().StatusCode)

		var response types.DomainResponse
		test.ParseResponseAndValidate(t, res, &response)

		test.Snapshoter.Skip([]string{"ID"}).Save(t, response)
	})
}

func TestPostDomain_Subdomain(t *testing.T) {
	test.WithTestServer(t, func(s *api.Server) {
		fix := test.Fixtures()

		payload := test.GenericPayload{
			"domain":         "maps.google.com",
			"parentDomainID": fix.Domain1.ID,
		}

		res := test.PerformRequest(t, s, "POST", "/api/v1/domains", payload, test.HeadersWithAuth(t, fix.User1AccessToken1.Token))
		require.Equal(t, http.StatusOK, res.Result().StatusCode)

		var response types.DomainResponse
		test.ParseResponseAndValidate(t, res, &response)

		test.Snapshoter.Skip([]string{"ID"}).Save(t, response)
	})
}
