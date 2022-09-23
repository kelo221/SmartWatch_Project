package main

import (
	dbModels "SmartWatch_Project/db/models"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"log"
)

//go:generate sqlboiler --wipe psql
func main() {
	db := connectDB()
	boil.SetDB(db)

	fmt.Println("DB connected.")

	user := &dbModels.User{
		Username: "1234",
		Password: "1234",
	}

	err := user.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		fmt.Println("cant insert")
	}

	fmt.Println(user.ID)

	userQ, err := dbModels.Users(dbModels.UserWhere.Username.EQ("1234")).OneG(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(userQ)
}

func connectDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:1234@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
