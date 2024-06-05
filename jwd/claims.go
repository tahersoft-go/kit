package jwd

import (
	"context"
	"errors"
	"log"

	"github.com/golang-jwt/jwt/v4"
)

func Claims(ctx context.Context) (jwt.MapClaims, error) {
	// Retrieve the token from the context
	token, ok := ctx.Value(UserContextKey).(*jwt.Token)
	if !ok || token == nil {
		// No token found, return nil claims without error
		log.Println("No valid JWT token found in context (Claims method)")
		return nil, nil
	}

	// Attempt to convert claims to jwt.MapClaims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("conversion to jwt.MapClaims failed for your token in context (Claims method)")
	}

	return claims, nil
}
