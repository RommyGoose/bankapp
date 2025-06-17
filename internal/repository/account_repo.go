package repository

import (
    "database/sql"
    "bankapp/internal/models"
)

type AccountRepo struct {
    DB *sql.DB
}

func NewAccountRepo(db *sql.DB) *AccountRepo {
    return &AccountRepo{DB: db}
}

func (r *AccountRepo) Create(userID int) (int, error) {
    var id int
    err := r.DB.QueryRow("INSERT INTO accounts (user_id, balance) VALUES ($1, 0) RETURNING id", userID).Scan(&id)
    return id, err
}
