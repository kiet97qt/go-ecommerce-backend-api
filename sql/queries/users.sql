-- Basic queries for sqlc to generate models from the users/roles tables

-- name: GetUserByID :one
SELECT
  id,
  username,
  is_active,
  created_at,
  updated_at
FROM users
WHERE id = ? LIMIT 1;

-- name: ListUsers :many
SELECT
  id,
  username,
  is_active,
  created_at,
  updated_at
FROM users
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: ListUserRoles :many
SELECT
  ur.user_id,
  ur.role_id
FROM user_roles ur
WHERE ur.user_id = ?;


