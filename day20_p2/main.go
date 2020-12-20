package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/devries/advent_of_code_2020/utils"
)

const (
	Top int = iota
	Right
	Bottom
	Left
	TopRev
	RightRev
	BottomRev
	LeftRev
)

var monster = `                  # 
#    ##    ##    ###
 #  #  #  #  #  #   `

type Borders map[int][]rune

type Tile struct {
	Number    int
	Photo     map[utils.Point]rune
	Neighbors []int
}

type Transform func(map[utils.Point]rune, int, int) map[utils.Point]rune

func main() {
	f, err := os.Open("input.txt")
	utils.Check(err, "error opening file")
	defer f.Close()

	photos := parseInput(f)

	keys := make([]int, 0, len(photos))
	for k := range photos {
		keys = append(keys, k)
	}

	borderMap := make(map[int]Borders)

	for _, k := range keys {
		borderMap[k] = getBorders(photos[k])
	}

	borderMatches := make(map[int]map[int]int)

	for i := 0; i < len(keys); i++ {
		borderMatches[keys[i]] = make(map[int]int)
		for j := 0; j < len(keys); j++ {
			if j == i {
				continue
			}
			b1 := borderMap[keys[i]]
			b2 := borderMap[keys[j]]
		PhotoCompare:
			for side := 0; side < 4; side++ {
				bOrig := b1[side]
				for _, bCompare := range b2 {
					match := true
					for k := 0; k < 10; k++ {
						if bOrig[k] != bCompare[k] {
							// do next border
							match = false
							break
						}
					}
					// Matching border
					if match {
						borderMatches[keys[i]][side] = keys[j]
						break PhotoCompare
					}
				}
			}
		}
	}

	// Find the first corners
	var corner int
	for k, v := range borderMatches {
		if len(v) == 2 {
			corner = k
			break
		}
	}

	// Rotate and flip to make it upper right corner
	cornerPhoto := photos[corner]
	if borderMatches[corner][Top] != 0 && borderMatches[corner][Left] != 0 {
		cornerPhoto = rotateRight(cornerPhoto, 10, 10)
		cornerPhoto = rotateRight(cornerPhoto, 10, 10)
	} else if borderMatches[corner][Top] != 0 && borderMatches[corner][Right] != 0 {
		cornerPhoto = rotateRight(cornerPhoto, 10, 10)
	} else if borderMatches[corner][Bottom] != 0 && borderMatches[corner][Left] != 0 {
		cornerPhoto = flipPhoto(cornerPhoto, 10, 10)
	}

	currentTile := makeTile(cornerPhoto, corner, borderMatches[corner])
	tiles := make(map[utils.Point]Tile)
	tiles[utils.Point{0, 0}] = currentTile
	i := 0
	j := 0

	for {
		// Find neighbor to write if any
		currentBorders := getBorders(currentTile.Photo)
		rightBorder := currentBorders[Right]
		found := false
		var nextPhoto map[utils.Point]rune
		var nextPhotoId int
	SearchTiles:
		for _, n := range currentTile.Neighbors {
			nextPhoto = photos[n]
			nextBorders := getBorders(nextPhoto)
			for direction, border := range nextBorders {
				if equalBorders(border, rightBorder) {
					switch direction {
					case Top:
						nextPhoto = rotateRight(nextPhoto, 10, 10)
						nextPhoto = flipPhoto(nextPhoto, 10, 10)
					case Right:
						nextPhoto = flipPhoto(nextPhoto, 10, 10)
					case Bottom:
						nextPhoto = rotateRight(nextPhoto, 10, 10)
					case TopRev:
						nextPhoto = rotateRight(nextPhoto, 10, 10)
						nextPhoto = rotateRight(nextPhoto, 10, 10)
						nextPhoto = rotateRight(nextPhoto, 10, 10)
					case RightRev:
						nextPhoto = rotateRight(nextPhoto, 10, 10)
						nextPhoto = rotateRight(nextPhoto, 10, 10)
					case BottomRev:
						nextPhoto = flipPhoto(nextPhoto, 10, 10)
						nextPhoto = rotateRight(nextPhoto, 10, 10)
					case LeftRev:
						nextPhoto = rotateRight(nextPhoto, 10, 10)
						nextPhoto = rotateRight(nextPhoto, 10, 10)
						nextPhoto = flipPhoto(nextPhoto, 10, 10)
					}
					nextPhotoId = n
					found = true
					break SearchTiles
				}
			}
		}
		if found {
			currentTile = makeTile(nextPhoto, nextPhotoId, borderMatches[nextPhotoId])
			i++
			tiles[utils.Point{i, j}] = currentTile
			continue
		}
		// Didn't find neighbor to right, go back to beginning of row and look down
		i = 0
		currentTile = tiles[utils.Point{i, j}]
		currentBorders = getBorders(currentTile.Photo)
		bottomBorder := currentBorders[Bottom]

	SearchNextRow:
		for _, n := range currentTile.Neighbors {
			nextPhoto = photos[n]
			nextBorders := getBorders(nextPhoto)
			for direction, border := range nextBorders {
				if equalBorders(border, bottomBorder) {
					switch direction {
					case Right:
						nextPhoto = rotateRight(nextPhoto, 10, 10)
						nextPhoto = rotateRight(nextPhoto, 10, 10)
						nextPhoto = rotateRight(nextPhoto, 10, 10)
					case Bottom:
						nextPhoto = flipPhoto(nextPhoto, 10, 10)
						nextPhoto = rotateRight(nextPhoto, 10, 10)
						nextPhoto = rotateRight(nextPhoto, 10, 10)
					case Left:
						nextPhoto = rotateRight(nextPhoto, 10, 10)
						nextPhoto = flipPhoto(nextPhoto, 10, 10)
					case RightRev:
						nextPhoto = flipPhoto(nextPhoto, 10, 10)
						nextPhoto = rotateRight(nextPhoto, 10, 10)
					case BottomRev:
						nextPhoto = rotateRight(nextPhoto, 10, 10)
						nextPhoto = rotateRight(nextPhoto, 10, 10)
					case LeftRev:
						nextPhoto = rotateRight(nextPhoto, 10, 10)
					case TopRev:
						nextPhoto = flipPhoto(nextPhoto, 10, 10)
					}
					nextPhotoId = n
					found = true
					break SearchNextRow
				}
			}
		}

		if found {
			currentTile = makeTile(nextPhoto, nextPhotoId, borderMatches[nextPhotoId])
			j++
			tiles[utils.Point{i, j}] = currentTile
		} else {
			// Found everything I guess
			break
		}
	}

	// construct final image
	fullPhoto := makeImageFromTiles(tiles)
	width, height := photoSize(fullPhoto)

	// Encode Monster
	monsterLines := strings.Split(monster, "\n")
	monsterImage := make(map[utils.Point]rune)

	for y, line := range monsterLines {
		for x, r := range line {
			if r == '#' {
				monsterImage[utils.Point{x, y}] = '#'
			}
		}
	}
	hashesInMonster := len(monsterImage)

	hashesInImage := 0
	for _, v := range fullPhoto {
		if v == '#' {
			hashesInImage++
		}
	}

	// Search for Monster
	for _, f := range []Transform{rotateRight, rotateRight, rotateRight, flipPhoto, rotateRight, rotateRight, rotateRight, rotateRight} {
		found := findMonsterInPhoto(fullPhoto, monsterImage, width, height)
		if found != 0 {
			highlightMonsterInPhoto(fullPhoto, monsterImage, width, height)
			printPhoto(fullPhoto)
			fmt.Println(hashesInImage - found*hashesInMonster)
			break
		}
		fullPhoto = f(fullPhoto, width, height)
	}
}

