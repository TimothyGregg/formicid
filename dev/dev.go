package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"os"

	firebase "firebase.google.com/go"
	game "github.com/TimothyGregg/formicid/game"
	"github.com/TimothyGregg/formicid/game/util/uid"
	rl "github.com/gen2brain/raylib-go/raylib"
	"google.golang.org/api/option"
)

func main() {
	for _, arg := range os.Args[1:] {
		switch arg {
		case "-s":
			//go api
		case "-l":
			defer local_draw()
		case "-j":
			print_json()
		case "-f":
			firebase_dump()
		}
	}
}

func firebase_dump() {
	conf := &firebase.Config{DatabaseURL: os.Getenv("FIREBASE_URL"), ProjectID: "formicid", StorageBucket: "formicid.appspot.com"}
	opt := option.WithCredentialsFile("/home/tim/Desktop/formicid-firebase-pk.json")
	app, err := firebase.NewApp(context.Background(), conf, opt)
	if err != nil {
		panic(fmt.Errorf("error initializing app: %v", err))
	}
	client, err := app.Database(context.Background())
	if err != nil {
		panic(err)
	}
	newgame := game.New_Game(uid.NewUID(0), 100, 100)
	ref := client.NewRef("/games/" + fmt.Sprint(newgame.UID.Value()))
	ref.Set(context.Background(), newgame)
}

func print_json() {
	g := game.New_Game(uid.NewUID(0), 50, 50)
	data, err := json.MarshalIndent(g, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}

func local_draw() {
	g := game.New_Game(uid.NewUID(0), 50, 50)

	size := g.Board.Size
	fmt.Println(size)
	border := 50
	rl.InitWindow(int32(size[0]+border), int32(size[1]+border), "raylib [core] example - basic window")

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		for _, node := range g.Board.Nodes {
			x, y, r := node.X, node.Y, node.Radius
			rl.DrawCircle(int32(x+int(border)/2), int32(y+int(border)/2), float32(r), rl.Lime)
			rl.DrawText(fmt.Sprint(node.UID), int32(x+int(border)/2+5), int32(y+int(border)/2+5), 20, rl.Blue)
		}

		/* for _, path := range g.Board.Paths {
			v1, v2 := path.Vertices()
			x1, y1 := v1.Position()
			x2, y2 := v2.Position()
			rl.DrawLine(int32(x1+int(border)/2), int32(y1+int(border)/2), int32(x2+int(border)/2), int32(y2+int(border)/2), rl.Red)
		} */

		for node_uid, arr := range g.Board.NodeConnections {
			for _, other := range arr {
				node, err := g.Board.NodeByID(node_uid)
				if err != nil {
					os.Exit(1)
				}
				x2, y2 := other.X, other.Y
				rl.DrawLine(int32(node.X+int(border)/2), int32(node.Y+int(border)/2), int32(x2+int(border)/2), int32(y2+int(border)/2), rl.Red)
			}
		}

		rl.EndDrawing()
		//start := time.Now()
		g.Board.Update()
		//fmt.Println(time.Since(start))
		//fmt.Println(int(g.Board.Element.Timer()))
	}

	rl.CloseWindow()
	fmt.Println("E/V^2 = " + fmt.Sprint(float64(len(g.Board.Paths))/math.Pow(float64(len(g.Board.Nodes)), 2.0)))
}
