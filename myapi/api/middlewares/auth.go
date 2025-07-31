package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/Danchitomoo/go_api_learning/apperrors"
	"github.com/Danchitomoo/go_api_learning/common"
	"google.golang.org/api/idtoken"
)

const (
	googleClientID = "myClientID"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorization := req.Header.Get("Authorization")

		authHeaders := strings.Split(authorization, " ")
		if len(authHeaders) != 2 {
			err := apperrors.RequiredAuthorizationHeader.Wrap(errors.New("incalid req header"), "invalid header")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		bearer, idToken := authHeaders[0], authHeaders[1]

		if bearer != "Bearer" || idToken == "" {
			err := apperrors.RequiredAuthorizationHeader.Wrap(errors.New("incalid req header"), "invalid header")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		tokenValidator, err := idtoken.NewValidator(context.Background())
		if err != nil {
			err = apperrors.CannotMakeValidator.Wrap(err, "internal auth error")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		payload, err := tokenValidator.Validate(context.Background(), idToken, googleClientID)
		if err != nil {
			err = apperrors.Unauthorizated.Wrap(err, "invalid id token")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		name, ok := payload.Claims["name"]
		if !ok {
			err = apperrors.Unauthorizated.Wrap(err, "invalid id token")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		req = common.SetUserName(req, name.(string))

		next.ServeHTTP(w, req)
	})
}
