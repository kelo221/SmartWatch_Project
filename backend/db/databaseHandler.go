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

func FetchMultipleByID(ID int, ctx context.Context) (string, dbModels.EventSlice) {

	author, err := dbModels.Users(dbModels.UserWhere.ID.EQ(ID)).OneG(ctx)
	if err != nil {
		return err.Error(), nil
	}

	events, err := author.EventidEvents().AllG(ctx)
	if err != nil {
		return err.Error(), nil
	}

	return "", events
}
