package service

import (
	"context"
	"go-echo/generated"
	"go-echo/model/request"
	"go-echo/repository"
	"net/http"
)

type ProfileService struct {
	Repository repository.RepositoryInterface
}

func NewProfileService(repo repository.RepositoryInterface) *ProfileService {
	return &ProfileService{Repository: repo}
}

func (s ProfileService) Detail(ctx context.Context, userID int64) (*generated.ProfileSuccessResponse, *generated.ErrorResponse) {
	intForbiddenRQ := int(http.StatusForbidden)
	respErr := generated.ErrorResponse{Code: &intForbiddenRQ}

	// get user
	user, err := s.Repository.UsersFirstByID(context.Background(), userID)
	if err != nil {
		respErr.Message = err.Error()
		return nil, &respErr
	}
	if user == nil {
		respErr.Message = "user not found"
		return nil, &respErr
	}

	response := generated.ProfileSuccessResponse{
		Phone:    &user.Phone,
		FullName: &user.FullName,
	}

	return &response, nil
}

func (s ProfileService) Update(ctx context.Context, userID int64, input generated.UpdateProfileBodyRequest) (*generated.ProfileSuccessResponse, *generated.ErrorResponse) {
	intForbiddenRQ := int(http.StatusForbidden)
	respErr := generated.ErrorResponse{Code: &intForbiddenRQ}

	// validation
	if err := request.ValidateUpdateProfileBodyRequest(input); err != nil {
		respErr.Message = err.Error()
		return nil, &respErr
	}

	// get user
	user, err := s.Repository.UsersFirstByID(context.Background(), userID)
	if err != nil {
		respErr.Message = err.Error()
		return nil, &respErr
	}
	if user == nil {
		respErr.Message = "user not found"
		return nil, &respErr
	}

	// check phone already exist
	countUser, err := s.Repository.UsersCountByPhone(context.Background(), input.Phone)
	if err != nil {
		respErr.Message = err.Error()
		return nil, &respErr
	}
	if countUser > 0 && user.Phone != input.Phone {
		intStatusConflict := int(http.StatusConflict)
		respErr.Message = "phone already exists"
		respErr.Code = &intStatusConflict
		return nil, &respErr
	}

	// update user
	user.Phone = input.Phone
	user.FullName = input.FullName
	_, err = s.Repository.UsersUpdateByID(context.Background(), userID, *user)
	if err != nil {
		respErr.Message = err.Error()
		return nil, &respErr
	}

	response := generated.ProfileSuccessResponse{
		Phone:    &user.Phone,
		FullName: &user.FullName,
	}

	return &response, nil
}
