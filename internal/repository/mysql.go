package repository

import (
	"context"
	"database/sql"
	"fmt"
	"golang/config"
	"golang/internal/entity"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type repository struct {
	db *sql.DB
}

func NewUserRepository() *repository {
    cfg := config.LoadConfig()

    dsn := cfg.MySQL.DSN()
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("failed to open mysql connection: %v", err)
    }

    if err := db.Ping(); err != nil {
        log.Fatalf("failed to ping mysql: %v", err)
    }

    log.Println("Connected to MySQL successfully!")

    return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, user *entity.User) error {
    query := `
        INSERT INTO users (name, email)
        VALUES (?, ?)
    `
    res, err := r.db.ExecContext(ctx, query, user.Name, user.Email)
    if err != nil {
        return fmt.Errorf("failed to insert user: %w", err)
    }

    id, err := res.LastInsertId()
    if err != nil {
        return fmt.Errorf("failed to get inserted ID: %w", err)
    }

    user.Id = int(id)
    return nil
}

func (r *repository) Get(ctx context.Context, user *entity.User) ([]*entity.User, error) {
	query := `
		SELECT * FROM users
	`
	rows,err := r.db.QueryContext(ctx,query)
	if err != nil {
		return nil,fmt.Errorf("failed to get all users: %w", err)
	}

	defer rows.Close()

	var users []*entity.User


	for rows.Next() {
		var user *entity.User
		if err := rows.Scan(&user.Id, &user.Email, &user.Name); err != nil {
			return nil,fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return users, nil
}