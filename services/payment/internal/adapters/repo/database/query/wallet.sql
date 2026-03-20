-- name: CreateWallet :exec
INSERT INTO payment.wallets(user_id)
VALUES(@user_id)
ON CONFLICT DO NOTHING;

-- name: GetWalletByUser :one
SELECT * FROM payment.wallets WHERE user_id = @user_id;

-- name: UpdateWallet :exec
UPDATE payment.wallets
SET balance = balance + @amount
WHERE id = @id;
