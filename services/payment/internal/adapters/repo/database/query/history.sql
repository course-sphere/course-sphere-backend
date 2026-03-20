-- name: CreateHistory :exec
INSERT INTO payment.histories(wallet_id, amount, detail)
VALUES(@wallet_id, @amount, @detail);

-- name: GetHistoriesByWallet :many
SELECT * FROM payment.histories WHERE wallet_id = @wallet_id;
