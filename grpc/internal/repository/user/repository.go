package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/prostasmosta/auth/grpc/internal/repository/user/converter"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/prostasmosta/auth/grpc/internal/repository"
	"github.com/prostasmosta/auth/grpc/internal/repository/user/model"

	grpcUser "github.com/prostasmosta/auth/grpc/pkg/user_v1"
)

const (
	tableName = "users"

	idColumn              = "id"
	nameColumn            = "name"
	emailColumn           = "email"
	roleColumn            = "role"
	passwordColumn        = "password"
	passwordConfirmColumn = "password_confirm"
	createdAtColumn       = "created_at"
	updatedAtColumn       = "updated_at"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, params *grpcUser.CreateRequest) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, roleColumn, passwordColumn, passwordConfirmColumn).
		Values(params.Info.Name, params.Info.Email, params.Info.Role, params.Password, params.PasswordConfirm).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	var id int64
	err = r.db.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*grpcUser.GetResponse, error) {
	builder := sq.Select(idColumn, nameColumn, emailColumn, roleColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var user model.GetUser
	err = r.db.QueryRow(ctx, query, args...).Scan(&user.ID, &user.Info.Name, &user.Info.Email, &user.Info.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}
