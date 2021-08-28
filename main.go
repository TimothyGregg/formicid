package main

import (
	"fmt"
)

// "github.com/gen2brain/raylib-go/raylib"

func main() {
	b := NewBoard()
	b.SetSize([2]int{100, 100})
	b.Set_Node_Radius(3)
	b.Naive_Fill()
	fmt.Println(b)
	/*
		raylib.InitWindow(800, 450, "raylib [core] example - basic window")

		raylib.SetTargetFPS(60)

		for !raylib.WindowShouldClose() {
			raylib.BeginDrawing()

			raylib.ClearBackground(raylib.RayWhite)

			raylib.DrawText("Congrats! You created your first window!", 190, 200, 20, raylib.LightGray)

			raylib.EndDrawing()
		}

		raylib.CloseWindow()
	*/
}
