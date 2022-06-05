package stamp_test

import (
	"context"
	"fmt"
	"net/http"
	"stamp/internal/api"
	"stamp/internal/models"
	"stamp/internal/test"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeleteVote(t *testing.T) {
	test.WithTestServer(t, func(s *api.Server) {
		ctx := context.Background()
		fix := test.Fixtures()

		exists, err := models.VoteExists(ctx, s.DB, fix.Domain2AnalyticsStampVote1.ID)
		require.NoError(t, err)
		require.True(t, exists)

		res := test.PerformRequest(t, s, "DELETE", fmt.Sprintf("/api/v1/stamp/%s", fix.Domain2AnalyticsStampVote1.ID), nil, test.HeadersWithAuth(t, fix.User2AccessToken1.Token))
		require.Equal(t, http.StatusNoContent, res.Result().StatusCode)

		exists, err = models.VoteExists(ctx, s.DB, fix.Domain2AnalyticsStampVote1.ID)
		require.NoError(t, err)
		require.False(t, exists)
	})
}
