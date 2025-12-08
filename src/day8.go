package src

import (
	"fmt"
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Point3D struct {
	x float64
	y float64
	z float64
}

func (a *Point3D) dist(b Point3D) float64 {
	return math.Sqrt(math.Pow(b.x-a.x, 2) + math.Pow(b.y-a.y, 2) + math.Pow(b.z-a.z, 2))
}

func (a *Point3D) equal(b Point3D) bool {
	return a.x == b.x && a.y == b.y && a.z == b.z
}

func parsePoints(input string) []Point3D {
	points := []Point3D{}

	for line := range strings.SplitSeq(input, "\n") {
		coords := strings.Split(line, ",")

		x, err := strconv.ParseFloat(coords[0], 64)

		if err != nil {
			panic(err)
		}

		y, err := strconv.ParseFloat(coords[1], 64)

		if err != nil {
			panic(err)
		}

		z, err := strconv.ParseFloat(coords[2], 64)

		if err != nil {
			panic(err)
		}

		points = append(points, Point3D{x: x, y: y, z: z})
	}

	return points
}

type Circuit struct {
	junctionBoxes []Point3D
}

func (c *Circuit) hasPoint(p Point3D) bool {
	for _, box := range c.junctionBoxes {
		if box.equal(p) {
			return true
		}
	}

	return false
}

func (c *Circuit) size() int {
	return len(c.junctionBoxes)
}

type Path struct {
	from Point3D
	to   Point3D
	dist float64
}

func pathsBetween(points []Point3D) []Path {
	paths := []Path{}

	for indexA, pointA := range points {
		for _, pointB := range points[indexA+1:] {
			paths = append(paths, Path{from: pointA, to: pointB, dist: pointA.dist(pointB)})
		}
	}

	sort.Slice(paths, func(i, j int) bool {
		return paths[i].dist < paths[j].dist
	})

	return paths
}

func mergeCircuits(circuits []Circuit, paths []Path) ([]Circuit, Path) {
	for _, path := range paths {
		circuitWithFromIndex := slices.IndexFunc(circuits, func(c Circuit) bool {
			return c.hasPoint(path.from)
		})

		circuitWithToIndex := slices.IndexFunc(circuits, func(c Circuit) bool {
			return c.hasPoint(path.to)
		})

		if circuitWithFromIndex == circuitWithToIndex {
			continue
		}

		circuits[circuitWithFromIndex].junctionBoxes = append(circuits[circuitWithFromIndex].junctionBoxes, circuits[circuitWithToIndex].junctionBoxes...)

		nextCircuits := []Circuit{}

		for circuitIndex, c := range circuits {
			if circuitIndex == circuitWithToIndex {
				continue
			}

			nextCircuits = append(nextCircuits, c)
		}

		circuits = nextCircuits

		if len(circuits) == 1 {
			return circuits, path
		}
	}

	return circuits, paths[len(paths)-1]
}

func Day8Part1(input string) {
	points := parsePoints(input)
	paths := pathsBetween(points)
	circuits := []Circuit{}

	for _, p := range points {
		circuits = append(circuits, Circuit{junctionBoxes: []Point3D{p}})
	}

	merged, _ := mergeCircuits(circuits, paths[0:1000])

	sort.Slice(merged, func(i, j int) bool {
		return merged[i].size() >= merged[j].size()
	})

	fmt.Printf("Result = %d\n", merged[0].size()*merged[1].size()*merged[2].size())
}

func Day8Part2(input string) {
	points := parsePoints(input)
	paths := pathsBetween(points)
	circuits := []Circuit{}

	for _, p := range points {
		circuits = append(circuits, Circuit{junctionBoxes: []Point3D{p}})
	}

	_, lastPath := mergeCircuits(circuits, paths)

	fmt.Printf("Result = %d\n", int(lastPath.from.x*lastPath.to.x))
}
