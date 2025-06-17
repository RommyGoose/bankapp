

CREATE TABLE credits (
    id SERIAL PRIMARY KEY,
    account_id INTEGER REFERENCES accounts(id),
    amount NUMERIC(12,2),
    interest_rate NUMERIC(5,2),
    term_months INTEGER
);

CREATE TABLE payment_schedules (
    id SERIAL PRIMARY KEY,
    credit_id INTEGER REFERENCES credits(id),
    amount NUMERIC(12,2),
    due_date DATE,
    paid BOOLEAN DEFAULT false
);
