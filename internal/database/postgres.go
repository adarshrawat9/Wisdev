package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)


var DB *pgxpool.Pool
func ConnectDB() error{

	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),

	)


	pool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil{
		return err
	}

	if err := pool.Ping(context.Background()); err != nil{
		pool.Close()
		return err
	}

	DB = pool

	return nil
}