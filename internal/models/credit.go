package models

type Credit struct {
    ID           int     `json:"id"`
    AccountID    int     `json:"account_id"`
    Amount       float64 `json:"amount"`
    InterestRate float64 `json:"interest_rate"`
    TermMonths   int     `json:"term_months"`
}
