package models

type Transaction struct {
    ID            int     `json:"id"`
    FromAccountID int     `json:"from_account_id"`
    ToAccountID   int     `json:"to_account_id"`
    Amount        float64 `json:"amount"`
    Timestamp     string  `json:"timestamp"`
}
