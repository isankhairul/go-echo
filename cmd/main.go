package main

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3filter"
	oapiMiddleware "github.com/oapi-codegen/echo-middleware"
	"os"
	"strings"

	"go-echo/generated"
	"go-echo/handler"
	"go-echo/repository"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func main() {
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
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(err)
	}

	// override with env vars
	viper.AutomaticEnv()
	viper.SetEnvPrefix("KD")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	e := echo.New()
	validatorOptions := &oapiMiddleware.Options{
		Options: openapi3filter.Options{
			AuthenticationFunc: handler.NewAuthenticator(),
		},
	}
	swagger, _ := generated.GetSwagger()
	e.Use(oapiMiddleware.OapiRequestValidatorWithOptions(swagger, validatorOptions))
	var server generated.ServerInterface = newServer()

	generated.RegisterHandlers(e, server)

	e.Logger.Fatal(e.Start(":1323"))
}

func newServer() *handler.Server {
	dbDsn := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=%s search_path=%s TimeZone=%s",
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.username"),
		viper.GetString("database.dbname"),
		viper.GetString("database.password"),
		"disable",
		viper.GetString("database.schemaname"),
		"Asia/Jakarta",
	)
	fmt.Println(dbDsn)
	var repo repository.RepositoryInterface = repository.NewRepository(repository.NewRepositoryOptions{
		Dsn: dbDsn,
	})
	opts := handler.NewServerOptions{
		Repository: repo,
	}
	return handler.NewServer(opts)
}
