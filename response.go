package infra

import (
	"github.com/guilhermealegre/cleanarch-infra/errors"
	"net/http"
	"reflect"
)

type ErrorResponse struct {
	Status int     `json:"status"`
	Errors []error `json:"errors"`
}

// swagger:model Response
type Response struct {
	// Data
	Data interface{} `json:"data"`
}

//swagger:response SuccessResponse
type SuccessResponse struct {
	// Success
	Success bool `json:"success"`
}

func SetResponse(app IInfra, ctx *Context, err error, data interface{}) error {
	ctx.Set("Content-Type", "application/json")

	return ctx.JSON(getResponse(app, ctx, err, data))
}

func getResponse(app IInfra, ctx *Context, err error, data interface{}) interface{} {
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		errors := ErrorResponse{
			Status: http.StatusBadRequest,
		}
		errors.AddError(err)
		return errors

	} else {
		ctx.Status(http.StatusOK)
		response := Response{}
		response.Data = data
		return response
	}
}

func (e *ErrorResponse) AddError(err error) {
	if err != nil {
		if reflect.TypeOf(err) != reflect.TypeOf(errors.ErrorDetails{}) {
			err = errors.ErrorGeneric.Formats(err.Error())
		}
		e.Errors = append(e.Errors, err)
	}
}
