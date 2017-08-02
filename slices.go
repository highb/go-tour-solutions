package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var data = make([][]uint8, dx)

	for x := 0; x < len(data); x++ {
	  data[x] = make([]uint8, dy)
	  for y := 0; y < len(data[x]); y++ {
	  	data[x][y] = uint8(x) ^ uint8(y)
	  }
	}

	return data
}

func main() {
	pic.Show(Pic)
}
