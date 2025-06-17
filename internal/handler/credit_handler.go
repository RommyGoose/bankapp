package handler

import (
    "encoding/json"
    "net/http"
    "strconv"
    "database/sql"

    "bankapp/internal/service"
    "bankapp/internal/repository"
)

var creditService *service.CreditService

func InitCreditHandler(db *sql.DB) {
    repo := repository.NewCreditRepo(db)
    creditService = service.NewCreditService(repo)
}

func ApplyCredit(w http.ResponseWriter, r *http.Request) {
    var req struct {
        AccountID int     `json:"account_id"`
        Amount    float64 `json:"amount"`
        Rate      float64 `json:"interest_rate"`
        Term      int     `json:"term_months"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    err := creditService.ApplyCredit(req.AccountID, req.Amount, req.Rate, req.Term)
    if err != nil {
        http.Error(w, "Failed to apply credit: "+err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"message": "credit approved"})
}
