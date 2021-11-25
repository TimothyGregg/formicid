package db

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"github.com/TimothyGregg/formicid/game"
	"github.com/TimothyGregg/formicid/game/elements"
	option "google.golang.org/api/option"
)

type FirebaseDatabase struct {
	Client *db.Client
}

func (database *FirebaseDatabase) StoreGame(root string, g *game.Game) error {
	path := fmt.Sprintf("%sgame/%d", root, g.UID.Value())
	err := database.SetRef(path, g)
	if err != nil {
		return err
	}
	err = database.StoreBoard(path, g.Board)
	if err != nil {
		return err
	}
	return nil
}

func (database *FirebaseDatabase) StoreBoard(root string, b *elements.Board) error {
	path := fmt.Sprintf("%sboard/%d", root, b.UID.Value())
	err := database.SetRef(path, b)
	if err != nil {
		return err
	}
	for _, node := range b.Nodes {
		err = database.StoreNode(root, node)
		if err != nil {
			return err
		}
	}
	for _, path := range b.Paths {
		err = database.StorePath(root, path)
		if err != nil {
			return err
		}
	}
	return nil
}

func (database *FirebaseDatabase) StoreNode(root string, n *elements.Node) error {
	return database.SetRef(fmt.Sprintf("%snode/%d", root, n.UID.Value()), n)
}

func (database *FirebaseDatabase) StorePath(root string, p *elements.Path) error {
	return database.SetRef(fmt.Sprintf("%s/path/%d", root, p.UID.Value()), p)
}

func (database *FirebaseDatabase) GetRef(ref string, v interface{}) error {
	return database.Client.NewRef(ref).Get(context.Background(), v)
}

func (database *FirebaseDatabase) SetRef(ref string, v interface{}) error {
	return database.Client.NewRef(ref).Set(context.Background(), v)
}

func GetDatabase() (*FirebaseDatabase, error) {
	conf := &firebase.Config{DatabaseURL: os.Getenv("FIREBASE_URL"), ProjectID: "formicid", StorageBucket: "formicid.appspot.com"}
	opt := option.WithCredentialsFile("/home/tim/Desktop/formicid-firebase-pk.json")
	app, err := firebase.NewApp(context.Background(), conf, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	client, err := app.Database(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error initializing client: %v", err)
	}
	return &FirebaseDatabase{client}, err
}

// func insert(ctx context.Context, client *db.Client, path string, v interface{}) {
// 	ref := client.NewRef(path)
// 	ref.Transaction(ctx, func(tn db.TransactionNode) (interface{}, error) {

// 		return nil, errors.New("Help")
// 	})
// }
