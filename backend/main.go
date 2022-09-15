package main

import (
	"SmartWatch_Project/ent"
	"SmartWatch_Project/ent/event"
	"SmartWatch_Project/ent/user"
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetUsername("user").
		SetPassword("password").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.Username("user")).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

func CreateEvents(ctx context.Context, client *ent.Client) (*ent.User, error) {
	// Create a new car with model "Tesla".
	tesla, err := client.Event.
		Create().
		SetEventName("test").
		SetIsSilent(true).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %w", err)
	}
	log.Println("car was created: ", tesla)

	// Create a new user, and add it the 2 cars.
	a8m, err := client.User.
		Create().
		SetUsername("user").
		SetPassword("password").
		AddEvents(tesla).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", a8m)
	return a8m, nil
}

func QueryEvents(ctx context.Context, a8m *ent.User) error {
	cars, err := a8m.QueryEvents().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %w", err)
	}
	log.Println("returned cars:", cars)

	// What about filtering specific cars.
	ford, err := a8m.QueryEvents().
		Where(event.EventName("test")).
		Only(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %w", err)
	}
	log.Println(ford)
	return nil
}

func main() {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=Smartwatch password=1234  sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {

		}
	}(client)
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	ctx := context.Background()

	_, err = CreateEvents(ctx, client)
	if err != nil {
		return
	}
}
