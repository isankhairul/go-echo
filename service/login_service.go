package service

import (
	"context"
	"fmt"
	"go-echo/generated"
	"go-echo/model/request"
	"go-echo/repository"
	"go-echo/util"
	"net/http"
)

type LoginService struct {
	Repository repository.RepositoryInterface
}

func NewLoginService(repo repository.RepositoryInterface) *LoginService {
	return &LoginService{Repository: repo}
}

func (s LoginService) Login(ctx context.Context, input generated.LoginBodyRequest) (*generated.LoginSuccessResponse, *generated.ErrorResponse) {
	intStatusBadRQ := int(http.StatusBadRequest)
	respErr := generated.ErrorResponse{Code: &intStatusBadRQ}

	// validation
	if err := request.ValidateLoginBodyRequest(input); err != nil {
		respErr.Message = err.Error()
		return nil, &respErr
	}

	// get user
	user, err := s.Repository.UsersFirstByPhone(context.Background(), input.Phone)
	if err != nil {
		respErr.Message = err.Error()
		return nil, &respErr
	}
	if user == nil {
		respErr.Message = "phone not registered"
		return nil, &respErr
	}

	// check password
	if err := util.VerifyPassword([]byte(input.Password), []byte(user.Password)); err != nil {
		respErr.Message = "wrong password"
		return nil, &respErr
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
		return nil, &respErr
	}
	strAccessToken := string(accessToken)
	strRefreshToken := string(refreshToken)
	response := generated.LoginSuccessResponse{
		Id:           &user.ID,
		AccessToken:  &strAccessToken,
		RefreshToken: &strRefreshToken,
	}

	return &response, nil
}
