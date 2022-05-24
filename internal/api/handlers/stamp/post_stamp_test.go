package stamp_test

import (
	"net/http"
	"stamp/internal/api"
	"stamp/internal/test"
	"stamp/internal/types"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPostStamp(t *testing.T) {
	test.WithTestServer(t, func(s *api.Server) {
		fix := test.Fixtures()

		payload := test.GenericPayload{
			"domainId": fix.Domain1.ID,
			"stampId":  fix.AnalyticsStamp.ID,
			"approved": true,
		}

		res := test.PerformRequest(t, s, "POST", "/api/v1/stamp", payload, test.HeadersWithAuth(t, fix.User1AccessToken1.Token))
		require.Equal(t, http.StatusOK, res.Result().StatusCode)

		var response types.PostStampResponse
		test.ParseResponseAndValidate(t, res, &response)

		test.Snapshoter.Skip([]string{"DomainStampID"}).Save(t, response)
	})
}

func TestPostStampExistingDomain(t *testing.T) {
	test.WithTestServer(t, func(s *api.Server) {
		fix := test.Fixtures()

		payload := test.GenericPayload{
			"domainId": fix.Domain2.ID,
			"stampId":  fix.AnalyticsStamp.ID,
			"approved": true,
			"rating":   9,
		}

		res := test.PerformRequest(t, s, "POST", "/api/v1/stamp", payload, test.HeadersWithAuth(t, fix.User1AccessToken1.Token))
		require.Equal(t, http.StatusOK, res.Result().StatusCode)

		var response types.PostStampResponse
		test.ParseResponseAndValidate(t, res, &response)

		test.Snapshoter.Save(t, response)
	})
}
