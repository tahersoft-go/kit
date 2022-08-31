package jwd

import (
	"context"
	"errors"

	"github.com/golang-jwt/jwt"
)

func Claims(ctx context.Context) (*CustomClaims, error) {
	if user, ok := ctx.Value(UserContextKey).(*jwt.Token); ok {
		return user.Claims.(*CustomClaims), nil
	}
	return nil, errors.New("conversion to jwt.token failed for your token in context (Claims method)")
}
