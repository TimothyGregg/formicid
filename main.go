package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	b := NewBoard()
	b.SetSize([2]float64{800, 450})
	b.Set_Node_Radius(10)
	b.Naive_Fill()
	b.Connect_Delaunay()
	fmt.Println(b)
	rl.InitWindow(800, 450, "raylib [core] example - basic window")

	rl.SetTargetFPS(60)
	fmt.Println(len(b.Get_Nodes()))
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)
		for _, node := range b.Get_Nodes() {
			x, y := node.Get()
			rl.DrawCircle(int32(x), int32(y), 10, rl.Lime)
		}

		for _, path := range b.Get_Paths() {
			v1, v2 := path.Get()
			x1, y1 := v1.Get()
			x2, y2 := v2.Get()
			rl.DrawLine(int32(x1), int32(y1), int32(x2), int32(y2), rl.Red)
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
