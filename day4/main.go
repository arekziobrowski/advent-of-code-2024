package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
)

func main() {
	part1()
}

func part1() {
	grid := readGrid()
	directions := []image.Point{
		{0, -1}, {1, 0}, {0, 1}, {-1, 0},
		{-1, -1}, {1, -1}, {1, 1}, {-1, 1},
	}

	var found []image.Point

	gridWidth := len(grid)
	gridHeight := gridWidth

	for point := range grid {
		for _, dir := range directions {
			if searchWord(grid, gridWidth, gridHeight, point, dir, "XMAS") {
				found = append(found, point)
			}
		}
	}

	fmt.Println("Part 1 -", len(found))
}

func searchWord(grid map[image.Point]rune, gridWidth, gridHeight int, start image.Point, dir image.Point, word string) bool {
	x, y := start.X, start.Y
	for _, char := range word {
		if x < 0 || y < 0 || y >= gridHeight || x >= gridWidth {
			return false
		}
		if grid[image.Point{X: x, Y: y}] != char {
			return false
		}
		x += dir.X
		y += dir.Y
	}
	return true
}

func readGrid() map[image.Point]rune {
	f, err := os.Open("day4/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	grid := make(map[image.Point]rune)
	scanner := bufio.NewScanner(f)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, char := range line {
			grid[image.Point{X: x, Y: y}] = char
		}
		y++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return grid
}
