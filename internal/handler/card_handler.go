package handler

import (
    "encoding/json"
    "net/http"
    "strconv"

    "bankapp/internal/service"
    "bankapp/internal/repository"
    "database/sql"
)

var cardService *service.CardService

func InitCardHandler(db *sql.DB) {
    repo := repository.NewCardRepo(db)
    cardService = service.NewCardService(repo)
}

func CreateCard(w http.ResponseWriter, r *http.Request) {
    var req struct {
        AccountID int    `json:"account_id"`
        CVV       string `json:"cvv"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    err := cardService.GenerateCard(req.AccountID, req.CVV)
    if err != nil {
        http.Error(w, "Failed to create card", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"message": "card created"})
}
