package service

import (
    "fmt"
    "time"
    "math"
    "bankapp/internal/models"
    "bankapp/internal/repository"
)

type CreditService struct {
    Repo *repository.CreditRepo
}

func NewCreditService(repo *repository.CreditRepo) *CreditService {
    return &CreditService{Repo: repo}
}

func (s *CreditService) ApplyCredit(accountID int, amount float64, rate float64, months int) error {
    credit := models.Credit{
        AccountID:    accountID,
        Amount:       amount,
        InterestRate: rate,
        TermMonths:   months,
    }
    creditID, err := s.Repo.Create(credit)
    if err != nil {
        return err
    }

    monthlyRate := rate / 12 / 100
    payment := amount * (monthlyRate * math.Pow(1+monthlyRate, float64(months))) / (math.Pow(1+monthlyRate, float64(months)) - 1)

    var schedule []models.PaymentSchedule
    due := time.Now().AddDate(0, 1, 0)

    for i := 0; i < months; i++ {
        schedule = append(schedule, models.PaymentSchedule{
            CreditID: creditID,
            Amount:   math.Round(payment*100) / 100,
            DueDate:  due.Format("2006-01-02"),
        })
        due = due.AddDate(0, 1, 0)
    }

    return s.Repo.CreateSchedule(creditID, schedule)
}
