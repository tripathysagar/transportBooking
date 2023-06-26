package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/tripathysagar/transport-booking/db"
	"github.com/tripathysagar/transport-booking/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DBNAME = "test_taxi_reservation"

type testdb struct {
	db.UserStore
}

func (tdb *testdb) teardown(t *testing.T) {
	if err := tdb.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}

}

func setup(t *testing.T) *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	return &testdb{
		UserStore: db.NewMongoStore(client, DBNAME),
	}
}

func TestPostUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	app := fiber.New()
	userHandler := NewUserHandler(tdb.UserStore)
	app.Post("/", userHandler.HandlePostUser)

	params := types.CreateUserParams{
		Email:     "foo@bar.com",
		FirstName: "foo",
		LastName:  "bar",
		Password:  "vnandasa",
	}

	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(resp)

}
