package handler

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go-echo/generated"
	"go-echo/model/request"
	"go-echo/util"
	"net/http"
)

// (GET /login)
func (s *Server) Login(ctx echo.Context) error {
	intStatusBadRQ := int(http.StatusBadRequest)
	respErr := generated.ErrorResponse{Code: &intStatusBadRQ}
	var input request.LoginRQ
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

	// get user
	user, err := s.Repository.UsersFirstByPhone(context.Background(), input.Phone)
	if err != nil {
		respErr.Message = err.Error()
		return ctx.JSON(http.StatusBadRequest, respErr)
	}
	if user == nil {
		respErr.Message = "phone not registered"
		return ctx.JSON(http.StatusBadRequest, respErr)
	}

	// check password
	if err := util.VerifyPassword([]byte(input.Password), []byte(user.Password)); err != nil {
		respErr.Message = "wrong password"
		return ctx.JSON(http.StatusBadRequest, respErr)
	}

	// generate token
	userMap := map[string]interface{}{
		"id":        user.ID,
		"phone":     user.Phone,
		"full_name": user.FullName,
	}
	accessToken, refreshToken, err := util.GenerateTokens(fmt.Sprint(user.ID), userMap)
	if err != nil {
		respErr.Message = err.Error()
		return ctx.JSON(http.StatusBadRequest, respErr)
	}
	strAccessToken := string(accessToken)
	strRefreshToken := string(refreshToken)

	response := generated.LoginSuccessResponse{
		Id:           &user.ID,
		AccessToken:  &strAccessToken,
		RefreshToken: &strRefreshToken,
	}

	return ctx.JSON(http.StatusOK, response)
}
