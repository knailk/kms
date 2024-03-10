package config

import (
	"context"
	"net/http"
	"time"

	"kms/app/entity"
	"kms/app/errs"
)

type contextKey string

const (
	appContextKey        = contextKey("app")
	contextKeyUser       = contextKey("user")
	authParamsContextKey = contextKey("authParams")
)

// NewContextWithApp returns a new context with the given App
func NewContextWithApp(ctx context.Context, a *entity.App) context.Context {
	return context.WithValue(ctx, appContextKey, a)
}

// AppFromRequest is a helper function which returns the App from the
// request context.
func AppFromRequest(r *http.Request) (*entity.App, error) {
	const op errs.Op = "entity/AppFromRequest"

	app, err := AppFromContext(r.Context())
	if err != nil {
		return nil, errs.E(op, err)
	}

	return app, nil
}

// AppFromContext returns the App from the given context
func AppFromContext(ctx context.Context) (*entity.App, error) {
	const op errs.Op = "entity/AppFromContext"

	a, ok := ctx.Value(appContextKey).(*entity.App)
	if !ok {
		return a, errs.E(op, errs.NotExist, "App not set to context")
	}
	return a, nil
}

// NewContextWithUser returns a new context with the given User
func NewContextWithUser(ctx context.Context, u *entity.User) context.Context {
	return context.WithValue(ctx, contextKeyUser, u)
}

// UserFromRequest returns the User from the request context
func UserFromRequest(r *http.Request) (u *entity.User, err error) {
	const op errs.Op = "entity/UserFromRequest"

	u, err = UserFromContext(r.Context())
	if err != nil {
		return nil, errs.E(op, err)
	}

	return u, nil
}

// UserFromContext returns the User from the given Context
func UserFromContext(ctx context.Context) (*entity.User, error) {
	const op errs.Op = "entity/UserFromContext"

	u, ok := ctx.Value(contextKeyUser).(*entity.User)
	if !ok {
		return nil, errs.E(op, errs.NotExist, "User not set properly to context")
	}
	return u, nil
}

// AuditFromRequest is a convenience function that sets up an Audit
// struct from the App and User set to the request context.
// The moment is also set to time.Now
func AuditFromRequest(r *http.Request) (adt entity.Audit, err error) {
	const op errs.Op = "entity/AuditFromRequest"

	var a *entity.App
	a, err = AppFromRequest(r)
	if err != nil {
		return entity.Audit{}, errs.E(op, err)
	}

	var u *entity.User
	u, err = UserFromRequest(r)
	if err != nil {
		return entity.Audit{}, errs.E(op, err)
	}

	adt.App = a
	adt.User = u
	adt.Moment = time.Now()

	return adt, nil
}

// NewContextWithAuthParams returns a new context with the given AuthenticationParams
func NewContextWithAuthParams(ctx context.Context, ap *entity.AuthenticationParams) context.Context {
	return context.WithValue(ctx, authParamsContextKey, ap)
}

// AuthParamsFromContext returns the AuthenticationParams from the given context
func AuthParamsFromContext(ctx context.Context) (*entity.AuthenticationParams, error) {
	const op errs.Op = "entity/AuthParamsFromContext"

	a, ok := ctx.Value(authParamsContextKey).(*entity.AuthenticationParams)
	if !ok {
		return a, errs.E(op, errs.NotExist, "Authentication Params not set to context")
	}
	return a, nil
}
