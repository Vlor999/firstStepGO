package main


import (
	"fmt"

	"math/rand"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
    "github.com/faiface/pixel/pixelgl"
    "golang.org/x/image/colornames"
    "golang.org/x/image/font/basicfont"
    "github.com/faiface/pixel/text"

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
            fmt.Println("Direction changed to:", currentDirection)
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
        isWin = false
        
        isTouching int = 0
        compteur int = 1
    )

	cfg := pixelgl.WindowConfig{
		Title:  "Go Snake",
		Bounds: pixel.R(0, 0, float64(maxX), float64(maxY)),
		VSync:  true,
	}

	snakeWindow, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

    compteur_window, err2 := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  "Compteur",
		Bounds: pixel.R(0, 0, 200, 100),
		VSync:  true,
	})
	if err2 != nil {
		panic(err2)
	}
    

    atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	txt := text.New(pixel.V(20, 70), atlas)
	imd := imdraw.New(nil)

	dequePosition := &try.Deque{}
    
	randomPoint := get2dRandPoint(radius, maxX - radius, radius,  maxY - radius)
    
	for !snakeWindow.Closed() || isWin {
		lastDirection = currentDirection
        txt.Clear()
        switch {
            case snakeWindow.Pressed(pixelgl.KeyRight):
                currentDirection = [2]int{1, 0}
            case snakeWindow.Pressed(pixelgl.KeyLeft):
                currentDirection = [2]int{-1, 0}
            case snakeWindow.Pressed(pixelgl.KeyUp):
                currentDirection = [2]int{0, 1}
            case snakeWindow.Pressed(pixelgl.KeyDown):
                currentDirection = [2]int{0, -1}
        }

        currentDirection = update_current(currentDirection, lastDirection)
    
        isGood := dequePosition.Verify(radius)
        if !isGood{
            fmt.Println("Collision")
            break
        }

        try.UpdateMap(currentDirection, dequePosition, isTouching == 0)
        head := dequePosition.GetFront()
        isInBounds := try.IsInBounds(head, maxX, maxY)
        if !isInBounds {
            isWin = false
            break
        }

        update_randomPoint_and_touching(dequePosition, &randomPoint, &isTouching, &compteur, radius, maxX, maxY)
        
        imd.Clear()
        for _, pos := range dequePosition.Data {
            imd.Color = colornames.Green
            imd.Push(pixel.V(float64(pos[0]), float64(pos[1])))
            imd.Circle(float64(radius), 0)
        }
        
		imd.Color = colornames.Red
		imd.Push(pixel.V(float64(randomPoint[0]), float64(randomPoint[1])))
        
		snakeWindow.Clear(colornames.Black)
		imd.Draw(snakeWindow)
		snakeWindow.Update()
        
        compteur_window.Clear(colornames.Black)
        fmt.Fprintf(txt, "%d", compteur)
        txt.Draw(compteur_window, pixel.IM)
        compteur_window.Update()
	}

}

func main() {
	pixelgl.Run(run)
}
