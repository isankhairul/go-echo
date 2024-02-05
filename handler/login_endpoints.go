package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-echo/generated"
	service2 "go-echo/service"
	"net/http"
)

// (GET /login)
func (s *Server) Login(ctx echo.Context) error {
	intStatusBadRQ := int(http.StatusBadRequest)
	respErr := generated.ErrorResponse{Code: &intStatusBadRQ}
	var input generated.LoginBodyRequest
	err := ctx.Bind(&input)
	if err != nil {
		respErr.Message = err.Error()
		return ctx.JSON(http.StatusBadRequest, respErr)
	}

	service := service2.NewLoginService(s.Repository)
	respSuccess, repError := service.Login(context.Background(), input)
	if repError != nil {
		return ctx.JSON(*repError.Code, *repError)
	}

	return ctx.JSON(http.StatusOK, *respSuccess)
}
