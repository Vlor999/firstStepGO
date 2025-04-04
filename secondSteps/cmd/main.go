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

/*
 get2dRandPoint generates a random 2D point within the specified ranges.
 The x-coordinate is randomly selected between minX (inclusive) and maxX (exclusive),
 and the y-coordinate is randomly selected between minY (inclusive) and maxY (exclusive).
 Parameters:
   - minX: The minimum value for the x-coordinate.
   - maxX: The maximum value for the x-coordinate (exclusive).
   - minY: The minimum value for the y-coordinate.
   - maxY: The maximum value for the y-coordinate (exclusive).
 Returns:
   A 2-element array [x, y] representing the randomly generated 2D point.
*/
func get2dRandPoint(minX int, maxX int, minY int, maxY int) [2]int {
    return [2]int{
        rand.Intn(maxX-minX) + minX,
        rand.Intn(maxY-minY) + minY,
    }
}

/*
 update_current updates the current direction based on the last direction.
 If the current direction is the exact opposite of the last direction, 
 it reverts to the last direction to prevent invalid movement. Otherwise, 
 it logs the change in direction. The function returns the updated direction.

 Parameters:
 - currentDirection: An array of two integers representing the current direction.
 - lastDirection: An array of two integers representing the last direction.

 Returns:
 - An array of two integers representing the updated direction.
*/
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


/*
update_randomPoint_Touching updates the random point and the touching state based on the current state of the deque position.
It checks if the random point is "touching" the snake (dequePosition) using the HandleSnakeApple function.
If the random point is touching, it updates the isTouching value and potentially generates a new random point.

Parameters:
- dequePosition: A pointer to a Deque representing the snake's position.
- randomPoint: A pointer to a 2D integer array representing the current random point.
- isTouching: A pointer to an integer representing the current touching state.
- radius: An integer representing the radius used for collision detection.
- maxX: An integer representing the maximum X-coordinate for generating random points.
- maxY: An integer representing the maximum Y-coordinate for generating random points.

Behavior:
- The function calculates the current touching state using HandleSnakeApple.
- If the random point is touching the snake, the isTouching value is updated.
- If the isTouching value reaches twice the radius, a new random point is generated.
- The isTouching value is decremented after processing.
*/
func update_randomPoint_and_touching(dequePosition *try.Deque, randomPoint *[2]int, isTouching *int, radius int, maxX int, maxY int) {
    currentTouching := try.HandleSnakeApple(dequePosition, *randomPoint, radius * 2)
    *isTouching = max(*isTouching, currentTouching)
    if *isTouching > 0 {
        if *isTouching == 2 * radius {
            *randomPoint = get2dRandPoint(radius, maxX - radius, radius,  maxY - radius)
        }
        *isTouching--
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

        currentDirection = update_current(currentDirection, lastDirection)
    
        isGood := dequePosition.Verify(radius)
        if !isGood{
            fmt.Println("Collision")
            break
        }

        dequePosition = try.UpdateMap(currentDirection, dequePosition, isTouching == 0)
        head := dequePosition.GetFront()
        isInBounds := try.IsInBounds(head, maxX, maxY)
        if !isInBounds {
            isWin = false
            break
        }

        update_randomPoint_and_touching(dequePosition, &randomPoint, &isTouching, radius, maxX, maxY)
        
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
