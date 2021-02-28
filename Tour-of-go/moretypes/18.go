// Implemented by Hunachi

import (
	"golang.org/x/tour/pic"
)

func plus(x,y int) uint8 {
	return uint8((x + y)/2)
}

func multi(x,y int) uint8 {
	return uint8(x * y)
}

func pow(x,y int) uint8 {
	return uint8(x << uint8(y))
}

func Pic(dx, dy int) [][]uint8 {
	result := [][]uint8{}
	for i := 0; i < dy; i++ {
		result = append(result, make([]uint8, dx, dx))
		for j := 0; j < dx; j++ {
			result[i][j] = plus(i,j)
		}
	}
	return result
}

func main() {
	pic.Show(Pic)
}