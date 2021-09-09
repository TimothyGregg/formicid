package main

import (
	"fmt"
	"log"
	"math"
	"net/http"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	webpage()
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

func local_draw() {
	g := NewGame(1920, 1080, 1000)

	size := g.Board.GetSize()
	border := float64(50)
	rl.InitWindow(int32(size[0] + border), int32(size[1] + border), "raylib [core] example - basic window")

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		for _, node := range g.Board.Nodes {
			x, y, _, _, r := node.Get()
			rl.DrawCircle(int32(x + int(border)/2), int32(y + int(border)/2), r, rl.Lime)
		}

		for _, path := range g.Board.Paths {
			v1, v2 := path.Get()
			x1, y1 := v1.Get()
			x2, y2 := v2.Get()
			rl.DrawLine(int32(x1 + int(border)/2), int32(y1 + int(border)/2), int32(x2 + int(border)/2), int32(y2 + int(border)/2), rl.Red)
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
	fmt.Println("E/V^2 = " + fmt.Sprint(float64(len(g.Board.Paths))/math.Pow(float64(len(g.Board.Nodes)), 2.0)))
}
