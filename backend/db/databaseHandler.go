package dbHandler

import (
	dbModels "SmartWatch_Project/db/models"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"log"
	"time"
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

	user, err := dbModels.Users(dbModels.UserWhere.ID.EQ(ID)).OneG(ctx)
	if err != nil {
		return err.Error(), nil
	}

	events, err := user.UseridEvents().AllG(ctx)
	if err != nil {
		return err.Error(), nil
	}

	return "", events
}

func DeleteEventByID(userID int, eventID int, ctx context.Context) string {

	queryResultCount, err := dbModels.Events(qm.Where("id=? and userID=?", eventID, userID)).DeleteAll(ctx, db)

	if queryResultCount == 0 {
		return "Deletion failed!"
	}

	if err != nil {
		return err.Error()
	}

	return ""
}

func UpdateEventTime(userID int, eventID int, newTime time.Time, ctx context.Context) string {

	queryResultCount, err := dbModels.Events(qm.Where("id=? and userID=?", eventID, userID)).UpdateAll(ctx, db, dbModels.M{"eventtime": newTime})
	if queryResultCount == 0 {
		return "Update failed!"
	}
	if err != nil {
		return err.Error()
	}

	return ""
}
