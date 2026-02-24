package methods

import (
	"github.com/jackc/pgx/v5"
)

type AuthMethods struct {
	conn *pgx.Conn
}

func NewAuthDB(conn *pgx.Conn) *AuthMethods {
	return &AuthMethods{
		conn: conn,
	}
}
