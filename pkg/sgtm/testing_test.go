package sgtm

import (
	"context"
	"flag"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"moul.io/sgtm/pkg/sgtmstore"
	"moul.io/zapconfig"
)

var debugFlag = flag.Bool("debug", false, "more verbose logging")

func TestingService(t *testing.T) (Service, func()) {
	opts := Opts{
		Logger:     TestingLogger(t),
		ServerBind: "127.0.0.1:0",
	}
	opts.applyDefaults()
	store := sgtmstore.TestingStore(t, opts.Logger)
	ctx, cancel := context.WithCancel(opts.Context)
	svc := Servic