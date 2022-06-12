package auth

import (
	"context"
	"crypto/rsa"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

const authHeader = "Tfs-Auth-Token"

type ctxMarker struct{}

var ctxMarkerKey = &ctxMarker{}

type User struct {
	Login string
}

var AnonymousUser User

func UserFromContext(ctx context.Context) User {
	u, ok := ctx.Value(ctxMarkerKey).(*User)
	if !ok || u == nil {
		return AnonymousUser
	}

	return *u
}

func Middleware(publicKey *rsa.PublicKey) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			signedToken := r.Header.Get(authHeader)

			var claims jwt.RegisteredClaims
			_, err := jwt.ParseWithClaims(signedToken, &claims,
				func(t *jwt.Token) (interface{}, error) {
					return publicKey, nil
				},
			)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			user := &User{
				Login: claims.ID,
			}
			authCtx := context.WithValue(
				r.Context(),
				ctxMarkerKey,
				user,
			)

			handler.ServeHTTP(w, r.WithContext(authCtx))
		}

		return http.HandlerFunc(fn)
	}
}

func WriteToken(w http.ResponseWriter, u *User, privateKey *rsa.PrivateKey) error {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.RegisteredClaims{
		ID: u.Login,
	})

	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return fmt.Errorf("can't sign token: %w", err)
	}

	w.Header().Add(authHeader, signedToken)
	return nil
}
