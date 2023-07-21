package utils

import (
	coalago "github.com/coalalib/coalago"
)

func CodeHandler(code coalago.CoapCode) int {
	switch code {
	case coalago.GET:
		return 100
	case coalago.POST:
		return 100
	case coalago.PUT:
		return 100
	case coalago.DELETE:
		return 100
	case coalago.CoapCodeEmpty:
		return 204
	case coalago.CoapCodeCreated:
		return 201
	case coalago.CoapCodeDeleted:
		return 200
	case coalago.CoapCodeValid:
		return 200
	case coalago.CoapCodeChanged:
		return 200
	case coalago.CoapCodeContent:
		return 200
	case coalago.CoapCodeContinue:
		return 100
	case coalago.CoapCodeBadRequest:
		return 400
	case coalago.CoapCodeUnauthorized:
		return 401
	case coalago.CoapCodeBadOption:
		return 402
	case coalago.CoapCodeForbidden:
		return 403
	case coalago.CoapCodeNotFound:
		return 404
	case coalago.CoapCodeMethodNotAllowed:
		return 405
	case coalago.CoapCodeNotAcceptable:
		return 406
	case coalago.CoapCodePreconditionFailed:
		return 412
	case coalago.CoapCodeRequestEntityTooLarge:
		return 413
	case coalago.CoapCodeUnsupportedContentFormat:
		return 415
	case coalago.CoapCodeInternalServerError:
		return 500
	case coalago.CoapCodeNotImplemented:
		return 501
	case coalago.CoapCodeBadGateway:
		return 502
	case coalago.CoapCodeServiceUnavailable:
		return 503
	case coalago.CoapCodeGatewayTimeout:
		return 504
	case coalago.CoapCodeProxyingNotSupported:
		return 505
	default:
		return 500
	}
}
