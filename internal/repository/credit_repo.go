package repository

import (
    "database/sql"
    "bankapp/internal/models"
)

type CreditRepo struct {
    DB *sql.DB
}

func NewCreditRepo(db *sql.DB) *CreditRepo {
    return &CreditRepo{DB: db}
}

func (r *CreditRepo) Create(credit models.Credit) (int, error) {
    var id int
    err := r.DB.QueryRow(`INSERT INTO credits (account_id, amount, interest_rate, term_months) 
        VALUES ($1, $2, $3, $4) RETURNING id`,
        credit.AccountID, credit.Amount, credit.InterestRate, credit.TermMonths).Scan(&id)
    return id, err
}

func (r *CreditRepo) CreateSchedule(creditID int, payments []models.PaymentSchedule) error {
    tx, err := r.DB.Begin()
    if err != nil {
        return err
    }
    for _, p := range payments {
        _, err := tx.Exec(`INSERT INTO payment_schedules (credit_id, amount, due_date, paid) 
            VALUES ($1, $2, $3, false)`, creditID, p.Amount, p.DueDate)
        if err != nil {
            tx.Rollback()
            return err
        }
    }
    return tx.Commit()
}
