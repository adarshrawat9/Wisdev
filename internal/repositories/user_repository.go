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


func (r *UserRepository) GetById(userId string) (*model.User, error){

	var user model.User

	err := r.db.QueryRow(context.Background(),
	  ` SELECT id, username, email,
               portfolio_website, github_username,
               avatar_url, bio,
               created_at, updated_at
        FROM users
        WHERE id = $1
		`,userId).Scan(
			 &user.ID,
             &user.Username,
             &user.Email,
             &user.PortfolioWebsite,
             &user.GithubUsername,
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


func (r *UserRepository) UpdateUserDetails(user *model.User) (*model.User, error){

    query := `
	 	UPDATE users
		SET
		    bio = $1,
			github_username = $2,
			portfolio_website = $3,
    		avatar_url = $4,
    		updated_at = NOW()
			WHERE id = $5
			RETURNING
			id,
            username,
            email,
            portfolio_website,
            github_username,
            avatar_url,
            bio,
            created_at,
            updated_at 
			`


	err := r.db.QueryRow(context.Background(),
				query,
				user.Bio,
				user.GithubUsername,
				user.PortfolioWebsite,
				user.AvatarURL,
				user.ID,
			).Scan(
			 &user.ID,
             &user.Username,
             &user.Email,
             &user.PortfolioWebsite,
             &user.GithubUsername,
             &user.AvatarURL,
             &user.Bio,
             &user.CreatedAt,
             &user.UpdatedAt,
			)
	if err != nil{
		return nil, err
	}		
	return user, nil
}


func (r *UserRepository) GetByUsername(username string) (*model.User, error){

	var user model.User
	query := `
		SELECT id,
				username,
				email,
				password_hash,
				github_username,
				portfolio_website,
				avatar_url,
				bio,
				created_at,
				updated_at FROM users WHERE username = $1
			`

	err :=	r.db.QueryRow(context.Background(),
				query,
				username,
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
		return nil , err
	}		
	return &user, nil

}