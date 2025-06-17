package service

import (
    "fmt"
    "time"
    "math/rand"
    "bankapp/internal/models"
    "bankapp/internal/repository"
    "bankapp/internal/utils"
)

type CardService struct {
    Repo *repository.CardRepo
}

func NewCardService(repo *repository.CardRepo) *CardService {
    return &CardService{Repo: repo}
}

func (s *CardService) GenerateCard(accountID int, cvv string) error {
    number := utils.GenerateCardNumber()
    expiry := time.Now().AddDate(3, 0, 0).Format("01/06")
    cvvHash, _ := utils.HashCVV(cvv)
    hmac := utils.ComputeHMAC(number, []byte("hmac_secret"))

    card := &models.Card{
        AccountID: accountID,
        Number:    number,
        Expiry:    expiry,
        CVVHash:   cvvHash,
        HMAC:      hmac,
    }

    return s.Repo.Create(card)
}
