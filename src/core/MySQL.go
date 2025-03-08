package core

import (
	"database/sql"
	"fmt"
	"log"
	

	_ "github.com/go-sql-driver/mysql" // Asegúrate de importar el controlador MySQL
)

func ConnectToDB() (*sql.DB, error) {
	// Leer variables de entorno
	user := "root"
	password := "Castro2005"
	host := "localhost"
	port := "3306" // Asegúrate de definir este valor en el archivo .env
	nameDB := "y"

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
	db.SetMaxOpenConns(25)         // Máximo número de conexiones abiertas
	db.SetMaxIdleConns(10)         // Máximo número de conexiones inactivas
	db.SetConnMaxIdleTime(10)      // Tiempo máximo de inactividad en segundos
	db.SetConnMaxLifetime(5 * 60)  // Vida máxima de una conexión en segundos

	fmt.Println("Conexión exitosa a la base de datos")
	return db, nil
}
