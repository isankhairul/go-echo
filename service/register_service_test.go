package service

import (
	"context"
	"go-echo/generated"
	"go-echo/repository"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestRegisterFailedMinFullName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//request
	input := generated.RegisterBodyRequest{
		FullName: "fo",
		Phone:    "+6281293939831",
		Password: "Aab12222222222!",
	}

	// Setup
	repo := repository.NewMockRepositoryInterface(ctrl)
	svc := NewRegisterService(repo)

	respSuccess, resError := svc.Register(context.TODO(), input)
	t.Log(resError.Message)

	if respSuccess != nil {
		t.Error(respSuccess)
	}
}

func TestRegisterFailedFormatPhone(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//request
	input := generated.RegisterBodyRequest{
		FullName: "foo",
		Phone:    "+6121293939831",
		Password: "Aab12222222222!",
	}

	// Setup
	repo := repository.NewMockRepositoryInterface(ctrl)
	svc := NewRegisterService(repo)

	respSuccess, resError := svc.Register(context.TODO(), input)
	t.Log(resError.Message)

	if respSuccess != nil {
		t.Error(respSuccess)
	}
}

func TestRegisterFailedPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//request
	input := generated.RegisterBodyRequest{
		FullName: "foo",
		Phone:    "+6281293939831",
		Password: "Aab12222222222",
	}

	// Setup
	repo := repository.NewMockRepositoryInterface(ctrl)
	svc := NewRegisterService(repo)

	respSuccess, resError := svc.Register(context.TODO(), input)
	t.Log(resError.Message)

	if respSuccess != nil {
		t.Error(respSuccess)
	}
}

func TestRegisterFailedAlreadyExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	//request
	input := generated.RegisterBodyRequest{
		FullName: "foo",
		Phone:    "+6281293939831",
		Password: "Aab12222222222!",
	}
	var count int64 = 1

	// Setup
	repo := repository.NewMockRepositoryInterface(ctrl)
	repo.EXPECT().UsersCountByPhone(context.TODO(), input.Phone).Return(count, nil).Times(1)
	svc := NewRegisterService(repo)

	respSuccess, resError := svc.Register(context.TODO(), input)
	t.Log(resError.Message)

	if respSuccess != nil {
		t.Error(respSuccess)
	}
}
