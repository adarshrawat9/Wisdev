package repositories

import (
	"Wisdev/internal/model"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)


type UserRepository struct{
	db *pgxpool.Pool
}


func NewUserRepository(db *pgxpool.Pool) *UserRepository{
	return &UserRepository{
		db: db,
	}
}

func(r *UserRepository) CreateUser(user *model.User) error{

	query := `
		INSERT INTO users(
			username,
			email,
			password_hash
			)
			VALUES($1, $2, $3)
			RETURNING id, created_at, updated_at
		`

		return r.db.QueryRow(
			context.Background(),
			query,
			user.Username,
			user.Email,
			user.PasswordHash,
		).Scan(
			&user.ID,
			&user.CreatedAt,
			&user.UpdatedAt,
		)


	}


func (r *UserRepository)GetByEmail(email string)(*model.User, error){

	query := `
		SELECT 
		   id,
		   username,
		   email,
		   password_hash,
		   github_username,
		   portfolio_website,
		   avatar_url,
		   bio,
		   created_at,
		   updated_at
		   FROM users
		   WHERE email = $1
		   `

	var user model.User
	
	err := r.db.QueryRow(
		context.Background(),
	    query,
		email,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.GithubUsername,
		&user.PortfolioWebsite,
		&user.AvatarURL,
		&user.Bio,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil{
		return nil, err
	}

	return &user, nil
}	
