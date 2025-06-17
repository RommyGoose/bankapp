package service

import (
    "bankapp/internal/repository"
)

type AccountService struct {
    Repo *repository.AccountRepo
}

func NewAccountService(repo *repository.AccountRepo) *AccountService {
    return &AccountService{Repo: repo}
}

func (s *AccountService) CreateAccount(userID int) (int, error) {
    return s.Repo.Create(userID)
}
