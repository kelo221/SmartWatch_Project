// Package dbHandler takes care of communicating with PostgresSQL
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

// ConnectToDB Establishes connection to a db variable that is used for database related requests with SQLBoiler.
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

// NewEventPopUp Pushes an event to the database.
func NewEventPopUp(ctx context.Context, newEvent dbModels.Event) string {

	err := newEvent.InsertG(ctx, boil.Infer())
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	return ""
}

// CreateUser Pushes an user to the database.
func CreateUser(user dbModels.User) string {

	err := user.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	return ""
}

// FetchMultipleByID Fetches all events that match the user id, return a slice of events.
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

// DeleteEventByDate Deletes an event that match the user id and the given id.
func DeleteEventByDate(userID int, eventDate string, ctx context.Context) string {

	queryResultCount, err := dbModels.Events(qm.Where("eventtime=? and userID=?", eventDate, userID)).DeleteAll(ctx, db)
	if queryResultCount == 0 {
		return "Deletion failed!"
	}
	if err != nil {
		return err.Error()
	}
	return ""
}

// DeleteEventByID Deletes an event that match the user id and the given id.
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

// UpdateEventTime Updates an event with a new time.
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

// UpdateEvent Updates an event.
func UpdateEvent(userID int, eventDate string, newEvent dbModels.Event, ctx context.Context) string {

	queryResultCount, err := dbModels.Events(qm.Where("eventtime=? and userID=?", eventDate, userID)).UpdateAll(ctx, db,
		dbModels.M{
			"eventtime":      newEvent.Eventtime,
			"eventname":      newEvent.Eventname,
			"issilent":       newEvent.Issilent,
			"snoozedisabled": newEvent.Snoozedisabled,
		})
	if queryResultCount == 0 {
		return "Update failed!"
	}
	if err != nil {
		return err.Error()
	}
	return ""

}
