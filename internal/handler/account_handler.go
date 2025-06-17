package handler

import (
    "encoding/json"
    "net/http"
    "strconv"
    "database/sql"

    "bankapp/internal/repository"
    "bankapp/internal/service"
)

var accountService *service.AccountService

func InitAccountHandler(db *sql.DB) {
    repo := repository.NewAccountRepo(db)
    accountService = service.NewAccountService(repo)
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
    userIDStr := r.Context().Value("userID").(string)
    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    accountID, err := accountService.CreateAccount(userID)
    if err != nil {
        http.Error(w, "Failed to create account", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]interface{}{
        "account_id": accountID,
    })
}
