package postgreSQL

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
)

type Params struct {
	Psql struct {
		Login              string `yaml:"login"`
		Password           string `yaml:"password"`
		SslMode            string `yaml:"ssl_mode"`
		DriverName         string `yaml:"driver_name"`
		DatabaseName       string `yaml:"database_name"`
		AttemptsConnection int    `yaml:"attempts_connection"`
	}
}

type PostgreSQL struct {
}

func open() (*sql.DB, error) {
	paramsFile, err := os.ReadFile(PARAMS_PATH)
	checkError(err)
	var params Params
	err = yaml.Unmarshal(paramsFile, &params)
	checkError(err)

	connStr := "user=" + params.Psql.Login + " password=" + params.Psql.Password + " dbname=" + params.Psql.DatabaseName + " sslmode=" + params.Psql.SslMode
	var db *sql.DB
	//на случай, если не удалсь подключиться - пробуем несколько раз
	for i := 0; i <= params.Psql.AttemptsConnection; i++ {
		db, err = sql.Open(params.Psql.DriverName, connStr)
		if err != nil {
			log.Println("Unsuccessful attempt to connect to DB: ", err)
			time.Sleep(1000)
			log.Println("Next try...")
			continue
		} else {
			return db, nil
		}
	}
	return nil, errors.New("Failed to connect to DB!")
}

func (p *PostgreSQL) SendQuery(query string) (sql.Result, error) {
	db, err := open()
	checkError(err)

	defer closeDB(db)

	res, err := db.Exec(query)
	return res, err
}
func (p *PostgreSQL) GetRow(query string) sql.Row {
	db, err := open()
	checkError(err)
	defer closeDB(db)
	res := db.QueryRow(query)
	return *res
}
func (p *PostgreSQL) GetRows(query string) *sql.Rows {
	db, err := open()
	checkError(err)
	defer closeDB(db)
	res, err1 := db.Query(query)
	checkError(err1)
	return res
}

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func closeDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		//TODO
	}
}
