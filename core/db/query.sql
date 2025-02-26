-- name: GetScans :one
SELECT * FROM scans
WHERE id = ? LIMIT 1;

-- name: CountScans :one
SELECT COUNT(*) FROM scans;

