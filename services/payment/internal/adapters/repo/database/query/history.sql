-- name: CreateHistory :exec
INSERT INTO payment.histories(wallet_id, amount, description)
VALUES(@wallet_id, @amount, @description);

-- name: GetHistoriesByWallet :many
SELECT * FROM payment.histories WHERE wallet_id = @wallet_id;
