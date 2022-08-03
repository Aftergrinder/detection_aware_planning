package sgtm

import (
	"context"
	"fmt"
	"os"
	"time"

	"google.golang.org/grpc/metadata"
	"moul.io/sgtm/pkg/sgtmpb"
)

func (svc *Service) claimsFromContext(ctx context.Context) (*jwtClaims, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("cannot get metadata from context")
	}

	oauthToken, ok := md["oa