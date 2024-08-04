package auth

import (
	"context"
	"encore.dev/beta/auth"
	"encore.dev/beta/errs"
)

var secrets struct {
	token string
}

//encore:authhandler
func AuthHandler(ctx context.Context, token string) (auth.UID, error) {
	uid := auth.UID(secrets.token)

	if token == secrets.token {
		return uid, nil
	}

	return "", &errs.Error{
		Code:    errs.Unauthenticated,
		Message: "invalid token",
	}
}
