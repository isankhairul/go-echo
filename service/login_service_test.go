package service

import (
	"context"
	"github.com/spf13/viper"
	"go-echo/generated"
	"go-echo/model/entity"
	"go-echo/repository"
	"go-echo/util"
	"go.uber.org/mock/gomock"
	"os"
	"strings"
	"testing"
)

func init() {
	// Load configuration
	viper.SetConfigType("yaml")
	var profile string = "dev"
	if os.Getenv("ENV") != "" {
		profile = "prd"
	}

	var configFileName []string
	configFileName = append(configFileName, "config-")
	configFileName = append(configFileName, profile)
	viper.SetConfigName(strings.Join(configFileName, ""))
	viper.AddConfigPath("../")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(err)
	}

	// override with env vars
	viper.AutomaticEnv()
	viper.SetEnvPrefix("KD")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func TestLoginFailedWrongPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//request
	input := generated.LoginBodyRequest{
		Phone:    "+6281293939831",
		Password: "Aab1222222222!",
	}
	userPassword := "Aab1222222222!!"
	userPasswordHash, _ := util.HashPassword([]byte(userPassword))
	user := entity.Users{
		Phone:    "+6281293939831",
		Password: userPasswordHash,
	}

	// Setup
	repo := repository.NewMockRepositoryInterface(ctrl)
	repo.EXPECT().UsersFirstByPhone(context.TODO(), input.Phone).Return(&user, nil).Times(1)
	svc := NewLoginService(repo)

	respSuccess, resError := svc.Login(context.TODO(), input)
	t.Log(resError.Message)

	if respSuccess != nil {
		t.Error(respSuccess)
	}
}

func TestLoginSuccess(t *testing.T) {
	// if go test config value must be ../private_key

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//request
	input := generated.LoginBodyRequest{
		Phone:    "+6281293939831",
		Password: "Aab1222222222!",
	}
	userPassword := "Aab1222222222!"
	userPasswordHash, _ := util.HashPassword([]byte(userPassword))
	user := entity.Users{
		Phone:    "+6281293939831",
		Password: userPasswordHash,
	}

	// Setup
	repo := repository.NewMockRepositoryInterface(ctrl)
	repo.EXPECT().UsersFirstByPhone(context.TODO(), input.Phone).Return(&user, nil).Times(1)
	svc := NewLoginService(repo)

	_, resError := svc.Login(context.TODO(), input)

	if resError != nil {
		t.Error(resError)
	}
}
