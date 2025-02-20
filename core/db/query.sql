-- name: GetScans :one
SELECT * FROM scans
WHERE id = ? LIMIT 1;

-- name: Count :one
SELECT COUNT(*) FROM scans;

