package fhttp

import (
	"context"

	"github.com/fox-one/pigeon-sdk-go/utils/id"
)

const (
	HeaderKeyRequestID = "X-Request-Id"
)

type ContextKey string

const (
	ContextKeyBasicAuth  ContextKey = "ck_basic_auth"
	ContextKeyBearerAuth ContextKey = "ck_bearer_auth"
	ContextKeyRequestID  ContextKey = "ck_request_id"
)

type FContext struct {
	context.Context
}

func NewFContext(ctx context.Context) FContext {
	return FContext{
		Context: ctx,
	}
}

func (c *FContext) With(key interface{}, value interface{}) *FContext {
	c.Context = context.WithValue(c, key, value)
	return c
}

func (c *FContext) WithRequestID(requestID string) *FContext {
	c.Context = context.WithValue(c, ContextKeyRequestID, requestID)
	return c
}

func ContextWith(ctx context.Context, key interface{}, value interface{}) context.Context {
	return context.WithValue(ctx, key, value)
}

func ContextWithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, ContextKeyRequestID, requestID)
}

func RequestIdFromContext(ctx context.Context) string {
	if v, ok := ctx.Value(ContextKeyRequestID).(string); ok {
		return v
	}

	return id.GenUUIDString()
}
