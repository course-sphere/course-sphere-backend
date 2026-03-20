-- name: CreateWallet :exec
INSERT INTO payment.wallets(user_id)
VALUES(@user_id)
ON CONFLICT DO NOTHING;

-- name: GetWallet :one
SELECT * FROM payment.wallets WHERE user_id = @user_id;

-- name: UpdateWalletBalance :exec
UPDATE payment.wallets
SET balance = balance + @amount
WHERE id = @id;
