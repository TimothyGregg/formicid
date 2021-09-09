package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	g := NewGame(720, 480, 100)

	size := g.Board.GetSize()
	rl.InitWindow(int32(size[0]), int32(size[1]), "raylib [core] example - basic window")

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		for _, node := range g.Board.Nodes {
			x, y, _, _, r := node.Get()
			rl.DrawCircle(int32(x), int32(y), r, rl.Lime)
		}

		for _, path := range g.Board.Paths {
			v1, v2 := path.Get()
			x1, y1 := v1.Get()
			x2, y2 := v2.Get()
			rl.DrawLine(int32(x1), int32(y1), int32(x2), int32(y2), rl.Red)
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
