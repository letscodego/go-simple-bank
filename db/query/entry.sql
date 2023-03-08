-- name: CreateEntry :execresult
INSERT INTO entries (
  account_id, 
  amount
) VALUES (
  ?, ? );

-- name: GetEntry :one
SELECT * FROM entries
WHERE id = ? LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entries
WHERE account_id = ?
ORDER BY id
LIMIT ?, ?;

-- name: DeleteEntry :exec
DELETE FROM entries
WHERE id = ?;

-- name: UpdateEntry :exec
UPDATE entries 
SET amount = ?
WHERE id = ?;