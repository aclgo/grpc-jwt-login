package repository

const (
	queryAddUser = `INSERT INTO users (id, name, last_name, password, email,
	role, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
	RETURNING id, name, last_name, password, email, role, created_at, updated_at`

	queryByID = `select id, name, last_name, password, email, role, created_at,
	updated_at from users where id=$1`

	queryFindByEmail = `select id, name, last_name, password, email, role,
	created_at,updated_at from users where email=$1`

	queryUpdate = `update users set
	id, name, last_name, password, email, role, created_at,
	updated_at from users where id=$1`
)
