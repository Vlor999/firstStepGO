package main

import (
    "fmt"

    "github.com/faiface/pixel"
    "github.com/faiface/pixel/imdraw"
    "github.com/faiface/pixel/pixelgl"
    "golang.org/x/image/colornames"

    "snake/try"
)

func run() {
    cfg := pixelgl.WindowConfig{
        Title:  "Canvas interactif en Go",
        Bounds: pixel.R(0, 0, 800, 600),
        VSync:  true,
    }

    win, err := pixelgl.NewWindow(cfg)
    if err != nil {
        panic(err)
    }

    imd := imdraw.New(nil)
    var lastDirection [2]int
    var currentDirection [2]int
    dequePosition := &try.Deque{}

    initialPos := []int{400, 300}
    dequePosition.PushBack(initialPos)
    dequePosition.Print()

    for !win.Closed() {
        if win.Pressed(pixelgl.MouseButtonLeft) {
            mousePos := win.MousePosition()
            imd.Color = colornames.Red
            imd.Push(mousePos)
            imd.Circle(5, 0)
        }
        
        lastDirection = currentDirection
        if win.Pressed(pixelgl.KeyRight) {
            currentDirection = [2]int{1, 0}
        }
        if win.Pressed(pixelgl.KeyLeft) {
            currentDirection = [2]int{-1, 0}
        }
        if win.Pressed(pixelgl.KeyUp) {
            currentDirection = [2]int{0, 1}
        }
        if win.Pressed(pixelgl.KeyDown) {
            currentDirection = [2]int{0, -1}
        }

        if lastDirection != currentDirection {
            fmt.Println("Direction changed to:", currentDirection)
            dequePosition.Print()
        }

        dequePosition = try.UpdateMap(currentDirection, dequePosition)
        
        imd.Clear()
        for _, pos := range dequePosition.Data {
            imd.Color = colornames.Green
            imd.Push(pixel.V(float64(pos[0]), float64(pos[1])))
            imd.Circle(10, 0)
        }

        win.Clear(colornames.Black)
        imd.Draw(win)
        win.Update()
    }
}

func main() {
    pixelgl.Run(run)
}