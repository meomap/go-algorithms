// Package closest provides ClosestPair Implementation
package closest

import (
	"fmt"
	"math"

	merge "github.com/meomap/go-algorithms/algorithms/sorting/merge-sort"
)

// points sorted by x axis
var compareByX = func(a, b interface{}) int {
	p1, ok := a.(Point)
	if !ok {
		panic("Expect a to be a Point struct")
	}
	p2, ok := b.(Point)
	if !ok {
		panic("Expect b to be a Point struct")
	}

	if p1.X > p2.X {
		return 1
	} else if p1.X == p2.X {
		return 0
	}
	return -1
}

func sortPointsByDimension(in []Point, compareFn merge.CompareFunc) (out []Point) {
	lenIn := len(in)
	points := make([]interface{}, len(in))
	out = make([]Point, lenIn)
	for i := 0; i < lenIn; i++ {
		points[i] = in[i]
	}
	points = merge.Sort(points, compareFn)
	for i := 0; i < lenIn; i++ {
		out[i], _ = points[i].(Point)
	}
	return
}

// points sorted by y axis
var compareByY = func(a, b interface{}) int {
	p1, ok := a.(Point)
	if !ok {
		panic("Expect a to be a Point struct")
	}
	p2, ok := b.(Point)
	if !ok {
		panic("Expect b to be a Point struct")
	}

	if p1.Y > p2.Y {
		return 1
	} else if p1.Y == p2.Y {
		return 0
	}
	return -1
}

type Point struct {
	X float64
	Y float64
}

func (p Point) String() string {
	return fmt.Sprintf("(%f,%f)", p.X, p.Y)
}

// Line represents straight line connected two Point in coordinates
type Line struct {
	Start  Point
	End    Point
	Length float64
}

func (l Line) String() string {
	return fmt.Sprintf("start=%q end=%q len=%f", l.Start, l.End, l.Length)
}

func NewPoint(x, y float64) Point {
	return Point{X: x, Y: y}
}

func calcDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

func closestSplitPairs(in []Point, delta float64) Line {
	// sort by y first to ensure vertial expanding
	inSortedY := sortPointsByDimension(in, compareByY)
	shortest := Line{Length: 9999}
	lenIn := len(inSortedY)

	// only need to expand to atmost 7 nodes as proved
	// by correctness that there's no more pair with
	// shorter path than delta outside 8 box
	var scanTimes int
	if lenIn > 7 {
		scanTimes = lenIn - 7
	} else {
		scanTimes = lenIn
	}

	for i := 0; i < scanTimes; i++ {
		for j := 1; j <= 7 && i+j < lenIn-1; j++ {
			line := NewLine(inSortedY[i], inSortedY[i+j])
			if line.Length < shortest.Length {
				shortest = line
			}
		}
	}
	return shortest
}

func NewLine(start, end Point) Line {
	distance := calcDistance(start.X, start.Y, end.X, end.Y)
	return Line{
		Start:  start,
		End:    end,
		Length: distance,
	}
}

// BruteForce scan through list of points to find out line with shortest length
func BruteForce(in []Point) Line {
	shortest := Line{}
	lenIn := len(in)
	for i := 0; i < lenIn; i++ {
		for j := i + 1; j < lenIn; j++ {
			line := NewLine(in[i], in[j])
			if shortest.Length == 0 || line.Length < shortest.Length {
				shortest = line
			}
		}
	}
	return shortest
}

func minLine(line1, line2, line3 Line) Line {
	minDist := math.Min(math.Min(line1.Length, line2.Length), line3.Length)
	switch minDist {
	case line1.Length:
		return line1
	case line2.Length:
		return line2
	default:
		return line3
	}
}

func DivideAndConquor(in []Point) Line {
	// base case
	if len(in) == 2 {
		return NewLine(in[0], in[1])
	} else if len(in) < 2 {
		return Line{Length: 99999}
	}

	inSortedX := sortPointsByDimension(in, compareByX)
	lenIn := len(in)

	// must round up to ceilling to prevent single element left at base stack
	midIdx := lenIn/2 + 1
	midPoint := inSortedX[midIdx]
	if midIdx >= lenIn {
		fmt.Printf("inSortedX = %+v\n", inSortedX)
		fmt.Printf("midIdx = %+v\n", midIdx)
	}
	// min distance found at left & right
	lLine := DivideAndConquor(inSortedX[:midIdx])
	rLine := DivideAndConquor(inSortedX[midIdx:])

	minDist := math.Min(lLine.Length, rLine.Length)
	// build list of split points within mid vertical x
	// that span around min distance
	var splitPoints []Point
	for i := 0; i < lenIn; i++ {
		p := inSortedX[i]
		if calcDistance(p.X, p.Y, midPoint.X, midPoint.Y) <= minDist {
			splitPoints = append(splitPoints, p)
		}
	}
	splitLine := closestSplitPairs(splitPoints, minDist)

	return minLine(lLine, rLine, splitLine)
}
