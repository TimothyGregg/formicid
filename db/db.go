package main

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go"
	option "google.golang.org/api/option"
)

func main() {
	app, err := getApp()
	fmt.Println(err)
	client, err := app.Database(context.Background())
	ref := client.NewRef("/test/2")
	ref.Set(context.Background(), "hello")
}

func getApp() (*firebase.App, error) {
	conf := &firebase.Config{DatabaseURL: os.Getenv("FIREBASE_URL"), ProjectID: "formicid", StorageBucket: "formicid.appspot.com"}
	opt := option.WithCredentialsFile("/home/tim/Desktop/formicid-firebase-pk.json")
	app, err := firebase.NewApp(context.Background(), conf, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return app, err
}
