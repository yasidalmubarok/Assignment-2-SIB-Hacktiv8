package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "yasidmubarok"
	password = "postgres"
	dialect  = "postgres"
	dbname   = "order-pg"
)

var (
	db  *sql.DB
	err error
)

func handleDatabaseConnection() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open(dialect, psqlInfo)

	if err != nil {
		log.Panic("error occured while validating database arguments:", err.Error())
	}

	err = db.Ping()

	if err != nil {
		log.Panic("error occured while opening a connection to database:", err.Error())
	}
}

func handleCreateRequiredTables() {
	orderTable := `
		CREATE TABLE IF NOT EXISTS "orders" (
			order_id SERIAL PRIMARY KEY,
			customer_name VARCHAR(255) NOT NULL,
			ordered_at timestamptz DEFAULT now(),
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now()
		);
	`

	itemTable := `
		CREATE TABLE IF NOT EXISTS "items" (
			item_id SERIAL PRIMARY KEY,
			item_code VARCHAR(191) NOT NULL,
			quantity INT NOT NULL,
			description TEXT NOT NULL,	
			order_id INT NOT NULL,
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now(),
			CONSTRAINT item_order_id_fk
			FOREIGN KEY(order_id)
				REFERENCES orders(order_id)
					ON DELETE CASCADE
		);
	`

	_, err := db.Exec(orderTable)

	if err != nil {
		log.Panic("error occured while create order table:", err.Error())
	}

	_, err = db.Exec(itemTable)

	if err != nil {
		log.Panic("error occured while create item table:", err.Error())
	}
}

func InitiliazeDatabase() {
	handleDatabaseConnection()
	handleCreateRequiredTables()
}

func GetDatabaseInstance() *sql.DB {
	if db == nil {
		log.Panic("DB is still nill!!")
	}

	return db
}
