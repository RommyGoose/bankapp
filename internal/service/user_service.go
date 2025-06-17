package service

import (
    "database/sql"
    "errors"
    "time"

    "golang.org/x/crypto/bcrypt"
    "github.com/golang-jwt/jwt/v5"
    "bankapp/internal/models"
    "bankapp/internal/config"
)

type UserService struct {
    DB *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
    return &UserService{DB: db}
}

func (s *UserService) Register(user models.User) error {
    var exists bool
    err := s.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email=$1 OR username=$2)", user.Email, user.Username).Scan(&exists)
    if err != nil {
        return err
    }
    if exists {
        return errors.New("email or username already exists")
    }

    hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    _, err = s.DB.Exec("INSERT INTO users (email, username, password_hash) VALUES ($1, $2, $3)",
        user.Email, user.Username, string(hash))
    return err
}

func (s *UserService) Login(email, password string) (string, error) {
    var id int
    var passwordHash string
    err := s.DB.QueryRow("SELECT id, password_hash FROM users WHERE email=$1", email).Scan(&id, &passwordHash)
    if err != nil {
        return "", errors.New("invalid email or password")
    }

    err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
    if err != nil {
        return "", errors.New("invalid email or password")
    }

    claims := jwt.RegisteredClaims{
        Subject:   string(rune(id)),
        ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
        IssuedAt:  jwt.NewNumericDate(time.Now()),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signedToken, err := token.SignedString([]byte(config.JwtSecret))
    if err != nil {
        return "", err
    }

    return signedToken, nil
}
