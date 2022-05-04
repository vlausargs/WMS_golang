-- name: CreateWHSE :one
INSERT INTO "WHSE"
(whseid, "desc", create_date)
VALUES($1,$2, now())
RETURNING *;

-- name: GetWHSE :one
SELECT * FROM "WHSE" where whseid = $1;

-- name: GetAllWHSE :many
SELECT * FROM "WHSE" ;