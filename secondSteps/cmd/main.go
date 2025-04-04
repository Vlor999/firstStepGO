package main

import (
	"fmt"

	"math/rand"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"

	"snake/try"
)

func get2dRandPoint(minX int, maxX int, minY int, maxY int) [2]int {
    return [2]int{
        rand.Intn(maxX-minX) + minX,
        rand.Intn(maxY-minY) + minY,
    }
}

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

    var isWin = false
	var radius int = 10
	randomPoint := get2dRandPoint(radius, maxX - radius, radius,  maxY - radius)
    var isTouching int = 0

	for !win.Closed() || isWin {
		lastDirection = currentDirection
        switch {
            case win.Pressed(pixelgl.KeyRight):
                currentDirection = [2]int{1, 0}
            case win.Pressed(pixelgl.KeyLeft):
                currentDirection = [2]int{-1, 0}
            case win.Pressed(pixelgl.KeyUp):
                currentDirection = [2]int{0, 1}
            case win.Pressed(pixelgl.KeyDown):
                currentDirection = [2]int{0, -1}
        }

        if lastDirection != currentDirection {
            if currentDirection[0] == -lastDirection[0] && currentDirection[1] == -lastDirection[1] {
                currentDirection = lastDirection
            } else {
                dequePosition.Verify()
                fmt.Println("Direction changed to:", currentDirection)
            }
        }

        dequePosition = try.UpdateMap(currentDirection, dequePosition, isTouching == 0)
        head := dequePosition.GetFront()
        isInBounds := try.IsInBounds(head, maxX, maxY)
        if !isInBounds {
            isWin = false
            break
        }
        currentTouching := try.HandleSnakeApple(dequePosition, randomPoint, radius * 2)
        isTouching = max(isTouching, currentTouching)
        if isTouching > 0 {
            if isTouching == 2 * radius {
                randomPoint = get2dRandPoint(radius, maxX - radius, radius,  maxY - radius)
            }
            isTouching--
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
