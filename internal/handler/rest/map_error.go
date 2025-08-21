package rest

import (
	"net/http"
	"whatsapp-api/internal/handler"
	"whatsapp-api/internal/service"
	"whatsapp-api/model"
	"whatsapp-api/model/constant"
)

// ErrorResponseMap maps error type to response code and HTTP status
func ErrorResponseMap(err error) (httpStatus int, errResp model.ErrorResponse) {
	var code string
	if serr, ok := err.(*service.ServiceError); ok {
		switch serr.Type {
		case service.ValidationError:
			code = constant.ValidationError
			httpStatus = http.StatusBadRequest
		case service.OtherError:
			code = constant.OtherError
			httpStatus = http.StatusInternalServerError
		}
	}

	if gerr, ok := err.(*handler.GrpcError); ok {
		switch gerr.Type {
		case handler.ConnectionError:
			code = constant.InternalError
			httpStatus = http.StatusInternalServerError
		}
	}

	// Default error
	return httpStatus, model.ErrorResponse{
		Error: model.Error{
			Code:    code,
			Message: err.Error(),
		},
	}
}
