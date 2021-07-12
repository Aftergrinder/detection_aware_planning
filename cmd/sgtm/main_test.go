package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.T) {
	ctx := context.Background()
	svcOpts.EnableDiscord = false
	svcOpts.EnableServer