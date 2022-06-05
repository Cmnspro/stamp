package domain_test

import (
	"fmt"
	"net/http"
	"stamp/internal/api"
	"stamp/internal/test"
	"stamp/internal/types"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetDomain(t *testing.T) {
	test.WithTestServer(t, func(s *api.Server) {
		fix := test.Fixtures()

		res := test.PerformRequest(t, s, "GET", fmt.Sprintf("/api/v1/domains/%s", fix.Domain2.ID), nil, test.HeadersWithAuth(t, fix.User1AccessToken1.Token))
		require.Equal(t, http.StatusOK, res.Result().StatusCode)

		var response types.Domain
		test.ParseResponseAndValidate(t, res, &response)

		test.Snapshoter.Save(t, response)
	})
}
