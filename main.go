package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"

	game "github.com/TimothyGregg/Antmound/game"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	for _, arg := range os.Args[1:] {
		switch arg {
		case "-w":
			defer webpage()
		case "-l":
			local_draw()
		case "-j":
			print_json()
		}
	}
}

func webpage() {
	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)
	log.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func print_json() {
	g := game.NewGame(100, 100, 10)
	data, _ := json.MarshalIndent(g.Board, "", "\t")
	fmt.Println(string(data))
}

func local_draw() {
	g := game.NewGame(1820, 980, 1000)

	size := g.Board.Get_Size()
	border := 50
	rl.InitWindow(int32(size[0]+border), int32(size[1]+border), "raylib [core] example - basic window")

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		for _, node := range g.Board.Nodes {
			x, y, r := node.Get()
			rl.DrawCircle(int32(x+int(border)/2), int32(y+int(border)/2), float32(r), rl.Lime)
			rl.DrawText(fmt.Sprint(node.UID), int32(x+int(border)/2+5), int32(y+int(border)/2+5), 20, rl.Blue)
		}

		/* for _, path := range g.Board.Paths {
			v1, v2 := path.Vertices()
			x1, y1 := v1.Position()
			x2, y2 := v2.Position()
			rl.DrawLine(int32(x1+int(border)/2), int32(y1+int(border)/2), int32(x2+int(border)/2), int32(y2+int(border)/2), rl.Red)
		} */

		for node, arr := range g.Board.Get_node_connections() {
			for _, other := range arr {
				x1, y1, _ := node.Get()
				x2, y2, _ := other.Get()
				rl.DrawLine(int32(x1+int(border)/2), int32(y1+int(border)/2), int32(x2+int(border)/2), int32(y2+int(border)/2), rl.Red)
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
