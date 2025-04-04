package main

import (
	"fmt"

	"math/rand/v2"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"

	"snake/try"
)

func run() {
	maxX := 800
	maxY := 600
	cfg := pixelgl.WindowConfig{
		Title:  "Canvas interactif en Go",
		Bounds: pixel.R(0, 0, float64(maxX), float64(maxY)),
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
	var isWin bool = false
	randomPoint := [2]int{rand.IntN(maxX), rand.IntN(maxY)}
	var radius int = 10

	for !win.Closed() || isWin {
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
        head := dequePosition.GetFront()
        isInBounds := try.IsInBounds(head, maxX, maxY)
        if !isInBounds {
            isWin = false
            break
        }
        isTouching := try.HandleSnakeApple(dequePosition, randomPoint, radius * 2)
        if isTouching {
            for i := 0; i < radius * 2 ; i++{
                tail := dequePosition.GetQueue()
                newTail := []int{tail[0] - currentDirection[0], tail[1] - currentDirection[1]}
                dequePosition.PushBack(newTail)
            }
            randomPoint = [2]int{rand.IntN(maxX), rand.IntN(maxY)}
        }
        
        imd.Clear()
        for _, pos := range dequePosition.Data {
            imd.Color = colornames.Green
            imd.Push(pixel.V(float64(pos[0]), float64(pos[1])))
            imd.Circle(float64(radius), 0)
        }

		imd.Color = colornames.Red
		imd.Push(pixel.V(float64(randomPoint[0]), float64(randomPoint[1])))

		win.Clear(colornames.Black)
		imd.Draw(win)
		win.Update()
	}

}

func main() {
	pixelgl.Run(run)
}
