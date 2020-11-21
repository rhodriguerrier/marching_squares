package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"image/color"
	"time"
	"math/rand"
	"fmt"
	"math"
)

func isoLine(nodes []int, dist float64) []int {
	activeNodes := ""
	for _, val := range nodes {
		if val > 2 {
			activeNodes += fmt.Sprint(1)
		} else {
			activeNodes += fmt.Sprint(0)
		}
	}
	
	/*Specifies the nodes pairs between which a contour line start or end will be placed.
	  0 = Top Left, 1 = Top Right, 2 = Bottom Right and 3 = Bottom Left
	*/
	lookupContour := map[string][]int {
		"0000": []int{},
		"0001": []int{0, 3, 3, 2},
		"0010": []int{1, 2, 3, 2},
		"0011": []int{0, 3, 1, 2},
		"0100": []int{0, 1, 1, 2},
		"0101": []int{0, 1, 0, 3, 1, 2, 3, 2},
		"0110": []int{0, 1, 3, 2},
		"0111": []int{0, 1, 0, 3},
		"1000": []int{0, 1, 0, 3},
		"1001": []int{0, 1, 3, 2},
		"1010": []int{0, 1, 1, 2, 0, 3, 3, 2},
		"1011": []int{0, 1, 1, 2},
		"1100": []int{0, 3, 1, 2},
		"1101": []int{1, 2, 3, 2},
		"1110": []int{0, 3, 3, 2},
		"1111": []int{},
	}
	pairs := lookupContour[activeNodes]
	if len(pairs) == 0 {
		return []int{}
	}
	var temp float64
	var firstNode int
	var secondNode int
	var linePoints []int
	var toAddX float64
	var toAddY float64
	for i := 0; i < len(pairs); i += 2 {
		firstNode = pairs[i]
		secondNode = pairs[i+1]
		temp = math.Abs(float64(nodes[firstNode] - nodes[secondNode]))
		toAddX = math.Abs(float64(2 - nodes[firstNode])) * (dist / temp)
		toAddY = math.Abs(float64(2 - nodes[firstNode])) * (dist / temp)
		if firstNode == 0 && secondNode == 1 {
			toAddY = 0.0
			linePoints = append(linePoints, int(math.Round(toAddX)), int(math.Round(toAddY)))
		} else if firstNode == 3 && secondNode == 2 {
			toAddY = dist
			linePoints = append(linePoints, int(math.Round(toAddX)), int(math.Round(toAddY)))
		} else if firstNode == 0 && secondNode == 3 {
			toAddX = 0.0
			linePoints = append(linePoints, int(math.Round(toAddX)), int(math.Round(toAddY)))
		} else if firstNode == 1 && secondNode == 2 {
			toAddX = dist
			linePoints = append(linePoints, int(math.Round(toAddX)), int(math.Round(toAddY)))
		}
	}
	return linePoints

}

func main() {
	a := app.New()
	w := a.NewWindow("Marching Squares")
	rand.Seed(time.Now().UTC().UnixNano())
	thresholdVal := 2
	nodeDiam := 6
	nodeDist := 22
	rect := &canvas.Rectangle{
		FillColor: color.Color(color.RGBA{0, 0, 0, 160}),
	}
	rect.Resize(fyne.Size{Width: 644, Height: 644})
	container := fyne.NewContainer(rect)
	var c *canvas.Circle
	arr := []*canvas.Circle{}
	nodeVals := []int{}
	var temp int
	r, g, b := uint8(0), uint8(0), uint8(0)
	for i := 0; i <= 638; i += nodeDist {
		for j := 0; j <= 638; j += nodeDist {
			temp = rand.Intn(6)
			if temp == thresholdVal {
				temp = 1
			}
			nodeVals = append(nodeVals, temp)
			c = &canvas.Circle{
				Position1: fyne.Position{X: j, Y: i}, Position2: fyne.Position{X: j+nodeDiam, Y: i+nodeDiam},
				FillColor: color.Color(color.RGBA{r, g, b, uint8(51*temp)}),
			}
			arr = append(arr, c)
			container.Objects = append(container.Objects, c)
		}
	}

	var squareVals []int
	var isoPoints []int
	var xPos int
	var yPos int
	var l *canvas.Line
	for i := 0; i < 869; i++ {
		squareVals = append(squareVals, nodeVals[i], nodeVals[i+1], nodeVals[i+31], nodeVals[i+30])
		isoPoints = isoLine(squareVals, float64(nodeDist))
		xPos = (*arr[i]).Position1.X + (nodeDiam / 2)
		yPos = (*arr[i]).Position1.Y + (nodeDiam / 2)
		for j := 0; j < len(isoPoints); j += 4 {
			l = &canvas.Line{
				Position1: fyne.Position{X: xPos+isoPoints[j], Y: yPos+isoPoints[j+1]}, Position2: fyne.Position{X: xPos+isoPoints[j+2], Y: yPos+isoPoints[j+3]},
				StrokeColor: color.Color(color.RGBA{255, 255, 255, 255}), StrokeWidth: float32(0.5),
			}
			container.Objects = append(container.Objects, l)
		}
		squareVals = nil
	}

	
	w.SetContent(container)
	w.Resize(fyne.NewSize(652,652))
	w.ShowAndRun()
}