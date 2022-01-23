package database

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/otp-email/model"
)

func initDb() *pgx.Conn {
	conn, _ := pgx.Connect(context.Background(), "postgres://postgres:password@localhost:5432")
	return conn
}

func InsertUser(u *model.User) error {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:password@192.168.1.4:5432/accounts")
	defer conn.Close(context.Background())
	if err != nil {
		return err
	}
	_, err = conn.Exec(context.Background(), INSERT_USER, &u.Name, &u.Email, &u.Password, &u.Address, &u.IsVerified)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserVerification(email string) error {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:password@192.168.1.4:5432/accounts")
	defer conn.Close(context.Background())
	if err != nil {
		return err
	}
	_, err = conn.Exec(context.Background(), UPDATE_VERIFY_USER, true, email)
	if err != nil {
		return err
	}
	return nil
}

func SelectPasswordUser(email string) (string, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:password@192.168.1.4:5432/accounts")
	defer conn.Close(context.Background())
	if err != nil {
		return "", err
	}
	var password string
	if err = conn.QueryRow(context.Background(), SELECT_PASSWORD_USER, email).Scan(&password); err != nil {
		return "", err
	}
	return password, err
}

func SelectPasswordAdmin(email string) (string, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:password@192.168.1.4:5432/accounts")
	defer conn.Close(context.Background())
	if err != nil {
		return "", err
	}
	var password string
	if err = conn.QueryRow(context.Background(), SELECT_PASSWORD_ADMIN, email).Scan(&password); err != nil {
		return "", err
	}
	return password, err
}
