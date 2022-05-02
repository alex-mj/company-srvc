package repository

type PostgresDB struct {
}

func NewPostgresStorage() *PostgresDB {
	return &PostgresDB{}
}
