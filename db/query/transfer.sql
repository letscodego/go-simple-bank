-- name: CreateTransfer :execresult
INSERT INTO transfers (
  from_account_id, 
  to_account_id,
  amount
) VALUES (
  ?, ?, ? );

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = ? LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM transfers
ORDER BY id
LIMIT ?, ?;

-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id = ?;

-- name: UpdateTransfer :execresult
UPDATE transfers 
SET amount = ?
WHERE id = ?;