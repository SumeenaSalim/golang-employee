package dbconnect

import (
	"context"
	"log"
	"os"

	"github.com/edgedb/edgedb-go"
	"github.com/joho/godotenv"
)

var DbClient *edgedb.Client

func ConnectDatabase() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	opts := edgedb.Options{
		Database:    os.Getenv("EDGEDB_SERVER_DATABASE"),
		User:        os.Getenv("EDGEDB_SERVER_USER"),
		Password:    edgedb.NewOptionalStr(os.Getenv("EDGEDB_SERVER_PASSWORD")),
		Host:        "edgedb",
		TLSSecurity: os.Getenv("EDGEDB_CLIENT_TLS_SECURITY"),
		Port:        5656,
	}

	DbClient, err = edgedb.CreateClient(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}
}
