package errs

import (
	"encoding/json"
	"errors"
	"fmt"
	"kms/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// ErrResponse is used as the Response Body
type ErrResponse struct {
	Error ServiceError `json:"error"`
}

// ServiceError has fields for Service errors. All fields with no data will
// be omitted
type ServiceError struct {
	Kind    string `json:"kind,omitempty"`
	Code    string `json:"code,omitempty"`
	Param   string `json:"param,omitempty"`
	Message string `json:"message,omitempty"`
}

// HTTPErrorResponse takes a context from gin adn error, performs a
// type switch to determine if the type is an Error (which meets
// the Error interface as defined in this package), then sends the
// Error as a response to the client. If the type does not meet the
// Error interface as defined in this package, then a proper error
// is still formed and sent to the client, however, the Kind and
// Code will be Unanticipated. Logging of error is also done using
func HTTPErrorResponse(ctx *gin.Context, err error) {
	if err == nil {
		nilErrorResponse(ctx)
		return
	}

	var e *Error
	if errors.As(err, &e) {
		switch e.Kind {
		case Unauthenticated:
			unauthenticatedErrorResponse(ctx, e)
			return
		case Unauthorized:
			unauthorizedErrorResponse(ctx, e)
			return
		default:
			typicalErrorResponse(ctx, e)
			return
		}
	}

	unknownErrorResponse(ctx, err)
}

// typicalErrorResponse replies to the request with the specified error
// message and HTTP code. It does not otherwise end the request; the
// caller should ensure no further writes are done to w.
//
// Taken from standard library and modified.
// https://golang.org/pkg/net/http/#Error
func typicalErrorResponse(ctx *gin.Context, e *Error) {
	const op Op = "errs/typicalErrorResponse"

	httpStatusCode := httpErrorStatusCode(e.Kind)

	// We can retrieve the status here and write out a specific
	// HTTP status code. If the error is empty, just send the HTTP
	// Status Code as response. Error should not be empty, but it's
	// theoretically possible, so this is just in case...
	if e.isZero() {
		logger.ErrorF(fmt.Sprintf("error sent to %s, but empty - very strange, investigate", op), logrus.Fields{})
		ctx.Status(http.StatusInternalServerError)
		return
	}

	// typical errors
	const errMsg = "error response sent to client"

	ops := OpStack(e)
	if len(ops) > 0 {
		j, _ := json.Marshal(ops)
		// log the error with the op stack
		logger.ErrorF(errMsg, logrus.Fields{
			"HTTPStatusCode": httpStatusCode,
			"Kind":           e.Kind.String(),
			"Stack":          string(j),
			"Parameter":      string(e.Param),
			"Error":          e.Err.Error(),
			"Code":           string(e.Code),
		})
	} else {
		// no op stack present, log the error without that field
		logger.ErrorF(errMsg, logrus.Fields{
			"HTTPStatusCode": httpStatusCode,
			"Kind":           e.Kind.String(),
			"Parameter":      string(e.Param),
			"Error":          e.Err.Error(),
			"Code":           string(e.Code),
		})
	}

	// get ErrResponse
	er := newErrResponse(e)

	// Write Content-Type headers
	ctx.Header("Content-Type", "application/json")
	ctx.Header("X-Content-Type-Options", "nosniff")

	// Write  HTTP StatusCode and response body (json)
	ctx.JSON(httpStatusCode, er)
}

func newErrResponse(err *Error) ErrResponse {
	return ErrResponse{
		Error: ServiceError{
			Kind:    err.Kind.String(),
			Code:    string(err.Code),
			Param:   string(err.Param),
			Message: err.Error(),
		},
	}
}

// unauthenticatedErrorResponse responds with http status code 401
// (Unauthorized / Unauthenticated), an empty response body and a
// WWW-Authenticate header.
func unauthenticatedErrorResponse(ctx *gin.Context, e *Error) {
	if e.Realm == "" {
		e.Realm = "default"
	}

	ops := OpStack(e)
	if len(ops) > 0 {
		j, _ := json.Marshal(ops)
		// log the error with the op stack
		logger.ErrorF("unauthenticatedErrorResponse", logrus.Fields{
			"HTTPStatusCode": http.StatusUnauthorized,
			"Message":        "Unauthenticated Request",
			"Stack":          string(j),
			"Error":          e.Err.Error(),
		})
	} else {
		// no op stack present, log the error without that field
		logger.ErrorF("unauthenticatedErrorResponse", logrus.Fields{
			"HTTPStatusCode": http.StatusUnauthorized,
			"Message":        "Unauthenticated Request",
			"Realm":          string(e.Realm),
			"Error":          e.Err.Error(),
		})
	}

	ctx.Header("WWW-Authenticate", fmt.Sprintf(`Bearer realm="%s"`, e.Realm))
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": e.Err.Error()})
}

// unauthorizedErrorResponse responds with http status code 403 (Forbidden)
// and an empty response body.
func unauthorizedErrorResponse(ctx *gin.Context, e *Error) {

	ops := OpStack(e)
	if len(ops) > 0 {
		j, _ := json.Marshal(ops)
		// log the error with the op stack
		logger.ErrorF("unauthorizedErrorResponse", logrus.Fields{
			"HTTPStatusCode": http.StatusForbidden,
			"Message":        "Unauthorized Request",
			"Stack":          string(j),
			"Error":          e.Err.Error(),
		})
	} else {
		// no op stack present, log the error without that field
		logger.ErrorF("unauthorizedErrorResponse", logrus.Fields{
			"HTTPStatusCode": http.StatusForbidden,
			"Message":        "Unauthorized Request",
			"Error":          e.Err.Error(),
		})
	}

	ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": e.Err.Error()})
}

// nilErrorResponse responds with http status code 500 (Internal Server Error)
// and an empty response body. nil error should never be sent, but in case it is...
func nilErrorResponse(ctx *gin.Context) {
	logger.ErrorF("nilErrorResponse", logrus.Fields{
		"HTTPStatusCode": http.StatusInternalServerError,
		"Message":        "nil error - no response body sent",
	})

	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": " no response body sent"})
}

// unknownErrorResponse responds with http status code 500 (Internal Server Error)
// and a json response body with unanticipated_error kind
func unknownErrorResponse(ctx *gin.Context, err error) {
	er := ErrResponse{
		Error: ServiceError{
			Kind:    Unanticipated.String(),
			Code:    "Unanticipated",
			Message: "Unexpected error - contact support",
		},
	}

	logger.ErrorF("Unknown Error", logrus.Fields{"Error": err.Error()})

	// Write Content-Type headers
	ctx.Header("Content-Type", "application/json")
	ctx.Header("X-Content-Type-Options", "nosniff")

	// Write HTTP Statuscode and response body (json)
	ctx.JSON(http.StatusInternalServerError, er)
}

// httpErrorStatusCode maps an error Kind to an HTTP Status Code
func httpErrorStatusCode(k Kind) int {
	switch k {
	case Invalid, Exist, NotExist, Private, BrokenLink, Validation, InvalidRequest:
		return http.StatusBadRequest
	// the zero value of Kind is Other, so if no Kind is present
	// in the error, Other is used. Errors should always have a
	// Kind set, otherwise, a 500 will be returned and no
	// error message will be sent to the caller
	case Other, IO, Internal, Database, Unanticipated:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
