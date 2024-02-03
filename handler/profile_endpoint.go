package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-echo/generated"
	"go-echo/model/request"
	"net/http"
)

// (GET /profile)
func (s *Server) Profile(ctx echo.Context) error {
	intStatusForbiddenRQ := int(http.StatusForbidden)
	respErr := generated.ErrorResponse{Code: &intStatusForbiddenRQ}

	privateClaim := ctx.Get(JWTClaimsContextKey).(map[string]interface{})
	id := privateClaim["id"].(float64)

	// get user
	user, err := s.Repository.UsersFirstByID(context.Background(), uint64(id))
	if err != nil {
		respErr.Message = err.Error()
		return ctx.JSON(intStatusForbiddenRQ, respErr)
	}
	if user == nil {
		respErr.Message = "user not found"
		return ctx.JSON(intStatusForbiddenRQ, respErr)
	}

	response := generated.ProfileSuccessResponse{
		Phone:    &user.Phone,
		FullName: &user.FullName,
	}

	return ctx.JSON(http.StatusOK, response)
}

// (PUT /profile)
func (s *Server) UpdateProfile(ctx echo.Context) error {
	intStatusForbiddenRQ := int(http.StatusForbidden)
	respErr := generated.ErrorResponse{Code: &intStatusForbiddenRQ}

	var input request.UpdateProfileRQ
	err := ctx.Bind(&input)
	if err != nil {
		respErr.Message = err.Error()
		return ctx.JSON(http.StatusBadRequest, respErr)
	}

	privateClaim := ctx.Get(JWTClaimsContextKey).(map[string]interface{})
	id := privateClaim["id"].(float64)

	// get user
	user, err := s.Repository.UsersFirstByID(context.Background(), uint64(id))
	if err != nil {
		respErr.Message = err.Error()
		return ctx.JSON(intStatusForbiddenRQ, respErr)
	}
	if user == nil {
		respErr.Message = "user not found"
		return ctx.JSON(intStatusForbiddenRQ, respErr)
	}

	// check phone already exist
	countUser, err := s.Repository.UsersCountByPhone(context.Background(), input.Phone)
	if err != nil {
		respErr.Message = err.Error()
		return ctx.JSON(http.StatusBadRequest, respErr)
	}
	if countUser != nil && (*countUser > 0 && user.Phone != input.Phone) {
		intStatusConflict := int(http.StatusConflict)
		respErr.Message = "phone already exists"
		respErr.Code = &intStatusConflict
		return ctx.JSON(http.StatusConflict, respErr)
	}

	// update user
	user.Phone = input.Phone
	user.FullName = input.FullName
	_, err = s.Repository.UsersUpdateByID(context.Background(), uint64(id), *user)
	if err != nil {
		respErr.Message = err.Error()
		return ctx.JSON(http.StatusBadRequest, respErr)
	}

	response := generated.ProfileSuccessResponse{
		Phone:    &user.Phone,
		FullName: &user.FullName,
	}

	return ctx.JSON(http.StatusOK, response)
}
