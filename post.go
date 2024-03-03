package post

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type UserData struct {
	ID          int
	Username    string
	Name        string
	Surname     string
	Description string
}

var (
	Hostname = ""
	Port     = 5432
	Username = ""
	Password = ""
	Database = ""
)

func openConnection() (*sql.DB, error) {
	connenction_string := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Hostname, Port, Username, Password, Database)
	db, err := sql.Open("postgres", connenction_string)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func AddUser(d UserData) int {
	db, err := openConnection()
	if err != nil {
		fmt.Println(err)
		return -1
	}
	defer db.Close()

	insertStatement := `insert into "users" ("username") values ($1)`
	_, err = db.Exec(insertStatement, d.Username)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return 1
}
