CREATE TABLE
    "accounts" (
        id BIGSERIAL PRIMARY KEY,
        user_id INTEGER NOT NULL,
        balance DOUBLE PRECISION NOT NULL DEFAULT 0.00,
        currency VARCHAR(10) NOT NULL,
        created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (user_id) REFERENCES users(id)
    );

ALTER TABLE "accounts"
ADD
    CONSTRAINT "unique_user_currency" UNIQUE (user_id, currency);

CREATE TABLE
    "entries" (
        id BIGSERIAL PRIMARY KEY,
        account_id INTEGER NOT NULL,
        amount DOUBLE PRECISION NOT NULL,
        created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (account_id) REFERENCES accounts(id)
    );

CREATE TABLE
    "transfers" (
        id BIGSERIAL PRIMARY KEY,
        from_account_id INTEGER NOT NULL,
        to_account_id INTEGER NOT NULL,
        amount DOUBLE PRECISION NOT NULL,
        created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (from_account_id) REFERENCES accounts(id),
        FOREIGN KEY (to_account_id) REFERENCES accounts(id)
    );