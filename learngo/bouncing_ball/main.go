package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

const maxX = 20
const maxY = 50

func main() {
	//var positions [maxX][maxY]bool
	xpos, ypos := 0, 0
	var backward bool = true
	var forward bool = false
	var right bool = false
	var left bool = true
	fmt.Println()
	for {
		//fmt.Printf("xpos: %d, ypos: %d\n", xpos, ypos)
		screen.Clear()
		screen.MoveTopLeft()
		printBall(xpos, ypos)
		if isXBoundry(xpos) {
			backward = !backward
			forward = !forward
		}
		if isYBoundry(ypos) {
			left = !left
			right = !right
		}

		if forward {
			xpos = xpos + 1
		} else {
			xpos = xpos - 1
		}

		if right {
			ypos = ypos + 1
		} else {
			ypos = ypos - 1
		}

		time.Sleep(1 * time.Second)
	}
}

func isXBoundry(x int) bool {
	if x >= (maxX-1) || x <= 0 {
		return true
	}
	return false
}

func isYBoundry(y int) bool {
	if y >= (maxY-1) || y <= 0 {
		return true
	}
	return false
}

func printBall(x int, y int) {
	buffer := make([]rune, 0, maxX*maxY)
	for i := 0; i < maxX; i++ {
		for j := 0; j < maxY; j++ {
			if i == x && j == y {
				buffer = append(buffer, '⚽')
			} else {
				buffer = append(buffer, 'x')
			}
		}
		buffer = append(buffer, '\n')
	}
	fmt.Println(string(buffer))
}
