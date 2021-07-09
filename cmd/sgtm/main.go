
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"syscall"
	"time"

	"github.com/Bearer/bearer-go"
	"github.com/bwmarrin/snowflake"
	"github.com/getsentry/sentry-go"
	"github.com/oklog/run"
	ff "github.com/peterbourgon/ff/v3"
	"github.com/peterbourgon/ff/v3/ffcli"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"moul.io/srand"
	"moul.io/zapconfig"
	"moul.io/zapgorm2"

	"moul.io/sgtm/internal/sgtmversion"
	"moul.io/sgtm/pkg/sgtm"
	"moul.io/sgtm/pkg/sgtmstore"
)

func main() {
	err := app(os.Args)
	switch {
	case err == nil:
	case err == run.SignalError{Signal: os.Interrupt}:
	default:
		log.Fatalf("error: %v", err)
		os.Exit(1)
	}
}

var svcOpts sgtm.Opts

func app(args []string) error {
	svcOpts = sgtm.DefaultOpts()
	rootFlags := flag.NewFlagSet("root", flag.ExitOnError)
	rootFlags.BoolVar(&svcOpts.DevMode, "dev-mode", svcOpts.DevMode, "start in developer mode")
	rootFlags.BoolVar(&svcOpts.EnableDiscord, "enable-discord", svcOpts.EnableDiscord, "enable discord bot")
	rootFlags.StringVar(&svcOpts.DiscordToken, "discord-token", svcOpts.DiscordToken, "discord bot token")
	rootFlags.StringVar(&svcOpts.DiscordAdminChannel, "discord-admin-channel", svcOpts.DiscordAdminChannel, "discord channel ID for admin messages")
	rootFlags.StringVar(&svcOpts.DBPath, "db-path", svcOpts.DBPath, "database path")
	rootFlags.BoolVar(&svcOpts.EnableServer, "enable-server", svcOpts.EnableServer, "enable HTTP+gRPC Server")
	rootFlags.StringVar(&svcOpts.ServerBind, "server-bind", svcOpts.ServerBind, "server bind (HTTP + gRPC)")
	rootFlags.StringVar(&svcOpts.ServerCORSAllowedOrigins, "server-cors-allowed-origins", svcOpts.ServerCORSAllowedOrigins, "allowed CORS origins")
	rootFlags.DurationVar(&svcOpts.ServerRequestTimeout, "server-request-timeout", svcOpts.ServerRequestTimeout, "server request timeout")
	rootFlags.DurationVar(&svcOpts.ServerShutdownTimeout, "server-shutdown-timeout", svcOpts.ServerShutdownTimeout, "server shutdown timeout")
	rootFlags.BoolVar(&svcOpts.ServerWithPprof, "server-with-pprof", svcOpts.ServerWithPprof, "enable pprof on HTTP server")
	rootFlags.StringVar(&svcOpts.DiscordClientID, "discord-client-id", svcOpts.DiscordClientID, "discord client ID (oauth)")
	rootFlags.StringVar(&svcOpts.DiscordClientSecret, "discord-client-secret", svcOpts.DiscordClientSecret, "discord client secret (oauth)")
	rootFlags.StringVar(&svcOpts.JWTSigningKey, "jwt-signing-key", svcOpts.JWTSigningKey, "HMAC secret to sign JWT tokens")
	rootFlags.StringVar(&svcOpts.Hostname, "hostname", svcOpts.Hostname, "I.e., https://sgtm.club")
	rootFlags.StringVar(&svcOpts.SoundCloudClientID, "soundcloud-client-id", svcOpts.SoundCloudClientID, "SoundCloud client ID")
	rootFlags.StringVar(&svcOpts.BearerToken, "bearer-token", svcOpts.BearerToken, "Bearer.sh token")
	rootFlags.StringVar(&svcOpts.IPFSAPI, "ipfs-api", svcOpts.IPFSAPI, "IPFS API multiaddress, if not provided or empry, will use the ipfs cli without an '--api' arg")
	rootFlags.BoolVar(&svcOpts.EnableProcessingWorker, "enable-processing-worker", svcOpts.EnableProcessingWorker, "enable processing worker")

	root := &ffcli.Command{
		FlagSet: rootFlags,
		Options: []ff.Option{
			ff.WithEnvVarPrefix("SGTM"),
			ff.WithConfigFile("config.txt"),
			ff.WithConfigFileParser(ff.PlainParser),
			ff.WithAllowMissingConfigFile(true),
		},