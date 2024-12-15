package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
	"os"
	"regexp"
	"strconv"
)

type robot struct {
	iPosX int
	iPosY int
	velX  int
	velY  int
	posX  int
	posY  int
	fPosX int
	fPosY int
}

var gridX = 11
var gridY = 7
var predictTime = 100

func readInput() (robots []*robot) {
	fileName := "input"
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if fileName == "input" {
		gridX = 101
		gridY = 103
	}

	r := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		val := scanner.Text()

		match := r.FindStringSubmatch(val)
		rob := &robot{}
		rob.iPosX, _ = strconv.Atoi(match[1])
		rob.iPosY, _ = strconv.Atoi(match[2])
		rob.velX, _ = strconv.Atoi(match[3])
		rob.velY, _ = strconv.Atoi(match[4])
		robots = append(robots, rob)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func main() {
	sol1()
	sol2()
}

func sol1() {
	robots := readInput()

	// grid is always odd, div/2 and drop the half to get the quandrant boundries
	halfX := gridX / 2
	halfY := gridY / 2

	quadrant1 := 0
	quadrant2 := 0
	quadrant3 := 0
	quadrant4 := 0

	for _, r := range robots {
		r.fPosX = (r.iPosX + (r.velX * predictTime)) % gridX
		if r.fPosX < 0 {
			r.fPosX += gridX
		}
		r.fPosY = (r.iPosY + (r.velY * predictTime)) % gridY
		if r.fPosY < 0 {
			r.fPosY += gridY
		}

		if r.fPosX < halfX {
			if r.fPosY < halfY {
				quadrant1++
			} else if r.fPosY > halfY {
				quadrant2++
			}
		} else if r.fPosX > halfX {
			if r.fPosY < halfY {
				quadrant3++
			} else if r.fPosY > halfY {
				quadrant4++
			}
		}
	}

	log.Println(quadrant1 * quadrant2 * quadrant3 * quadrant4)

}

func sol2() {
	robots := readInput()

	var grid [][]rune
	grid = make([][]rune, gridY)
	for i := range grid {
		grid[i] = make([]rune, gridX)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	for _, r := range robots {
		r.posX = r.iPosX
		r.posY = r.iPosY
	}
	var palette = []color.Color{
		color.RGBA{0x00, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0x00, 0xff, 0xff},
		color.RGBA{0x00, 0xff, 0x00, 0xff}, color.RGBA{0x00, 0xff, 0xff, 0xff},
		color.RGBA{0xff, 0x00, 0x00, 0xff}, color.RGBA{0xff, 0x00, 0xff, 0xff},
		color.RGBA{0xff, 0xff, 0x00, 0xff}, color.RGBA{0xff, 0xff, 0xff, 0xff},
	}
	var images []*image.Paletted
	var delays []int

	for i := 0; i < 6243; i++ {
		for _, r := range robots {
			grid[r.posY][r.posX] = '.'
			r.posX += r.velX
			r.posX %= gridX
			if r.posX < 0 {
				r.posX += gridX
			}
			r.posY += r.velY
			r.posY %= gridY
			if r.posY < 0 {
				r.posY += gridY
			}
			grid[r.posY][r.posX] = 'R'
		}

		draw := false
		for y, row := range grid {
			for x, val := range row {
				if val == 'R' {
					// check below and to the right to find lines
					downscore := 0
					for yp := y + 1; yp < gridY; yp++ {
						if grid[yp][x] == 'R' {
							downscore++
						} else {
							break
						}
					}
					if downscore > 5 {
						rightscore := 0
						for xp := x + 1; xp < gridX; xp++ {
							if grid[y][xp] == 'R' {
								rightscore++
							} else {
								break
							}
						}
						if rightscore > 5 {
							draw = true
						}
					}
				}
			}
		}

		// if !draw {
		// 	continue
		// }

		if draw {
			log.Println(i)
		}

		img := image.NewPaletted(image.Rect(0, 0, gridX+100, gridY), palette)
		if draw {
			delays = append(delays, 10000)
		} else {
			delays = append(delays, 1)
		}
		for y, row := range grid {
			for x, val := range row {
				c := color.RGBA{B: 255, A: 255}
				if val == 'R' {
					c = color.RGBA{B: 255, G: 255, A: 255}
				}
				img.Set(x, y, c)
			}
		}

		timeX := i % 100
		timeY := i / 100

		img.Set(gridX+timeX, timeY, color.White)

		// nimg := image.NewPaletted(image.Rect(0, 0, (gridX+100)*4, gridY*4), palette)
		images = append(images, img)

		// for x := 0; x < gridX+100; x++ {
		// 	for y := 0; y < gridY; y++ {
		// 		for w := 0; w < 4; w++ {
		// 			for h := 0; h < 4; h++ {
		// 				nimg.Set(x*4+w, y*4+h, img.At(x, y))
		// 			}
		// 		}
		// 	}
		// }
	}
	f, err := os.OpenFile("out.gif", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
	})
}
