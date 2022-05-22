package util_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"stamp/internal/util"
)

func TestRunningInTest(t *testing.T) {
	require.True(t, util.RunningInTest(), "Should be true while we are running in the test env/context")
}
