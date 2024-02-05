package repository

import (
	"context"
	"go-echo/model/entity"
)

func (r *Repository) GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT name FROM test WHERE id = $1", input.Id).Scan(&output.Name)
	if err != nil {
		return
	}
	return
}

func (r *Repository) UsersFirstByPhone(ctx context.Context, phone string) (output *entity.Users, err error) {
	var users entity.Users
	err = r.Db.QueryRowContext(ctx, "SELECT id, phone, full_name, password FROM users WHERE phone = $1", phone).Scan(&users.ID, &users.Phone, &users.FullName, &users.Password)
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *Repository) UsersFirstByID(ctx context.Context, id int64) (output *entity.Users, err error) {
	var users entity.Users
	err = r.Db.QueryRowContext(ctx, "SELECT id, phone, full_name, password FROM users WHERE id = $1", id).Scan(&users.ID, &users.Phone, &users.FullName, &users.Password)
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *Repository) UsersCreate(ctx context.Context, Users entity.Users) (output *int64, err error) {
	var UsersID int64
	err = r.Db.QueryRowContext(ctx,
		`INSERT INTO users (phone, password, full_name) 
				VALUES($1, $2, $3) RETURNING id`, Users.Phone, Users.Password, Users.FullName).
		Scan(&UsersID)

	if err != nil {
		return nil, err
	}
	return &UsersID, nil
}

func (r *Repository) UsersUpdateByID(ctx context.Context, id int64, Users entity.Users) (output *int64, err error) {
	var UsersID int64
	_, err = r.Db.ExecContext(ctx,
		`UPDATE users 
				SET  phone=$1, full_name=$2, updated_at=now()
				WHERE id = $3`, Users.Phone, Users.FullName, id)
	if err != nil {
		return nil, err
	}
	return &UsersID, nil
}

func (r *Repository) UsersCountByPhone(ctx context.Context, phone string) (output int64, err error) {
	var count int64
	err = r.Db.QueryRowContext(ctx, "SELECT count(1) FROM users WHERE phone = $1", phone).Scan(&count)
	if err != nil {
		return count, err
	}
	return count, nil
}
