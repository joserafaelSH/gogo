package config

import "fmt"

const (
	HostBb1     = "localhost"
	PortBb1     = 5434
	UserBb1     = "postgres"
	PasswordBb1 = "root"
	DbnameBb1   = "pg1"
)
const (
	HostBb2     = "localhost"
	PortBb2     = 5433
	UserBb2     = "postgres"
	PasswordBb2 = "root"
	DbnameBb2   = "pg2"
)

func CreteUrlConnectionDb1() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		HostBb1, PortBb1, UserBb1, PasswordBb1, DbnameBb1)
}

func CreteUrlConnectionDb2() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		HostBb2, PortBb2, UserBb2, PasswordBb2, DbnameBb2)
}
