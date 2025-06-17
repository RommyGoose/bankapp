package scheduler

import (
    "database/sql"
    "log"
    "time"
)

func StartPaymentScheduler(db *sql.DB) {
    ticker := time.NewTicker(12 * time.Hour)
    go func() {
        for {
            <-ticker.C
            err := ProcessPayments(db)
            if err != nil {
                log.Println("Payment scheduler error:", err)
            }
        }
    }()
}

func ProcessPayments(db *sql.DB) error {
    tx, err := db.Begin()
    if err != nil {
        return err
    }

    // Найти все просроченные и неоплаченные платежи
    rows, err := tx.Query(`
        SELECT ps.id, ps.credit_id, ps.amount, c.account_id
        FROM payment_schedules ps
        JOIN credits c ON c.id = ps.credit_id
        WHERE ps.paid = false AND ps.due_date <= CURRENT_DATE
    `)
    if err != nil {
        tx.Rollback()
        return err
    }
    defer rows.Close()

    for rows.Next() {
        var scheduleID, creditID, accountID int
        var amount float64

        if err := rows.Scan(&scheduleID, &creditID, &amount, &accountID); err != nil {
            tx.Rollback()
            return err
        }

        // Попробовать списать со счёта
        result, err := tx.Exec(`UPDATE accounts SET balance = balance - $1 WHERE id = $2 AND balance >= $1`, amount, accountID)
        if err != nil {
            tx.Rollback()
            return err
        }

        rowsAffected, _ := result.RowsAffected()
        if rowsAffected == 1 {
            // Обновить статус платежа
            _, err = tx.Exec(`UPDATE payment_schedules SET paid = true WHERE id = $1`, scheduleID)
            if err != nil {
                tx.Rollback()
                return err
            }
        } else {
            // Недостаточно средств — начислить штраф 10%
            penalty := amount * 0.10
            _, err = tx.Exec(`UPDATE payment_schedules SET amount = amount + $1 WHERE id = $2`, penalty, scheduleID)
            if err != nil {
                tx.Rollback()
                return err
            }
        }
    }

    return tx.Commit()
}
