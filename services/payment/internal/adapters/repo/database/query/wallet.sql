-- name: CreateWallet :exec
INSERT INTO payment.wallets(user_id)
VALUES(@user_id)
ON CONFLICT DO NOTHING;

-- name: GetWallet :one
SELECT amount FROM payment.wallets WHERE user_id = @user_id;

-- name: DepositWallet :exec
UPDATE payment.wallets
SET amount = amount + @deposit_amount
WHERE user_id = @user_id;

-- name: WithdrawWallet :exec
UPDATE payment.wallets
SET amount = amount - @withdraw_amount
WHERE user_id = @user_id;
