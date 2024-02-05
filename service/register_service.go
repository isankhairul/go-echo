package service

import (
	"context"
	"go-echo/generated"
	"go-echo/model/entity"
	"go-echo/model/request"
	"go-echo/repository"
	"go-echo/util"
	"net/http"
)

type RegisterService struct {
	Repository repository.RepositoryInterface
}

func NewRegisterService(repo repository.RepositoryInterface) *RegisterService {
	return &RegisterService{Repository: repo}
}

func (s RegisterService) Register(ctx context.Context, input generated.RegisterBodyRequest) (*generated.RegisterSuccessResponse, *generated.ErrorResponse) {
	intStatusBadRQ := int(http.StatusBadRequest)
	respErr := generated.ErrorResponse{Code: &intStatusBadRQ}

	// validation
	if err := request.ValidateRegisterBodyRequest(input); err != nil {
		respErr.Message = err.Error()
		return nil, &respErr
	}

	// check phone already exist
	countUser, err := s.Repository.UsersCountByPhone(context.Background(), input.Phone)
	if err != nil {
		respErr.Message = err.Error()
		return nil, &respErr
	}
	if countUser > 0 {
		intStatusConflict := int(http.StatusConflict)
		respErr.Message = "phone already exists"
		respErr.Code = &intStatusConflict
		return nil, &respErr
	}

	var resp generated.RegisterSuccessResponse
	user := entity.Users{}
	password, err := util.HashPassword([]byte(input.Password))
	if err != nil {
		respErr.Message = err.Error()
		return nil, &respErr
	}
	user.FullName = input.FullName
	user.Password = password
	user.Phone = input.Phone

	id, err := s.Repository.UsersCreate(context.Background(), user)
	if err != nil {
		respErr.Message = err.Error()
		return nil, &respErr
	}
	resp.Id = id

	return &resp, nil
}
