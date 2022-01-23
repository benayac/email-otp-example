package database

var (
	INSERT_USER           = "INSERT INTO account_data(name, email, password, address, is_verified) VALUES ($1, $2, $3, $4, $5);"
	UPDATE_VERIFY_USER    = "UPDATE account_data SET is_verified = $1 WHERE email = $2"
	SELECT_PASSWORD_USER  = "SELECT password FROM account_data WHERE email = $1"
	SELECT_PASSWORD_ADMIN = "SELECT password FROM account_admin WHERE email = $1"
)
