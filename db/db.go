package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectaComBancoDeDados() *sql.DB {
	conexao := "host=atvosfunc.postgresql.dbaas.com.br port=5432 dbname=atvosfunc user=atvosfunc password=Atv@s2024 sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db

}
