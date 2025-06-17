package models

type Card struct {
    ID        int    `json:"id"`
    AccountID int    `json:"account_id"`
    Number    string `json:"number"`
    Expiry    string `json:"expiry"`
    CVVHash   string `json:"-"`
    HMAC      string `json:"-"`
}
