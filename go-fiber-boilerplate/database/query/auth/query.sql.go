package auth

const (
	CheckUserStatusSQL = `select id, is_active, is_suspended, is_allowed_cod from users where id = $1`
)
