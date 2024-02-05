package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-echo/generated"
	service2 "go-echo/service"
	"net/http"
)

// (GET /profile)
func (s *Server) Profile(ctx echo.Context) error {
	privateClaim := ctx.Get(JWTClaimsContextKey).(map[string]interface{})
	id := privateClaim["id"].(float64)

	service := service2.NewProfileService(s.Repository)
	respSuccess, repError := service.Detail(context.Background(), int64(id))
	if repError != nil {
		return ctx.JSON(*repError.Code, *repError)
	}

	return ctx.JSON(http.StatusOK, *respSuccess)
}

// (PUT /profile)
func (s *Server) UpdateProfile(ctx echo.Context) error {
	intStatusForbiddenRQ := int(http.StatusForbidden)
	respErr := generated.ErrorResponse{Code: &intStatusForbiddenRQ}

	var input generated.UpdateProfileBodyRequest
	err := ctx.Bind(&input)
	if err != nil {
		respErr.Message = err.Error()
		return ctx.JSON(http.StatusBadRequest, respErr)
	}

	privateClaim := ctx.Get(JWTClaimsContextKey).(map[string]interface{})
	id := privateClaim["id"].(float64)

	service := service2.NewProfileService(s.Repository)
	respSuccess, repError := service.Update(context.Background(), int64(id), input)
	if repError != nil {
		return ctx.JSON(*repError.Code, *repError)
	}

	return ctx.JSON(http.StatusOK, *respSuccess)
}
