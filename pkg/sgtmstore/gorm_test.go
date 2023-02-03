
package sgtmstore_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"moul.io/sgtm/pkg/sgtmpb"
	"moul.io/sgtm/pkg/sgtmstore"
)

func TestDBUserCreate(t *testing.T) {
	// logger := zapconfig.Configurator{}.MustBuild()
	logger := zap.NewNop()
	store := sgtmstore.TestingStore(t, logger)
	db := store.DB()

	tests := []struct {
		name            string
		input           *sgtmpb.User
		expectedOutput  *sgtmpb.User
		shouldHaveError bool
	}{
		{"nil", nil, nil, true},
		/*{
			"empty",
			&sgtmpb.User{},