package jwd

import (
	"context"
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

func Claims(ctx context.Context) (jwt.MapClaims, error) {
	if ctx.Value(UserContextKey) == nil {
		return nil, errors.New("user context key not found")
	}
	if claims, ok := ctx.Value(UserContextKey).(*jwt.Token).Claims.(jwt.MapClaims); ok {
		return claims, nil
	}
	return nil, errors.New("conversion to jwt.token failed for your token in context (Claims method)")
}
