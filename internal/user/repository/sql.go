package repository

const (
	queryAddUser = `INSERT INTO users (user_id, name, last_name, password, email,
	role, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
	RETURNING user_id, name, last_name, password, email, role, created_at, updated_at`

	queryByID = `select user_id, name, last_name, password, email, role, created_at,
	updated_at from users where user_id=$1`

	queryFindByEmail = `select user_id, name, last_name, password, email, role,
	created_at,updated_at from users where email=$1`

	queryUpdate = `update "users" set "name" = COALESCE($1, "name"), "last_name" = COALESCE($2, "last_name"),
	"password" = COALESCE($3, "password"), "email" = COALESCE($4, "email"),
	"updated_at" = COALESCE($5, "updated_at") where user_id=$6`
)
