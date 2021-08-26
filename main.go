package main

// "github.com/gen2brain/raylib-go/raylib"

func main() {
	b := NewBoard()
	err := b.Add_Node([2]int{1, 2})
	err = b.Add_Node([2]int{2, 2})
	err = b.Connect_Nodes(b.nodes[0], b.nodes[1])
	if err != nil {
		panic(err)
	}
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
