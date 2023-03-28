-- name: CreateAccount :execresult
INSERT INTO accounts (
  owner, 
  balance,
  currency
) VALUES (
  ?, ?, ? );

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = ? LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT * FROM accounts
WHERE id = ? LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
WHERE owner = ?
ORDER BY id
LIMIT ?, ?;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = ?;

-- name: UpdateAccount :exec
UPDATE accounts 
SET balance = ?
WHERE id = ?;

-- name: AddAccountBalance :exec
UPDATE accounts 
SET balance = balance + sqlc.arg(amount)
WHERE id = sqlc.arg(id);