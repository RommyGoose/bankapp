package repository

import (
    "database/sql"
    "bankapp/internal/models"
)

type CardRepo struct {
    DB *sql.DB
}

func NewCardRepo(db *sql.DB) *CardRepo {
    return &CardRepo{DB: db}
}

func (r *CardRepo) Create(card *models.Card) error {
    query := `INSERT INTO cards (account_id, number, expiry, cvv_hash, hmac) 
              VALUES ($1, $2, $3, $4, $5)`
    _, err := r.DB.Exec(query, card.AccountID, card.Number, card.Expiry, card.CVVHash, card.HMAC)
    return err
}
