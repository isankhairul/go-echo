package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-echo/generated"
	"go-echo/model/entity"
	"go-echo/model/request"
	"go-echo/util"
	"net/http"
)

// this register
// (POST /register)
func (s *Server) Register(ctx echo.Context) error {
	intStatusBadRQ := int(http.StatusBadRequest)
	respErr := generated.ErrorResponse{Code: &intStatusBadRQ}
	var input request.RegisterRQ
	err := ctx.Bind(&input)
	if err != nil {
		respErr.Message = err.Error()
		return ctx.JSON(http.StatusBadRequest, respErr)
	}

	// validation
	if err := input.Validate(); err != nil {
		respErr.Message = err.Error()
		return ctx.JSON(http.StatusBadRequest, respErr)
	}

	// check phone already exist
	countUser, err := s.Repository.UsersCountByPhone(context.Background(), input.Phone)
	if err != nil {
		respErr.Message = err.Error()
		return ctx.JSON(http.StatusBadRequest, respErr)
	}
	if countUser != nil && *countUser > 0 {
		intStatusConflict := int(http.StatusConflict)
		respErr.Message = "phone already exists"
		respErr.Code = &intStatusConflict
		return ctx.JSON(http.StatusConflict, respErr)
	}

	var resp generated.RegisterSuccessResponse
	user := entity.Users{}
	password, err := util.HashPassword([]byte(input.Password))
	if err != nil {
		respErr.Message = err.Error()
		return ctx.JSON(http.StatusBadRequest, respErr)
	}

	user.FullName = input.FullName
	user.Password = password
	user.Phone = input.Phone

	id, err := s.Repository.UsersCreate(context.Background(), user)
	if err != nil {
		respErr.Message = err.Error()
		return ctx.JSON(http.StatusBadRequest, respErr)
	}
	resp.Id = id

	return ctx.JSON(http.StatusOK, resp)
}
