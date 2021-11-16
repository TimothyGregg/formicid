package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func main() {
	connString := "postgresql://" + os.Getenv("formicid_db_user") + ":" + 
	os.Getenv("formicid_db_pass") + "@" +
	os.Getenv("formicid_db_server") + ":" +
	os.Getenv("formicid_db_port") + "/" +
	os.Getenv("formicid_db_name")
	conn, err := pgx.Connect(context.Background(), connString)
	fmt.Println(connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	os.Exit(0)
}