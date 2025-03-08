package dependencies

import (
	"database/sql"
	"fmt"
	"hex_sub/src/core"
	"hex_sub/src/payments/application"
	"hex_sub/src/payments/application/services"
	"hex_sub/src/payments/infraestructure"
	"hex_sub/src/payments/infraestructure/adapters"
	"hex_sub/src/payments/infraestructure/controllers"
)

var (
	mySQL infraestructure.MySQL
	db    *sql.DB
	rabbitmqAdapter *adapters.RabbitMQAdapter
)

func Init() {
	db, err := core.ConnectToDB()

	if err != nil {
		fmt.Println("server error")
		return
	}

	mySQL = *infraestructure.NewMySQL(db)
	rabbitmqAdapter, err = adapters.NewRabbitMQAdapter()
	if err != nil {
		fmt.Println("Error iniciando RabbitMQ:", err)
		return
	}

}
func CloseDB() {
	if db != nil {
		db.Close()
		fmt.Println("Conexión a la base de datos cerrada.")
	}
}

func GetCreatePaymentController()*controllers.CreatePaymentController{
	serviceNotification := services.NewServiceNotification(rabbitmqAdapter)
	caseCreatePayment := application.NewCreatePayment(&mySQL,serviceNotification)
	return controllers.NewCreatePaymentController(caseCreatePayment,serviceNotification)
}
