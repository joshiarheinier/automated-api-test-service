package middleware

import (
	"context"
	"github.com/joshia/automated-api-test-service/testapp/lib/uuid"
	"net/http"
)

func InjectRequestId(r *http.Request) *http.Request {
	ctx := r.Context()
	ctx = context.WithValue(ctx, "requestId", uuid.NewRequestId())
	r = r.WithContext(ctx)
	return r
}