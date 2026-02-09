package otvet

import (
	"fmt"

	"github.com/Polshkrev/gopolutils"
)

// Representation of an HTTP status code.
type StatusCode = uint16

const (
	StatusOK             StatusCode = 200 // RFC 9110, 15.3.1
	StatusCreated        StatusCode = 201 // RFC 9110, 15.3.2
	StatusAccepted       StatusCode = 202 // RFC 9110, 15.3.3
	StatusNoContent      StatusCode = 204 // RFC 9110, 15.3.5
	StatusPartialContent StatusCode = 206 // RFC 9110, 15.3.7

	StatusNotModified       StatusCode = 304 // RFC 9110, 15.4.5
	StatusUseProxy          StatusCode = 305 // RFC 9110, 15.4.6
	StatusTemporaryRedirect StatusCode = 307 // RFC 9110, 15.4.8
	StatusPermanentRedirect StatusCode = 308 // RFC 9110, 15.4.9

	StatusBadRequest            StatusCode = 400 // RFC 9110, 15.5.1
	StatusUnauthorized          StatusCode = 401 // RFC 9110, 15.5.2
	StatusPaymentRequired       StatusCode = 402 // RFC 9110, 15.5.3
	StatusForbidden             StatusCode = 403 // RFC 9110, 15.5.4
	StatusNotFound              StatusCode = 404 // RFC 9110, 15.5.5
	StatusMethodNotAllowed      StatusCode = 405 // RFC 9110, 15.5.6
	StatusNotAcceptable         StatusCode = 406 // RFC 9110, 15.5.7
	StatusRequestTimeout        StatusCode = 408 // RFC 9110, 15.5.9
	StatusConflict              StatusCode = 409 // RFC 9110, 15.5.10
	StatusRequestEntityTooLarge StatusCode = 413 // RFC 9110, 15.5.14
	StatusRequestURITooLong     StatusCode = 414 // RFC 9110, 15.5.15
	StatusUnsupportedMediaType  StatusCode = 415 // RFC 9110, 15.5.16
	StatusTeapot                StatusCode = 418 // RFC 9110, 15.5.19 (Unused)
	StatusUnprocessableEntity   StatusCode = 422 // RFC 9110, 15.5.21
	StatusLocked                StatusCode = 423 // RFC 4918, 11.3
	StatusTooManyRequests       StatusCode = 429 // RFC 6585, 4

	StatusInternalServerError           StatusCode = 500 // RFC 9110, 15.6.1
	StatusNotImplemented                StatusCode = 501 // RFC 9110, 15.6.2
	StatusBadGateway                    StatusCode = 502 // RFC 9110, 15.6.3
	StatusServiceUnavailable            StatusCode = 503 // RFC 9110, 15.6.4
	StatusGatewayTimeout                StatusCode = 504 // RFC 9110, 15.6.5
	StatusHTTPVersionNotSupported       StatusCode = 505 // RFC 9110, 15.6.6
	StatusInsufficientStorage           StatusCode = 507 // RFC 4918, 11.5
	StatusNetworkAuthenticationRequired StatusCode = 511 // RFC 6585, 6
)

// Determine the string value representation of a given status code.
// Returns the string representation for the given status code.
// If the string representation of the given status code can not be determined, a [gopolutils.ValueError] is returned with an empty string.
func StatusToString(code StatusCode) (string, *gopolutils.Exception) {
	switch code {
	case StatusOK:
		return "OK", nil
	case StatusCreated:
		return "Created", nil
	case StatusAccepted:
		return "Accepted", nil
	case StatusNoContent:
		return "No Content", nil
	case StatusPartialContent:
		return "Partial Content", nil
	case StatusNotModified:
		return "Not Modified", nil
	case StatusUseProxy:
		return "Use Proxy", nil
	case StatusTemporaryRedirect:
		return "Temporary Redirect", nil
	case StatusPermanentRedirect:
		return "Permanent Redirect", nil
	case StatusBadRequest:
		return "Bad Request", nil
	case StatusUnauthorized:
		return "Unauthorized", nil
	case StatusPaymentRequired:
		return "Payment Required", nil
	case StatusForbidden:
		return "Forbidden", nil
	case StatusNotFound:
		return "Not Found", nil
	case StatusMethodNotAllowed:
		return "Method Not Allowed", nil
	case StatusNotAcceptable:
		return "Not Acceptable", nil
	case StatusRequestTimeout:
		return "Request Timeout", nil
	case StatusConflict:
		return "Conflict", nil
	case StatusRequestEntityTooLarge:
		return "Request Entity Too Large", nil
	case StatusRequestURITooLong:
		return "Request URI Too Long", nil
	case StatusUnsupportedMediaType:
		return "Unsupported Media Type", nil
	case StatusTeapot:
		return "I'm a teapot", nil
	case StatusUnprocessableEntity:
		return "Unprocessable Entity", nil
	case StatusLocked:
		return "Locked", nil
	case StatusTooManyRequests:
		return "Too Many Requests", nil
	case StatusInternalServerError:
		return "Internal Server Error", nil
	case StatusNotImplemented:
		return "Not Implemented", nil
	case StatusBadGateway:
		return "Bad Gateway", nil
	case StatusServiceUnavailable:
		return "Service Unavailable", nil
	case StatusGatewayTimeout:
		return "Gateway Timeout", nil
	case StatusHTTPVersionNotSupported:
		return "HTTP Version Not Supported", nil
	case StatusInsufficientStorage:
		return "Insufficient Storage", nil
	case StatusNetworkAuthenticationRequired:
		return "Network Authentication Required", nil
	default:
		return "", gopolutils.NewNamedException(gopolutils.ValueError, fmt.Sprintf("Can not determine string value of status code: %d.\n", code))
	}
}
