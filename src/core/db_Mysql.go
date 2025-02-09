package core

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Conn_Mysql struct {
	DB  *sql.DB
	Err string
} // para mandar mensaje de las filas afectadas

func GetDBPool() *Conn_Mysql {
	error := ""

	err := godotenv.Load() //para obtener los datos de las credenciales en la variable de entorno
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	// Esto es para obtener las variables
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbSchema := os.Getenv("DB_SCHEMA")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPass, dbHost, dbSchema)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		error = fmt.Sprintf("%v", err)
		return &Conn_Mysql{DB: db, Err: error}
	}
	// Configuración del pool de conexiones
	db.SetMaxOpenConns(10)

	// Probar la conexión
	if err := db.Ping(); err != nil {
		db.Close()
		error = fmt.Sprintf("error al verificar la conexión a la base de datos: %w", err)
	}

	return &Conn_Mysql{DB: db, Err: error}
}

func (conn *Conn_Mysql) ExecutePreparedQuery(query string, values ...interface{}) (sql.Result, error) {
	stmt, err := conn.DB.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("error al preparar la consulta: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(values...)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta preparada: %w", err)
	}

	return result, nil
}

func (conn *Conn_Mysql) FetchRows(query string, values ...interface{}) (*sql.Rows, *sql.Rows) {
	rows, err := conn.DB.Query(query, values...)
	if err != nil {
		fmt.Printf("error al ejecutar la consulta SELECT: %w", err)
	}

	return rows, nil

}
