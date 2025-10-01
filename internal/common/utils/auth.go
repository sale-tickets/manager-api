package utils

import (
	"context"

	"google.golang.org/grpc/metadata"
)

func GetProfileId(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	ids := md.Get("uuid")
	if len(ids) == 0 {
		return ""
	}

	return ids[0]
}
