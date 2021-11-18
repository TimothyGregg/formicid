package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/rds/auth"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error: " + err.Error())
	}

	authToken, err := auth.BuildAuthToken(
		context.TODO(),
		os.Getenv("formicid_db_server")+":"+os.Getenv("formicid_db_port"), // Database Endpoint (With Port)
		os.Getenv("formicid_db_region"),                                   // AWS Region
		os.Getenv("formicid_db_user"),                                     // Database Account
		cfg.Credentials,
	)
	if err != nil {
		panic("failed to create authentication token: " + err.Error())
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		os.Getenv("formicid_db_server"),
		os.Getenv("formicid_db_port"),
		os.Getenv("formicid_db_user"),
		authToken,
		os.Getenv("formicid_db_name"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Amazon make things painful
	// connString := "postgresql://" + os.Getenv("formicid_db_user") + ":" +
	// 	os.Getenv("formicid_db_pass") + "@" +
	// 	os.Getenv("formicid_db_server") + ":" +
	// 	os.Getenv("formicid_db_port") + "/" +
	// 	os.Getenv("formicid_db_name")
	// conn, err := pgx.Connect(context.Background(), connString)
	// fmt.Println(connString)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	// 	os.Exit(1)
	// }
	// defer conn.Close(context.Background())
	// os.Exit(0)
}
