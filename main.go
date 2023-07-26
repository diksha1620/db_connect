// package main

// import (
// 	"github.com/diksha1620/go-gorm-gin/routes"
// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	router := gin.Default()
// 	routes.UserRouter(router)

// 	router.Run()
// }

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Dishu@2026"
	dbname   = "test"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	insertStmt := `INSERT INTO "Employee" ("Name", "EmpId") VALUES ('Rohit', 21)`
	_, e := db.Exec(insertStmt)
	CheckError(e)

	insertDynStmt := `INSERT INTO "Employee" ("Name", "EmpId") VALUES ($1, $2)`
	_, e = db.Exec(insertDynStmt, "Krish", 3) // Note that EmpId is provided as an integer (3) and not a string ("03")
	CheckError(e)

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
