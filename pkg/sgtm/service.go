package sgtm

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"
	"moul.io/banner"

	"moul.io/sgtm/pkg/sgtmpb"
	"moul.io/sgtm/pkg/sgtmstore"
)

type Service struct {
	sgtmpb.UnimplementedWebAPIServer

	store         sg