package middleware

import (
	"context"
	"encoding/json"
	"github.com/joshia/automated-api-test-service/testapp/config/apperror"
	"github.com/joshia/automated-api-test-service/testapp/lib/message"
	"github.com/joshia/automated-api-test-service/testapp/lib/uuid"
	"log"
	"net/http"
)

func InjectRequestId(r *http.Request) *http.Request {
	ctx := r.Context()
	ctx = context.WithValue(ctx, "requestId", uuid.NewRequestId())
	r = r.WithContext(ctx)
	return r
}

func RewriteErrorResponse(b []byte) []byte {
	errv := &apperror.V1Error{}
	err := json.Unmarshal(b, errv)
	if err != nil {
		log.Fatalf(apperror.ErrFailedToDecodeConfigurationFile, err)
	}
	errRes := message.SetErrorResponse(errv)
	res, err := json.Marshal(errRes)
	if err != nil {
		log.Fatalf(apperror.ErrFailedToDecodeConfigurationFile, err)
	}
	return res
}