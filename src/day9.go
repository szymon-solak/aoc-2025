package src

import (
	"cmp"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Point2D struct {
	x int
	y int
}

func parseTiles(input string) []Point2D {
	points := []Point2D{}

	for line := range strings.SplitSeq(input, "\n") {
		coords := strings.Split(line, ",")

		x, err := strconv.Atoi(coords[0])
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(coords[1])
		if err != nil {
			panic(err)
		}

		points = append(points, Point2D{x: x, y: y})
	}

	return points
}

type Rect struct {
	x1   int
	x2   int
	y1   int
	y2   int
	area int
}

func getRectFromPoints(a, b Point2D) Rect {
	x := math.Abs(float64(a.x-b.x)) + 1
	y := math.Abs(float64(a.y-b.y)) + 1
	area := int(x * y)

	if a.x > b.x {
		if a.y > b.y {
			return Rect{b.x, a.x, b.y, a.y, area}
		} else {
			return Rect{b.x, a.x, a.y, b.y, area}
		}
	} else {
		if a.y > b.y {
			return Rect{a.x, b.x, b.y, a.y, area}
		} else {
			return Rect{a.x, b.x, a.y, b.y, area}
		}
	}
}

func getRects(points []Point2D) []Rect {
	rects := []Rect{}

	for index, pointA := range points {
		for _, pointB := range points[index+1:] {
			rects = append(rects, getRectFromPoints(pointA, pointB))
		}
	}

	slices.SortFunc(rects, func(a, b Rect) int {
		return cmp.Compare(b.area, a.area)
	})

	return rects
}

func getEdges(points []Point2D) []Rect {
	edges := []Rect{}

	for index, point := range points {
		nextIndex := (index + 1) % len(points)
		nextPoint := points[nextIndex]
		edges = append(edges, getRectFromPoints(point, nextPoint))
	}

	slices.SortFunc(edges, func(a, b Rect) int {
		return cmp.Compare(b.area, a.area)
	})

	return edges
}

func getLargestRectInsidePolygon(rects []Rect, edges []Rect) *Rect {
rectLoop:
	for _, rect := range rects {
		for _, edge := range edges {
			if edge.x2 > rect.x1 && edge.x1 < rect.x2 && edge.y2 > rect.y1 && edge.y1 < rect.y2 {
				continue rectLoop
			}
		}

		return &rect
	}

	return nil
}

func Day9Part1(input string) {
	fmt.Println(getRects(parseTiles(input))[0].area)
}

func Day9Part2(input string) {
	tiles := parseTiles(input)
	fmt.Println(getLargestRectInsidePolygon(getRects(tiles), getEdges(tiles)).area)
}