func findMonsterInPhoto(photo map[utils.Point]rune, monsterImage map[utils.Point]rune, width int, height int) int {
	found := 0

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			isFound := true
			for pt := range monsterImage {
				if photo[utils.Point{i + pt.X, j + pt.Y}] != '#' {
					isFound = false
					break
				}
			}
			if isFound {
				found++
			}
		}
	}

	return found
}

func highlightMonsterInPhoto(photo map[utils.Point]rune, monsterImage map[utils.Point]rune, width int, height int) {
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			isFound := true
			for pt := range monsterImage {
				if photo[utils.Point{i + pt.X, j + pt.Y}] != '#' {
					isFound = false
					break
				}
			}
			if isFound {
				for pt := range monsterImage {
					photo[utils.Point{i + pt.X, j + pt.Y}] = 'R'
				}
			}
		}
	}
}

func parseInput(r io.Reader) map[int]map[utils.Point]rune {
	lines := utils.ReadLines(r)

	photos := make(map[int]map[utils.Point]rune)
	var gridTitle int
	var current map[utils.Point]rune
	var err error
	y := 0

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "Tile") {
			// Title line
			if current != nil {
				photos[gridTitle] = current
			}

			colonIndex := strings.Index(line, ":")
			gridTitle, err = strconv.Atoi(line[5:colonIndex])
			utils.Check(err, "error parsing tile number")

			current = make(map[utils.Point]rune)
			y = 0
			continue
		}

		for x, r := range line {
			current[utils.Point{x, y}] = r
		}
		y++
	}

	photos[gridTitle] = current

	return photos
}

