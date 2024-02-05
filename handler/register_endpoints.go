package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-echo/generated"
	service2 "go-echo/service"
	"net/http"
)

// this register
// (POST /register)
func (s *Server) Register(ctx echo.Context) error {
	var input generated.RegisterBodyRequest
	err := ctx.Bind(&input)
	if err != nil {
		intStatusBadRQ := int(http.StatusBadRequest)
		respErr := generated.ErrorResponse{Code: &intStatusBadRQ}
		respErr.Message = err.Error()
		return ctx.JSON(http.StatusBadRequest, respErr)
	}

	service := service2.NewRegisterService(s.Repository)
	respSuccess, repError := service.Register(context.Background(), input)
	if repError != nil {
		return ctx.JSON(*repError.Code, *repError)
	}

	return ctx.JSON(http.StatusOK, *respSuccess)
}
