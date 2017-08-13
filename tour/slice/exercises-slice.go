package main

import (
	"golang.org/x/tour/pic"
	// "log"d
)

func main() {
	pic.Show(Pic)
}

func Pic(dx, dy int) [][]uint8 {

	pic := make([][]uint8, dy)

	for k := range pic {
		pic[k] = make([]uint8, dx)
	}

	// log.Fatalln(pic)
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			pic[i][j] = uint8(i * j)
		}
	}
	return pic
}

// func Pic(dx, dy int) [][]uint8 {
//     pic := make([][]uint8, dy)
//     for i := range pic {
//         pic[i] = make([]uint8, dx)
//         for j := range pic {
//             pic[i][j] = uint8((j ^ i) * (10 ^ (i * j)))
//         }
//     }
//     return pic
// }
