// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import (
	"context"
	"go-echo/model/entity"
)

type RepositoryInterface interface {
	GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error)
	UsersCreate(ctx context.Context, Users entity.Users) (output *int64, err error)
	UsersFirstByPhone(ctx context.Context, phone string) (output *entity.Users, err error)
	UsersCountByPhone(ctx context.Context, phone string) (output int64, err error)
	UsersFirstByID(ctx context.Context, id int64) (output *entity.Users, err error)
	UsersUpdateByID(ctx context.Context, id int64, Users entity.Users) (output *int64, err error)
}
