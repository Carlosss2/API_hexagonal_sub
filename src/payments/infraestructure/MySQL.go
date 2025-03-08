package infraestructure

import (
	"database/sql"
	"fmt"
	

	
)

type MySQL struct {
	DB *sql.DB
}

func NewMySQL(db *sql.DB)*MySQL{
	return &MySQL{DB: db}
}

func (mysql *MySQL) Save(menssage string) error {
	_, err := mysql.DB.Exec("INSERT INTO payments (message) VALUES (?)", menssage)

	if err != nil {
		return fmt.Errorf("[MySQL] Error a save the payment : %w", err)
	}
	return nil
}