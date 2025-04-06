package main

import (
	"fmt"
	"math/rand"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"

	"snake/try"
)

func get2dRandPoint(minX int, maxX int, minY int, maxY int) [2]int {
    return [2]int{
        rand.Intn(maxX-minX) + minX,
        rand.Intn(maxY-minY) + minY,
    }
}

func update_current(currentDirection [2]int, lastDirection [2]int) [2]int {
    if lastDirection != currentDirection {
        if currentDirection[0] == -lastDirection[0] && currentDirection[1] == -lastDirection[1] {
            currentDirection = lastDirection
        } else {
            fmt.Print("Direction changed to:", currentDirection, "\r")
        }
    }
    return currentDirection
}

func update_randomPoint_and_touching(dequePosition *try.Deque, randomPoint *[2]int, isTouching *int, compteur *int, radius int, maxX int, maxY int) {
    currentTouching := try.HandleSnakeApple(dequePosition, *randomPoint, radius * 2)
    *isTouching = max(*isTouching, currentTouching)
    if *isTouching > 0 {
        if *isTouching == 2 * radius {
            *randomPoint = get2dRandPoint(radius, maxX - radius, radius,  maxY - radius)
            *compteur++
        }
        *isTouching--
    }
}

func run() {
    const (
        maxX int = 800
        maxY int = 600
        radius int = 10
    )

	var (
		lastDirection [2]int
		currentDirection [2]int
		isWin bool = true
		isTouching int = 0
		compteur int = 1
		lastScore int = -1
	)

	cfg := pixelgl.WindowConfig{
		Title: "Go Snake",
		Bounds: pixel.R(0, 0, float64(maxX), float64(maxY)),
		VSync: true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	txt := text.New(pixel.V(20, float64(maxY)-30), atlas)
	imd := imdraw.New(nil)
	dequePosition := &try.Deque{}
	randomPoint := get2dRandPoint(radius, maxX-radius, radius, maxY-radius)

	for !win.Closed() && isWin {
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

		currentDirection = update_current(currentDirection, lastDirection)

		if !dequePosition.Verify(radius) {
            isWin = false
			break
		}

		try.UpdateMap(currentDirection, dequePosition, isTouching == 0)
		head := dequePosition.GetFront().GetValue()
		if !try.IsInBounds(head, maxX, maxY) {
			isWin = false
			break
		}

		update_randomPoint_and_touching(dequePosition, &randomPoint, &isTouching, &compteur, radius, maxX, maxY)

        if compteur == 100 {
            break // isWin = true
        }

		win.Clear(colornames.Black)
		imd.Clear()

		for node := dequePosition.GetFront(); node != nil; node = node.GetNext() {
			imd.Color = colornames.Limegreen
			imd.Push(pixel.V(float64(node.GetValue()[0]), float64(node.GetValue()[1])))
			imd.Circle(float64(radius), 0)
		}
		

		imd.Color = colornames.Red
		imd.Push(pixel.V(float64(randomPoint[0]), float64(randomPoint[1])))
		imd.Circle(float64(radius), 0)

		imd.Color = pixel.RGBA{R: 0.2, G: 0.2, B: 0.2, A: 0.7}
		imd.Push(pixel.V(10, float64(maxY)-10))
		imd.Push(pixel.V(150, float64(maxY)-40))
		imd.Rectangle(0)

		imd.Draw(win)

		if compteur != lastScore {
			txt.Clear()
			fmt.Fprintf(txt, "Score: %d", compteur)
			lastScore = compteur
		}
		txt.Draw(win, pixel.IM)

		win.Update()
	}

    if !isWin {
        for !win.Closed(){
            if win.Pressed(pixelgl.KeyEscape){
                win.SetClosed(true)
            }
            win.Clear(colornames.Black)
		    imd.Clear()
            txt.Clear()

            fmt.Fprintf(txt, "Press ESC to close the game\n\nGame Over: %d", compteur)
            txt.Draw(win, pixel.IM)
            win.Update()
        }
    } else {
        for !win.Closed(){
            if win.Pressed(pixelgl.KeyEscape){
                win.SetClosed(true)
            }
            win.Clear(colornames.Black)
		    imd.Clear()
            txt.Clear()

            fmt.Fprintf(txt, "Press ESC to close the game\n\nEnd of the game : %d", compteur)
            txt.Draw(win, pixel.IM)
            win.Update()
        }
    }
}

func main() {
	pixelgl.Run(run)
}
