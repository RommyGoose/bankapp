package utils

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
    "golang.org/x/crypto/bcrypt"
    "math/rand"
    "strconv"
)

func HashCVV(cvv string) (string, error) {
    return bcrypt.GenerateFromPassword([]byte(cvv), bcrypt.DefaultCost)
}

func ComputeHMAC(data string, secret []byte) string {
    h := hmac.New(sha256.New, secret)
    h.Write([]byte(data))
    return hex.EncodeToString(h.Sum(nil))
}

func GenerateCardNumber() string {
    prefix := "400000"
    number := prefix
    for len(number) < 15 {
        number += strconv.Itoa(rand.Intn(10))
    }
    return number + calculateLuhnChecksum(number)
}

func calculateLuhnChecksum(number string) string {
    sum := 0
    alt := false
    for i := len(number) - 1; i >= 0; i-- {
        n, _ := strconv.Atoi(string(number[i]))
        if alt {
            n *= 2
            if n > 9 {
                n -= 9
            }
        }
        sum += n
        alt = !alt
    }
    return strconv.Itoa((10 - sum%10) % 10)
}
