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

	oauthToken, ok := md["oauth-token"]
	if !ok || len(oauthToken) == 0 {
		return nil, fmt.Errorf("no such oauth-token")
	}

	return svc.parseJWTToken(oauthToken[0])
}

func (svc *Service) Me(ctx context.Context, req *sgtmpb.Me_Request) (*sgtmpb.Me_Response, error) {
	claims, err := svc.claimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	user, err := svc.store.GetUserByID(claims.Session.UserID)
	if err != nil {
		return nil, err
	}
	return &sgtmpb.Me_Response{User: user}, nil
}

func (svc *Service) Ping(context.Context, *sgtmpb.Pin