package core

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql" // Asegúrate de importar el controlador MySQL
)

func ConnectToDB() (*sql.DB, error) {
	// Cargar variables de entorno desde .env
	if err := godotenv.Load(); err != nil {
		log.Println("No se pudieron cargar las variables de entorno")
	}

	// Leer variables de entorno
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	nameDB := os.Getenv("DB_DATABASE")

	// Construir el DSN con el formato correcto
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, nameDB)

	// Crear conexión
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("Error al abrir la conexión: %v", err)
		return nil, fmt.Errorf("error al abrir la conexión: %v", err)
	}

	// Validar la conexión con Ping
	if err := db.Ping(); err != nil {
		log.Printf("Error al validar la conexión con la base de datos: %v", err)
		return nil, fmt.Errorf("error al validar la conexión: %v", err)
	}

	// Configurar el pool de conexiones
	db.SetMaxOpenConns(25)             // Máximo número de conexiones abiertas
	db.SetMaxIdleConns(10)             // Máximo número de conexiones inactivas
	db.SetConnMaxIdleTime(time.Minute) // Tiempo máximo de inactividad
	db.SetConnMaxLifetime(time.Minute * 5) // Vida máxima de una conexión

	fmt.Println("Conexión exitosa a la base de datos")
	return db, nil
}
