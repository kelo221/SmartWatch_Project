package dbHandler

import (
	dbModels "SmartWatch_Project/db/models"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"log"
)

var db *sql.DB

func ConnectToDB() {

	db = connectDB()
	boil.SetDB(db)

	fmt.Println("DB connected.")
	/*
		user := &dbModels.User{
			Username: "al2t",
			Password: "a2lt",
		}

		err := user.Insert(context.Background(), db, boil.Infer())
		if err != nil {
			fmt.Println("cant insert")
		}

		fmt.Println(user.ID)

		userQ, err := dbModels.Users(dbModels.UserWhere.Username.EQ("al2t")).OneG(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(userQ.Username)
		createEvent(context.Background(), *userQ)*/
}

func connectDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:1234@localhost:5432/smartwatch_project?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func CreateEvent(ctx context.Context, newEvent dbModels.Event) string {

	fmt.Println(newEvent)
	err := newEvent.InsertG(ctx, boil.Infer())
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	return ""
}

func CreateUser(user dbModels.User) string {

	err := user.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	return ""
}