func makeTile(photo map[utils.Point]rune, photoId int, neighbors map[int]int) Tile {
	newNeighbors := []int{}
	for _, v := range neighbors {
		newNeighbors = append(newNeighbors, v)
	}

	return Tile{photoId, photo, newNeighbors}
}

func getBorders(photo map[utils.Point]rune) Borders {
	result := make(Borders)

	result[Top] = make([]rune, 10)
	result[Right] = make([]rune, 10)
	result[Bottom] = make([]rune, 10)
	result[Left] = make([]rune, 10)
	result[TopRev] = make([]rune, 10)
	result[RightRev] = make([]rune, 10)
	result[BottomRev] = make([]rune, 10)
	result[LeftRev] = make([]rune, 10)

	for i := 0; i < 10; i++ {
		result[Top][i] = photo[utils.Point{i, 0}]
		result[Bottom][i] = photo[utils.Point{i, 9}]
		result[TopRev][9-i] = photo[utils.Point{i, 0}]
		result[BottomRev][9-i] = photo[utils.Point{i, 9}]

		result[Right][i] = photo[utils.Point{9, i}]
		result[Left][i] = photo[utils.Point{0, i}]
		result[RightRev][9-i] = photo[utils.Point{9, i}]
		result[LeftRev][9-i] = photo[utils.Point{0, i}]
	}

	return result
}

func equalBorders(a []rune, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i, va := range a {
		if va != b[i] {
			return false
		}
	}
	return true
}

func photoSize(photo map[utils.Point]rune) (int, int) {
	maxX := 0
	maxY := 0
	for k := range photo {
		if k.X > maxX {
			maxX = k.X
		}
		if k.Y > maxY {
			maxY = k.Y
		}
	}

	return maxX + 1, maxY + 1
}
func rotateRight(photo map[utils.Point]rune, width int, height int) map[utils.Point]rune {
	result := make(map[utils.Point]rune)

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			result[utils.Point{height - j - 1, i}] = photo[utils.Point{i, j}]
		}
	}

	return result
}

func flipPhoto(photo map[utils.Point]rune, width int, height int) map[utils.Point]rune {
	result := make(map[utils.Point]rune)

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			result[utils.Point{width - i - 1, j}] = photo[utils.Point{i, j}]
		}
	}

	return result
}

func makeImageFromTiles(tiles map[utils.Point]Tile) map[utils.Point]rune {
	fullPhoto := make(map[utils.Point]rune)

	for pt, tile := range tiles {
		for subp, value := range tile.Photo {
			if subp.X == 0 || subp.X == 9 || subp.Y == 0 || subp.Y == 9 {
				continue
			}
			x := pt.X*8 + subp.X - 1
			y := pt.Y*8 + subp.Y - 1

			fullPhoto[utils.Point{x, y}] = value
		}
	}

	return fullPhoto
}

func printPhoto(photo map[utils.Point]rune) {
	maxX, maxY := photoSize(photo)

	for j := 0; j < maxY; j++ {
		for i := 0; i < maxX; i++ {
			r := photo[utils.Point{i, j}]
			if r == 'R' {
				fmt.Printf("\033[35;1m#\033[0m")
			} else {
				fmt.Printf("%c", r)
			}
		}
		fmt.Printf("\n")
	}
}
