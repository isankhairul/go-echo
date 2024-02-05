package service

import (
	"context"
	"github.com/spf13/viper"
	"go-echo/model/entity"
	"go-echo/repository"
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

func TestDetailProfileFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//request
	var userIDRQ int64 = 2

	// Setup
	repo := repository.NewMockRepositoryInterface(ctrl)
	repo.EXPECT().UsersFirstByID(context.Background(), userIDRQ).Return(nil, nil).Times(1)
	svc := NewProfileService(repo)

	respSuccess, resError := svc.Detail(context.Background(), userIDRQ)

	if resError != nil {
		t.Log(*resError)
	}

	if respSuccess != nil {
		t.Log(*respSuccess)
		t.Error(respSuccess)
	}
}

func TestDetailProfileSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//request
	var userIDRQ int64 = 2
	user := entity.Users{
		ID:       2,
		Phone:    "+6281293939831",
		FullName: "test",
	}

	// Setup
	repo := repository.NewMockRepositoryInterface(ctrl)
	repo.EXPECT().UsersFirstByID(context.Background(), userIDRQ).Return(&user, nil).Times(1)
	svc := NewProfileService(repo)

	respSuccess, resError := svc.Detail(context.Background(), userIDRQ)

	if respSuccess != nil {
		t.Log(*respSuccess)
	}

	if resError != nil {
		t.Log(*resError)
		t.Error(resError)
	}
}
