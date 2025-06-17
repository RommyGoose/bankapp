package models

type PaymentSchedule struct {
    ID        int     `json:"id"`
    CreditID  int     `json:"credit_id"`
    Amount    float64 `json:"amount"`
    DueDate   string  `json:"due_date"`
    Paid      bool    `json:"paid"`
}
