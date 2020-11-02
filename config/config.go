package config

import (
	"database/sql"
	"fmt"
	_ "github.com/nakagami/firebirdsql"
	"log"
)

var DB *sql.DB

// Parametrs  создание подключения к БД
var Parametrs Config

// SetParametrs установка параметров приложения
func SetParametrs() {
	db, err := sql.Open("firebirdsql", "sysdba:serg@localhost:3050/C:/obmen/test.FDB")
	DB = db
	if err != nil {
		fmt.Printf("error : %v", err)
		log.Fatal(err)
	}

}
