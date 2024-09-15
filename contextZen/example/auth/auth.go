package auth

import (
	"context"
	"errors"
)

type ctxKey int

const (
	authToken ctxKey = iota
)

func SetAuthToken(ctx context.Context, token string) context.Context {
	tokenCtx := context.WithValue(ctx, authToken, token)
	return tokenCtx
}

func getAuthToken(ctx context.Context) (string, error) {
	if token, ok := ctx.Value(authToken).(string); ok {
		return token, nil
	}

	return "", errors.New("cannot find auth token")
}

func VerifyAuthToken(ctx context.Context) (int, error) {
	token, err := getAuthToken(ctx)
	if err != nil {
		return 0, err
	}

	userID := len(token)
	if userID < 3 {
		return 0, errors.New("forbidden")
	}

	return userID, nil
}
