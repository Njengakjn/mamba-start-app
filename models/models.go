package models

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
	"github.com/upper/db/v4/adapter/postgresql"
)

var pool *sql.DB
var upper db.Session

type Models struct {
}

func New(dbPool *sql.DB) Models {
	pool = dbPool

	if os.Getenv("DATABASE_TYPE") == "mysql" || os.Getenv("DATABASE_TYPE") == "mariadb" {
		upper, _ = mysql.New(pool)
	} else {
		upper, _ = postgresql.New(pool)
	}

	return Models{}
}

func getInsertID(i db.ID) int {
	id := fmt.Sprintf("%T", i)
	if id == "int64" {
		return int(i.(int64))
	}

	return i.(int)
}
